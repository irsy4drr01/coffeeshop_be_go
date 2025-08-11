package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepoInterface interface {
	// CreateProduct(data *models.Product) (*models.Product, error)
	GetAllProducts(ctx context.Context, params models.ProductQueryParams) ([]models.ProductDB, int, error)
	GetOneProduct(ctx context.Context, uuid string) (models.ProductDB, []models.ProductImgSlot, []models.Size, error)
	// UpdateProduct(uuid string, body map[string]any) (*models.Product, error)
	// DeleteProduct(uuid string) (*models.DeleteProduct, error)
}

type RepoProduct struct {
	db *sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db: db}
}

// func (r *RepoProduct) CreateProduct(data *models.Product) (*models.Product, error) {
// 	query := `
// 		INSERT INTO public.product (
// 			product_name,
// 			price,
// 			category,
// 			description
// 		)
// 		VALUES (:product_name, :price, :category, :description)
// 		RETURNING id, product_name, price, category, description, created_at, uuid;
// 	`

// 	var product models.Product
// 	stmt, err := r.DB.PrepareNamed(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = stmt.Get(&product, data)
// 	stmt.Close() // Menutup statement setelah digunakan
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }

// ---------------------------------

func (r *RepoProduct) GetAllProducts(ctx context.Context, params models.ProductQueryParams) ([]models.ProductDB, int, error) {
	var products []models.ProductDB
	var total int

	query := `
		SELECT 
			p.id, 
			p.product_name, 
			p.category_id, 
			p.price, 
			p.description, 
			p.total_sold, 
			p.total_likes, 
			p.is_deleted, 
			p.created_at, 
			p.updated_at,
			c.name AS category_name,
			COALESCE((
				SELECT image_file 
				FROM product_images 
				WHERE product_id = p.id 
				ORDER BY slot_number ASC LIMIT 1
			), 'product_default.webp') AS product_img,
			EXISTS (
				SELECT 1 
				FROM discount_products dp 
				JOIN discounts d ON dp.discount_id = d.id
				WHERE dp.product_id = p.id 
				AND d.expired > NOW() 
				AND d.is_actived = TRUE
			) AS is_discount,
			COALESCE((
				SELECT MAX(d.discount)
				FROM discount_products dp 
				JOIN discounts d ON dp.discount_id = d.id
				WHERE dp.product_id = p.id 
					AND d.expired > NOW() 
					AND d.is_actived = TRUE
			), 0) AS discount_rate
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.is_deleted = FALSE
	`

	args := []interface{}{}
	argPos := 1

	// Add search
	if params.SearchProductName != "" {
		query += fmt.Sprintf(" AND p.product_name ILIKE $%d ", argPos)
		args = append(args, "%"+params.SearchProductName+"%")
		argPos++
	}

	// Add sorting
	switch params.SortBy {
	case "newest":
		query += " ORDER BY p.created_at DESC "
	case "oldest":
		query += " ORDER BY p.created_at ASC "
	case "asc":
		query += " ORDER BY p.product_name ASC "
	case "desc":
		query += " ORDER BY p.product_name DESC "
	case "cheapest":
		query += " ORDER BY p.price ASC "
	case "priciest":
		query += " ORDER BY p.price DESC "
	case "most_liked":
		query += " ORDER BY p.total_likes DESC "
	default:
		query += `
			ORDER BY 
				CASE WHEN EXISTS (
					SELECT 1 
					FROM discount_products dp 
					JOIN discounts d ON dp.discount_id = d.id 
					WHERE dp.product_id = p.id 
						AND d.expired > NOW() 
						AND d.is_actived = TRUE
				) THEN 0 ELSE 1 END,
				c.name ASC		
		`
	}

	// Pagination
	offset := (params.Page - 1) * params.Limit
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d ", argPos, argPos+1)
	args = append(args, params.Limit, offset)

	err := r.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		log.Printf("[RepoProduct][GetAllProducts] failed to fetch products: %v", err)
		return nil, 0, fmt.Errorf("failed to fetch products from DB")
	}

	// Total count for pagination
	countArgs := []interface{}{}

	countQuery := `
		SELECT COUNT(*) 
		FROM products p
		WHERE p.is_deleted = FALSE
	`
	if params.SearchProductName != "" {
		countQuery += fmt.Sprintf(" AND p.product_name ILIKE $%d ", 1)
		countArgs = append(countArgs, "%"+params.SearchProductName+"%")
	}

	err = r.db.GetContext(ctx, &total, countQuery, countArgs...)
	if err != nil {
		log.Printf("[RepoProduct][GetAllProducts] failed to count products: %v", err)
		return nil, 0, fmt.Errorf("failed to count products from DB")
	}

	return products, total, nil
}

func (r *RepoProduct) GetOneProduct(ctx context.Context, uuid string) (models.ProductDB, []models.ProductImgSlot, []models.Size, error) {
	var product models.ProductDB
	var images []models.ProductImgSlot
	var sizes []models.Size

	// --- Query product detail
	// COALESCE((
	// 	SELECT image_file
	// 	FROM product_images
	// 	WHERE product_id = p.id AND slot_number = 1
	// ), 'product_default.webp') AS product_img,

	query := `
		SELECT 
			p.id, p.product_name, p.category_id, p.price, p.description, 
			p.total_sold, p.total_likes, p.is_deleted, p.created_at, p.updated_at,
			c.name AS category_name,
			EXISTS (
				SELECT 1 FROM discount_products dp 
				JOIN discounts d ON dp.discount_id = d.id
				WHERE dp.product_id = p.id AND d.expired > NOW() AND d.is_actived = TRUE
			) AS is_discount,
			COALESCE((
				SELECT MAX(d.discount)
				FROM discount_products dp 
				JOIN discounts d ON dp.discount_id = d.id
				WHERE dp.product_id = p.id AND d.expired > NOW() AND d.is_actived = TRUE
			), 0) AS discount_rate
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1 AND p.is_deleted = FALSE
	`

	err := r.db.GetContext(ctx, &product, query, uuid)
	if err != nil {
		log.Printf("[RepoProduct][GetOneProduct] failed to fetch product: %v", err)
		return product, nil, nil, fmt.Errorf("product not found")
	}

	// --- Query slot images
	queryImgs := `
		SELECT slot_number, image_file
		FROM product_images
		WHERE product_id = $1
		ORDER BY slot_number ASC
	`
	err = r.db.SelectContext(ctx, &images, queryImgs, uuid)
	if err != nil {
		log.Printf("[RepoProduct][GetOneProduct] failed to fetch slot images: %v", err)
	}

	// --- Query sizes only for drink
	if product.CategoryID == 1 || product.CategoryID == 2 {
		querySizes := `SELECT id as size_id, name FROM sizes WHERE id != 4 ORDER BY id ASC`
		err = r.db.SelectContext(ctx, &sizes, querySizes)
		if err != nil {
			log.Printf("[RepoProduct][GetOneProduct] failed to fetch sizes: %v", err)
		}
	}

	return product, images, sizes, nil
}

// func (r *RepoProduct) UpdateProduct(uuid string, body map[string]any) (*models.Product, error) {
// 	query := `UPDATE product SET `
// 	params := map[string]interface{}{}

// 	if productName, exists := body["product_name"]; exists {
// 		query += "product_name = :product_name, "
// 		params["product_name"] = productName
// 	}
// 	if price, exists := body["price"]; exists {
// 		query += "price = :price, "
// 		params["price"] = price
// 	}
// 	if category, exists := body["category"]; exists {
// 		query += "category = :category, "
// 		params["category"] = category
// 	}
// 	if image, exists := body["image"]; exists {
// 		query += "image = :image, "
// 		params["image"] = image
// 	}
// 	if description, exists := body["description"]; exists {
// 		query += "description = :description, "
// 		params["description"] = description
// 	}

// 	query += "updated_at = NOW() WHERE uuid = :uuid RETURNING id, product_name, price, category, image, description, updated_at, uuid"
// 	params["uuid"] = uuid

// 	var product models.Product
// 	stmt, args, err := sqlx.Named(query, params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	stmt = r.Rebind(stmt) // Rebind the statement according to the driver
// 	if err := r.QueryRowx(stmt, args...).StructScan(&product); err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }

// ---------------------------------

// func (r *RepoProduct) DeleteProduct(uuid string) (*models.DeleteProduct, error) {
// 	query := `
// 		UPDATE public.product
// 		SET
// 			is_deleted = true
// 		WHERE uuid = $1
// 		RETURNING id, uuid, product_name, is_deleted;
// 	`

// 	var product models.DeleteProduct
// 	if err := r.Get(&product, query, uuid); err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }

// ---------------------------------

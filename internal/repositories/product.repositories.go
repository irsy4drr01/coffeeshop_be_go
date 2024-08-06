package repositories

import (
	"strconv"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(data *models.Product) (string, *models.CreateProductResponse, error) {
	query := `
		INSERT INTO public.product (
			product_name,
			price,
			category,
			description
		)
		VALUES (:product_name, :price, :category, :description)
		RETURNING id, product_name, price, category, description, created_at, uuid;
	`

	var product models.CreateProductResponse
	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return "", nil, err
	}

	err = stmt.Get(&product, data)
	stmt.Close() // Menutup statement setelah digunakan
	if err != nil {
		return "", nil, err
	}
	return "Product created successfully.", &product, nil
}

func (r *RepoProduct) GetAllProducts(searchProductName string, minPrice int, maxPrice int, category string, sort string) (*models.Products, error) {
	query := `
        SELECT
            id,
            product_name,
            price,
            category,
            description,
            created_at,
            updated_at,
            uuid,
            is_deleted
        FROM public.product
        WHERE is_deleted = false
    `
	params := []any{}
	paramIndex := 1

	// Apply filters
	if searchProductName != "" {
		query += ` AND product_name ILIKE $` + strconv.Itoa(paramIndex)
		params = append(params, "%"+searchProductName+"%")
		paramIndex++
	}
	if minPrice > 0 {
		query += ` AND price >= $` + strconv.Itoa(paramIndex)
		params = append(params, minPrice)
		paramIndex++
	}
	if maxPrice > 0 {
		query += ` AND price <= $` + strconv.Itoa(paramIndex)
		params = append(params, maxPrice)
		paramIndex++
	}
	if category != "" {
		query += ` AND category = $` + strconv.Itoa(paramIndex)
		params = append(params, category)
		paramIndex++
	}

	// Sorting logic
	switch sort {
	case "a-z":
		query += ` ORDER BY product_name ASC`
	case "z-a":
		query += ` ORDER BY product_name DESC`
	case "cheapest":
		query += ` ORDER BY price ASC`
	case "priciest":
		query += ` ORDER BY price DESC`
	case "oldest":
		query += ` ORDER BY created_at ASC`
	case "newest":
		query += ` ORDER BY created_at DESC`
	default:
		query += ` ORDER BY created_at DESC`
	}

	data := models.Products{}
	if err := r.Select(&data, query, params...); err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoProduct) UpdateProduct(uuid string, body map[string]any) (string, *models.UpdateProductResponse, error) {
	query := `UPDATE product SET `
	params := map[string]interface{}{}

	if productName, exists := body["product_name"]; exists {
		query += "product_name = :product_name, "
		params["product_name"] = productName
	}
	if price, exists := body["price"]; exists {
		query += "price = :price, "
		params["price"] = price
	}
	if category, exists := body["category"]; exists {
		query += "category = :category, "
		params["category"] = category
	}
	if description, exists := body["description"]; exists {
		query += "description = :description, "
		params["description"] = description
	}

	query += "updated_at = NOW() WHERE uuid = :uuid RETURNING id, product_name, price, category, description, updated_at, uuid"
	params["uuid"] = uuid

	var product models.UpdateProductResponse
	stmt, args, err := sqlx.Named(query, params)
	if err != nil {
		return "", nil, err
	}

	stmt = r.Rebind(stmt) // Rebind the statement according to the driver
	if err := r.QueryRowx(stmt, args...).StructScan(&product); err != nil {
		return "", nil, err
	}
	return "Product updated successfully.", &product, nil
}

func (r *RepoProduct) DeleteProduct(uuid string) (string, *models.DeleteProductResponse, error) {
	query := `
		UPDATE public.product
		SET
			is_deleted = true
		WHERE uuid = $1
		RETURNING id, product_name, is_deleted;
	`

	var product models.DeleteProductResponse
	if err := r.Get(&product, query, uuid); err != nil {
		return "", nil, err
	}
	return "Product deleted successfully.", &product, nil
}

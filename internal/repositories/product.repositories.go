package repositories

import (
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

func (r *RepoProduct) GetAllProducts() (*models.Products, error) {
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
		ORDER BY created_at DESC;
	`
	data := models.Products{}

	if err := r.Select(&data, query); err != nil {
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

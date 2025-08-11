package repositories

import (
	"context"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type OrderHistoryRepoInterface interface {
	GetAllOrderHistories(ctx context.Context, userID string) (models.OrderHistoriesDB, error)
	GetOrderHistoryDetails(ctx context.Context, orderID string, userID string) (models.OrderHistoryDetailDB, error)
}

type RepoOrderHistory struct {
	db *sqlx.DB
}

func NewOrderHistory(db *sqlx.DB) *RepoOrderHistory {
	return &RepoOrderHistory{db: db}
}

func (r *RepoOrderHistory) GetAllOrderHistories(ctx context.Context, userID string) (models.OrderHistoriesDB, error) {
	var orderHistoriesDB models.OrderHistoriesDB

	query := `
		select
			o.id as order_id,
			COALESCE((
				SELECT image_file 
				FROM product_images 
				WHERE product_id = p.id AND slot_number = 1
			), 'product_default.webp') AS image,
			s.status,
			o.created_at,
			o.total_amount
		from orders o
		join statuses s on o.status_id = s.id
		join order_details od on o.id = od.order_id
		join products p on od.product_id = p.id
		where o.user_id = $1;
	`

	err := r.db.SelectContext(ctx, &orderHistoriesDB, query, userID)
	if err != nil {
		return nil, err
	}

	return orderHistoriesDB, nil
}

func (r *RepoOrderHistory) GetOrderHistoryDetails(ctx context.Context, orderID string, userID string) (models.OrderHistoryDetailDB, error) {
	var detail models.OrderHistoryDetailDB

	queryHeader := `
		SELECT
			o.id AS order_id,
			o.fullname,
			o.address,
			o.phone,			
			o.total_purchase,
			o.payment_method,
			o.delivery_method,
			o.delivery_method_fee,
			o.tax,
			o.total_amount,
			s.status AS status_name,
			o.created_at
		FROM orders o
		JOIN statuses s ON o.status_id = s.id
		WHERE o.id = $1 AND o.user_id = $2
		LIMIT 1;
	`

	err := r.db.GetContext(ctx, &detail, queryHeader, orderID, userID)
	if err != nil {
		return models.OrderHistoryDetailDB{}, err
	}

	var items models.OrderHistoryItemsDB
	queryItems := `
		SELECT
			pr.product_name,
			COALESCE((
				SELECT image_file FROM product_images 
				WHERE product_id = pr.id AND slot_number = 1
			), 'product_default.webp') AS image,
			od.qty,
			sz.name AS size_name,
			od.is_iced,
			pr.price AS base_price,
			od.final_price AS finale_price,
			od.discount_name,
			pr.category_id
		FROM order_details od
		JOIN products pr ON od.product_id = pr.id
		JOIN sizes sz ON od.size_id = sz.id
		WHERE od.order_id = $1;
	`

	err = r.db.SelectContext(ctx, &items, queryItems, orderID)
	if err != nil {
		return models.OrderHistoryDetailDB{}, err
	}

	detail.Items = items
	return detail, nil
}

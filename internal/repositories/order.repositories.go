package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/utils"
)

type OrderRepoInterface interface {
	CreateOrder(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderRepoResult, error)
}

type RepoOrder struct {
	db *sqlx.DB
}

func NewOrder(db *sqlx.DB) *RepoOrder {
	return &RepoOrder{db: db}
}

func (r *RepoOrder) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderRepoResult, error) {
	uuidClean := strings.ReplaceAll(uuid.New().String(), "-", "")
	orderID := fmt.Sprintf("TRX%s%s", time.Now().Format("20060102150405"), strings.ToUpper(uuidClean[:15]))

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return models.CreateOrderRepoResult{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// get helper data
	// Ambil phone dari profile jika ada, fallback ke req.Phone ===
	var phone string
	err = tx.GetContext(ctx, &phone, `SELECT phone FROM profiles WHERE id = $1`, req.UserID)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}
	if phone == "" {
		phone = req.Phone
	}
	if phone == "" {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, fmt.Errorf("phone is required")
	}

	var delivery models.DeliveryMethod
	err = tx.GetContext(ctx, &delivery, `SELECT id, name, fee FROM delivery_methods WHERE id = $1`, req.DeliveryMethodID)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}

	var payment models.PaymentMethod
	err = tx.GetContext(ctx, &payment, `SELECT id, name FROM payment_methods WHERE id = $1`, req.PaymentMethodID)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}

	var tax models.Tax
	err = tx.GetContext(ctx, &tax, `SELECT id, tax_value FROM tax ORDER BY id DESC LIMIT 1`)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}

	var iceCubePrice decimal.Decimal
	err = tx.GetContext(ctx, &iceCubePrice, `SELECT price FROM products WHERE product_name ILIKE 'Ice Cube' LIMIT 1`)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}

	// INSERT TO orders (initial amount 0, update later)
	_, err = tx.ExecContext(ctx, `
		INSERT INTO orders 
		(id, user_id, fullname, address, phone, payment_method, delivery_method, delivery_method_fee, tax, status_id, total_purchase, total_amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 1, 0, 0)
	`,
		orderID,
		req.UserID,
		req.Fullname,
		req.Address,
		phone,
		payment.Name,
		delivery.Name,
		delivery.Fee,
		tax.TaxValue,
	)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}

	// LOOP ITEM -> insert order_details, hitung total
	var items []models.OrderDetailsResponse
	totalPurchase := decimal.Zero

	for _, item := range req.Items {
		var product models.ProductWithSizeAndDiscount
		err = tx.GetContext(ctx, &product, `
			SELECT 
				p.id as product_id, 
				p.product_name, 
				p.category_id,
				p.price, 
				s.id as size_id, 
				s.name as size_name, 
				s.additional_price,
				d.id as discount_id,
				d.name as discount_name,
				d.discount as discount_value,
				d.is_actived,
				d.expired
			FROM products p
			JOIN sizes s ON s.id = $2
			LEFT JOIN discount_products dp ON dp.product_id = p.id
			LEFT JOIN discounts d ON dp.discount_id = d.id AND d.is_actived = TRUE AND d.expired > NOW()
			WHERE p.id = $1
		`, item.ProductID, item.SizeID)
		if err != nil {
			tx.Rollback()
			return models.CreateOrderRepoResult{}, err
		}

		if product.CategoryID != 1 && product.CategoryID != 2 {
			item.SizeID = 4
			item.IsIced = false
			product.SizeName = "-"
			product.Additional = 0.0
		}

		basePrice := product.BasePrice.Mul(decimal.NewFromFloat(1 + product.Additional))

		var isDiscount bool
		var discountName string
		var discountPercent float64

		if product.DiscountID != nil && product.DiscountValue != nil && product.DiscountActive != nil && *product.DiscountActive {
			isDiscount = true
			discountName = *product.DiscountName
			discountPercent = *product.DiscountValue
		} else {
			isDiscount = false
			discountName = ""
			discountPercent = 0.0
		}

		finalPrice := basePrice.Mul(decimal.NewFromFloat(1 - discountPercent))
		if item.IsIced {
			finalPrice = finalPrice.Add(iceCubePrice)
		}

		subTotal := finalPrice.Mul(decimal.NewFromInt(int64(item.Qty)))
		totalPurchase = totalPurchase.Add(subTotal)

		// var discount models.Discount
		// err = tx.GetContext(ctx, &discount, `
		// 	SELECT d.id, d.name, d.discount
		// 	FROM discounts d
		// 	JOIN discount_products dp ON dp.discount_id = d.id
		// 	WHERE dp.product_id = $1 AND d.is_actived = TRUE AND d.expired > NOW()
		// 	LIMIT 1
		// `, product.ProductID)
		// if err != nil && err != sql.ErrNoRows {
		// 	tx.Rollback()
		// 	return models.CreateOrderRepoResult{}, err
		// }

		// if err == nil {
		// 	isDiscount = true
		// 	discountName = discount.Name
		// 	discountPercent = discount.Discount
		// } else {
		// 	isDiscount = false
		// 	discountName = ""
		// 	discountPercent = 0.0
		// }

		// finalPrice := basePrice.Mul(decimal.NewFromFloat(1 - discountPercent))

		// if item.IsIced {
		// 	finalPrice = finalPrice.Add(iceCubePrice)
		// }

		// subTotal := finalPrice.Mul(decimal.NewFromInt(int64(item.Qty)))
		// totalPurchase = totalPurchase.Add(subTotal)

		// if product.DiscountValue != nil && utils.CheckDiscountValid(*product.DiscountExpiry, product.DiscountActive) {
		// 	discount = *product.DiscountValue
		// }
		// finalPrice := utils.CalculateFinalPrice(basePrice, discount)
		// if item.IsIced {
		// 	finalPrice = finalPrice.Add(iceCubePrice)
		// }
		// subTotal := finalPrice.Mul(decimal.NewFromInt(int64(item.Qty)))
		// totalPurchase = totalPurchase.Add(subTotal)

		// CEK & KURANGI STOCK
		var currentStock int
		err = tx.GetContext(ctx, &currentStock, `
			SELECT stock FROM product_stocks 
			WHERE product_id = $1 AND size_id = $2
			FOR UPDATE
		`, product.ProductID, item.SizeID)
		if err != nil {
			tx.Rollback()
			return models.CreateOrderRepoResult{}, fmt.Errorf("stock check failed: %w", err)
		}

		if currentStock < item.Qty {
			tx.Rollback()
			return models.CreateOrderRepoResult{}, fmt.Errorf("not enough stock for %s (size: %s). Please adjust the quantity or size", product.ProductName, product.SizeName)
		}

		_, err = tx.ExecContext(ctx, `
			UPDATE product_stocks 
			SET stock = stock - $1, updated_at = NOW()
			WHERE product_id = $2 AND size_id = $3
		`, item.Qty, product.ProductID, item.SizeID)
		if err != nil {
			tx.Rollback()
			return models.CreateOrderRepoResult{}, fmt.Errorf("failed to reduce stock: %w", err)
		}

		_, err = tx.ExecContext(ctx, `
			INSERT INTO order_details 
			(order_id, product_id, qty, size_id, is_iced, is_discount, discount_name, base_price, final_price, sub_total)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		`,
			orderID,
			product.ProductID,
			item.Qty,
			item.SizeID,
			item.IsIced,
			isDiscount,
			discountName,
			basePrice,
			finalPrice,
			subTotal,
		)
		if err != nil {
			tx.Rollback()
			return models.CreateOrderRepoResult{}, err
		}

		baseStr, finalStr := utils.CalculatePriceAndFinal(basePrice, isDiscount, decimal.NewFromFloat(discountPercent))

		items = append(items, models.OrderDetailsResponse{
			ProductID:   product.ProductID,
			ProductName: product.ProductName,
			Qty:         item.Qty,
			Size:        product.SizeName,
			IsIced:      item.IsIced,
			BasePrice:   baseStr,
			FinalPrice:  finalStr,
			SubTotal:    fmt.Sprintf("Rp. %s", subTotal.StringFixed(0)),
		})
	}

	// Update orders hanya SEKALI setelah loop
	taxAmount := totalPurchase.Mul(decimal.NewFromFloat(tax.TaxValue))
	totalAmount := totalPurchase.Add(decimal.NewFromInt(int64(delivery.Fee))).Add(taxAmount)

	_, err = tx.ExecContext(ctx, `
		UPDATE orders 
		SET total_purchase = $1, total_amount = $2
		WHERE id = $3
	`,
		totalPurchase,
		totalAmount,
		orderID,
	)
	if err != nil {
		tx.Rollback()
		return models.CreateOrderRepoResult{}, err
	}

	// Commit
	if err = tx.Commit(); err != nil {
		return models.CreateOrderRepoResult{}, err
	}

	resp := models.CreateOrderRepoResult{
		OrderID:        orderID,
		Phone:          phone,
		Items:          items,
		DeliveryMethod: delivery,
		PaymentMethod:  payment,
		TaxAmount:      taxAmount,
		TotalPurchase:  totalPurchase,
		TotalAmount:    totalAmount,
	}

	// Success response
	return resp, nil
}

package utils

import (
	"net/url"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
)

func BuildQueryString(params models.ProductQueryParams) string {
	query := url.Values{}

	if params.SearchProductName != "" {
		query.Add("search_product_name", params.SearchProductName)
	}
	if params.SortBy != "" {
		query.Add("sort_by", params.SortBy)
	}
	if params.MinPrice != "" {
		query.Add("min_price", params.MinPrice)
	}
	if params.MaxPrice != "" {
		query.Add("max_price", params.MaxPrice)
	}
	if params.Discount {
		query.Add("discount", "true")
	}

	return query.Encode()
}

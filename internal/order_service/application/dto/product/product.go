package order_api_product_dto

type ProductDTO struct {
	ID          string    `json:"id"`
	OrderID     *string   `json:"order_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	ShopUrl     string    `json:"shop_url"`
	ShopName    string    `json:"shop_name"`
	ImageUrls   *[]string `json:"image_urls"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   *string   `json:"updated_at"`
}

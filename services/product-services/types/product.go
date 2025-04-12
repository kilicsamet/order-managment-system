package types

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ListProductsResponse struct {
	Products []Product `json:"products"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

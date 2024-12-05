package response

type (
	CreateProductResponse struct {
		ID uint `json:"id"`
	}
	GetProductByIdResponse struct {
		ID          uint   `json:"id"`
		NamaProduct string `json:"nama_product"`
		Harga       uint   `json:"harga"`
		Quantity    uint   `json:"quantity"`
		IdBrand     uint   `json:"id_brand"`
	}
	UpdateProductResponse struct {
		NamaProduct string `json:"nama_product,omitempty"`
		Harga       uint   `json:"harga,omitempty"`
		Quantity    uint   `json:"quantity,omitempty"`
		IdBrand     uint   `json:"id_brand,omitempty"`
	}
	GetAllProductsDataResponse struct {
		ID          uint   `json:"id"`
		NamaProduct string `json:"nama_product"`
		Harga       uint   `json:"harga"`
		Quantity    uint   `json:"quantity"`
		IdBrand     uint   `json:"id_brand"`
	}
	PageInformation struct {
		CurrentPage uint `json:"current_page"`
		PageSize    uint `json:"page_size"`
		Records     uint `json:"records"`
		TotalPage   uint `json:"total_page"`
	}
	GetAllProductsResponse struct {
		PageInformation PageInformation              `json:"page_information"`
		ListData        []GetAllProductsDataResponse `json:"list_data"`
	}
)

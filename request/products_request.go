package request

type (
	CreateProductRequest struct {
		NamaProduct string `json:"nama_product"`
		Harga       uint   `json:"harga"`
		Quantity    uint   `json:"quantity"`
		IdBrand     uint   `json:"id_brand"`
	}
	UpdateProductRequest struct {
		NamaProduct string `json:"nama_product"`
		Harga       uint   `json:"harga"`
		Quantity    uint   `json:"quantity"`
		IdBrand     uint   `json:"id_brand"`
	}
	GetAllProductsRequest struct {
		Page, Size uint
	}
)

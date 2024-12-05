package entities

type Products struct {
	ID          uint   `gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	NamaProduct string `gorm:"column:nama_product;type:varchar(100);default:''"`
	Harga       uint   `gorm:"column:harga;type:integer;default:0"`
	Quantity    uint   `gorm:"column:quantity;type:integer;default:0"`
	IdBrand     uint   `gorm:"column:id_brand;type:integer;default:null"`
}

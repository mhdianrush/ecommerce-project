package entities

type Brands struct {
	IdBrand   uint       `gorm:"column:id_brand;type:integer;primaryKey;autoIncrement;not null"`
	NamaBrand string     `gorm:"column:nama_brand;type:varchar(100);default:''"`
	Products  []Products `gorm:"foreignKey:IdBrand;references:IdBrand;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

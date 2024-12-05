package config

import (
	"log"

	"github.com/mhdianrush/ecommerce-project/entities"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// var logger = logrus.New()
var DB *gorm.DB

func ConnectDB() {
	viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASS")
	name := viper.GetString("DB_NAME")
	applicationName := viper.GetString("APPLICATION_NAME")

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + name + " port=" + port + " sslmode=disable" + " TimeZone=UTC" + " application_name=" + applicationName

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("DB_PREFIX"),
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(
		&entities.Brands{},
		&entities.Products{},
	)

	DB = db
}

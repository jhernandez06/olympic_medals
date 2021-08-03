package data_base

import (
	"olympic-medals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Config() {
	dsn := "host=localhost user=wawandco password=wawandco dbname=wawandco port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.CountryMedal{})
	models.Position()

	for i := 0; i < len(models.AllCountries); i++ {
		models.AllCountries[i].AllMedals = models.AllCountries[i].GoldMedals +
			models.AllCountries[i].SilverMedals + models.AllCountries[i].BronzeMedals
		models.AllCountries[i].Position = i + 1
	}
	db.Create(&models.AllCountries)
}

package api

import (
	"encoding/json"
	"net/http"
	"olympic-medals/models"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {

	dsn := "host=localhost user=wawandco password=wawandco dbname=wawandco port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		panic("failed to connect database")
	}
	var allCountries []models.CountryMedal
	db.Find(&allCountries)
	json.NewEncoder(w).Encode(&allCountries)
}

func GetCountry(w http.ResponseWriter, r *http.Request) {

	dsn := "host=localhost user=wawandco password=wawandco dbname=wawandco port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		panic("failed to connect database")
	}
	params := mux.Vars(r)

	var country models.CountryMedal
	db.First(&country, params["id"])
	json.NewEncoder(w).Encode(&country)

}

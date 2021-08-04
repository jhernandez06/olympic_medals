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
	if err != nil {
		panic("failed to connect database")
	}
	defer sqlDB.Close()
	var AllCountries []models.CountryMedal
	db.Find(&AllCountries)
	models.Position(AllCountries)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&AllCountries)
}

func GetCountry(w http.ResponseWriter, r *http.Request) {

	dsn := "host=localhost user=wawandco password=wawandco dbname=wawandco port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	defer sqlDB.Close()
	params := mux.Vars(r)
	var country models.CountryMedal
	db.First(&country, params["id"])
	json.NewEncoder(w).Encode(&country)

}

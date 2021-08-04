package models_test

import (
	"olympic-medals/models"
	"testing"
)

var TestCountries = []models.CountryMedal{
	{
		Country:      "China",
		GoldMedals:   20,
		SilverMedals: 14,
		BronzeMedals: 13,
	},
	{
		Country:      "Estados Unidos",
		GoldMedals:   20,
		SilverMedals: 23,
		BronzeMedals: 16,
	},
	{
		Country:      "Japon",
		GoldMedals:   17,
		SilverMedals: 5,
		BronzeMedals: 9,
	},
	{
		Country:      "COR",
		GoldMedals:   12,
		SilverMedals: 19,
		BronzeMedals: 15,
	},
	{
		Country:      "Australia",
		GoldMedals:   12,
		SilverMedals: 19,
		BronzeMedals: 14,
	},
	{
		Country:      "Colombia",
		GoldMedals:   30,
		SilverMedals: 10,
		BronzeMedals: 14,
	},
}

func Test_Position(t *testing.T) {
	models.Position(TestCountries)

	if TestCountries[0].Country != "Colombia" {
		t.Error("The position is not correct")
		t.Fail()
	} else {
		t.Log("Correcto")
	}
	if TestCountries[2].Country != "China" {
		t.Error("The position is not correct")
		t.Fail()
	} else {
		t.Log("Correcto")
	}
	if TestCountries[5].Country != "Australia" {
		t.Error("The position is not correct")
		t.Fail()
	} else {
		t.Log("Correcto")
	}

}

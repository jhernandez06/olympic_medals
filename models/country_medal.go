package models

import (
	"sort"

	"gorm.io/gorm"
)

type CountryMedal struct {
	gorm.Model
	Position     int
	Country      string
	GoldMedals   int64
	SilverMedals int64
	BronzeMedals int64
	AllMedals    int64
}

var AllCountries = []CountryMedal{
	{

		Country:      "China",
		GoldMedals:   24,
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
		BronzeMedals: 3,
	},
	{
		Country:      "Australia",
		GoldMedals:   14,
		SilverMedals: 3,
		BronzeMedals: 14,
	},
	{
		Country:      "Gran Bretaña",
		GoldMedals:   10,
		SilverMedals: 10,
		BronzeMedals: 12,
	},
	{
		Country:      "Francia",
		GoldMedals:   5,
		SilverMedals: 10,
		BronzeMedals: 6,
	},
	{
		Country:      "Italia",
		GoldMedals:   4,
		SilverMedals: 10,
		BronzeMedals: 15,
	},
	{
		Country:      "Paises Bajos",
		GoldMedals:   4,
		SilverMedals: 12,
		BronzeMedals: 5,
	},
	{
		Country:      "Alemania",
		GoldMedals:   4,
		SilverMedals: 10,
		BronzeMedals: 12,
	},
	{
		Country:      "Nueva Zelanda",
		GoldMedals:   4,
		SilverMedals: 10,
		BronzeMedals: 14,
	},
	{
		Country:      "Colombia",
		GoldMedals:   30,
		SilverMedals: 10,
		BronzeMedals: 14,
	},
}

//Position
type lessFunc func(p1, p2 *CountryMedal) bool

type multiSorter struct {
	allCountries []CountryMedal
	less         []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(allCountries []CountryMedal) {
	ms.allCountries = allCountries
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.allCountries)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.allCountries[i], ms.allCountries[j] = ms.allCountries[j], ms.allCountries[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.allCountries[i], &ms.allCountries[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	return ms.less[k](p, q)
}

func Position() {
	// Closures that order the Change structure.
	gold := func(c1, c2 *CountryMedal) bool {
		return c1.GoldMedals > c2.GoldMedals
	}
	silver := func(c1, c2 *CountryMedal) bool {
		return c1.SilverMedals > c2.SilverMedals
	}
	bronze := func(c1, c2 *CountryMedal) bool {
		return c1.BronzeMedals > c2.BronzeMedals
	}

	OrderedBy(gold, silver, bronze).Sort(AllCountries)
	//fmt.Println("By Gold,Silver,Bronze:", allCountries)

}

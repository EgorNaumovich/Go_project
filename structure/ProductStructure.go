package structure

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID               int     `json:"id"`
	Type             string  `json:"type"`
	Name             string  `json:"name"`
	Amount           int     `json:"amount"`
	ShortDescription string  `json:"description"`
	Price            float32 `json:"price"`
	CreatedOn        string  `json:"-"`
	DeletedOn        string  `json:"-"`
}

var PList = []*Product{
	{
		ID:               1,
		Type:             "Headphones",
		Name:             "WH-CH100XM4B",
		Amount:           7,
		ShortDescription: "Wireless headphones Sony WH-CH100XM4B",
		Price:            111.99,
		CreatedOn:        time.Now().UTC().String(),
	},
	{
		ID:               2,
		Type:             "Smartphones",
		Name:             "Galaxy A03s",
		Amount:           3,
		ShortDescription: "Very good smartphone by Samsung",
		Price:            599.95,
		CreatedOn:        time.Now().UTC().String(),
	},
}

func GetPList() Products {
	return PList
}

type Products []*Product

func (p *Products) ToJSON(write io.Writer) error {
	err := json.NewEncoder(write)
	return err.Encode(p)
}

func (p *Product) FromJSON(read io.Reader) error {
	err := json.NewDecoder(read)
	return err.Decode(p)
}

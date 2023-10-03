package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	Id          string    `json:"id" bson:"id"`
	StoreName   string    `json:"storeName" bson:"storeName"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Quantity    int       `json:"quantity" bson:"quantity,omitempty"`
	Price       uint      `json:"price" bson:"price,omitempty"`
	ImageUri    []string  `json:"imageUri" bson:"imageUri"`
	TimeCreated time.Time `json:"timeCreated" bson:"timeCreated"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

type ListProducts struct {
	Products []*Product `json:"products"`
	Count    int64      `json:"count"`
}

type ListProductsParams struct {
	StoreName string `json:"storeName"`
	Limit     int    `json:"limit"`
}

type Cart struct {
	Id        string    `json:"id" bson:"id"`
	UserId    string    `json:"userId" bson:"userId"`
	CartItems []Product `json:"cartItems" bson:"cartItems,omitempty"`
}

type Order struct {
	Id     string `json:"id" bson:"id"`
	CartId string `json:"cartId" bson:"cartId"`
	Price  int    `json:"price" bson:"price"`
}

func (p Product) Validate() error {
	if err := validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Quantity, validation.Required),
		validation.Field(&p.Price, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

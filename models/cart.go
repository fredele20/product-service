package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Cart struct {
	UserId    string    `json:"userId" bson:"userId"`
	CartItems []Product `json:"cartItems" bson:"cartItems,omitempty"`
}

type CartCheckoutResponse struct {
	Cart       Cart `json:"cart"`
	TotalPrice int  `json:"totalPrice" bson:"price"`
}

type AddToCartRequest struct {
	UserId    string
	ProductId string
	Quantity  int `json:"quantity"`
}

type AddToCartResponse struct {

}

func (c AddToCartRequest) Validate() error {
	if err := validation.ValidateStruct(&c,
		validation.Field(&c.Quantity, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

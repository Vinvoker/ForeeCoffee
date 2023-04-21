package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID       int
	Username string
	Role     string
	jwt.StandardClaims
}

type User struct {
	ID       int    `json:"id"`
	Username int    `json:"username"`
	Role     string `json:"role"`
}

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Category   string `json:"category"`
	PictureUrl string `json:"picture_url"`
}

type OrderDetails struct {
	ID       int
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	ID              int            `json:"id"`
	TransactionTime time.Time      `json:"transaction_time"`
	Details         []OrderDetails `json:"details"`
	TotalPrice      int            `json:"total_price"`
}

type Investor struct {
	Username string
	Email    string
}

type ProductDetail struct {
	Name     string
	Quantity int
	Price    int
}

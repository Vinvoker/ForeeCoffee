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

type Branch struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Category   string `json:"category"`
	PictureUrl string `json:"picture_url"`
}

type ProductsDetails struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	Category   string   `json:"category"`
	PictureUrl string   `json:"picture_url"`
	Branch     []Branch `json:"branches"`
}

type BranchProductsForMenu struct { // No variable Quantity because customer doesnt need to see it
	Branch  Branch    `json:"branch"`
	Product []Product `json:"products"`
}

type BranchProductForInsert struct {
	Branch   string  `json:"branch"`
	Product  Product `json:"products"`
	Quantity int     `json:"quantity"`
}

type BranchProduct struct {
	Branch   Branch  `json:"branch"`
	Product  Product `json:"products"`
	Quantity int     `json:"quantity"`
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

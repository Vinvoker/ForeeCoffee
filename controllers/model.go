package controllers

import (
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
	Username string `json:"username"`
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

type ShortProduct struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductForMenu struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Category   string `json:"category"`
	PictureUrl string `json:"picture_url"`
	Status     string `json:"status"`
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
	Branch  Branch           `json:"branch"`
	Product []ProductForMenu `json:"products"`
}

type BranchProductForInsert struct {
	Branch   string         `json:"branch"`
	Product  ProductForMenu `json:"products"`
	Quantity int            `json:"quantity"`
}

type BranchProduct struct {
	Branch   Branch         `json:"branch"`
	Product  ProductForMenu `json:"products"`
	Quantity int            `json:"quantity"`
}

type OrderDetails struct {
	Product  ShortProduct `json:"product"`
	Quantity int          `json:"quantity"`
}

type Order struct {
	ID              int            `json:"id"`
	TransactionTime string         `json:"transaction_time"`
	BranchID        int            `json:"branch_id"`
	Status          string         `json:"status"`
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

type UpdateProductBranch struct {
	Branch   Branch  `json:"Branch"`
	Product  Product `json:"Product"`
	OldStock int     `json:"Old Stock"`
	NewStock int     `json:"New Stock"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

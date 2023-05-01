package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})
var ctx = context.Background()

func StartCRON(c *gin.Context) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	s := gocron.NewScheduler(location)
	// s.Every(1).Minutes().Do(func() {
	// 	SendEmail(c)
	// })
	s.Every(1).Day().At("09.00").Do(func() {
		SendEmail(c)
	})
	s.StartAsync()

	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "CRON job restarted successfully"}
	c.JSON(http.StatusOK, response)
}

func GetOwners(c *gin.Context) []Investor {

	var investor Investor
	var investors []Investor

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT username, email from users WHERE role = 'INVESTOR'")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with dataowner the query"})
		return investors
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&investor.Username, &investor.Email); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "Products not found"})
			return investors
		} else {
			investors = append(investors, investor)
		}
	}
	return investors
}

func RekapOrder(c *gin.Context) []ProductDetail {
	db := connect()
	defer db.Close()

	var productDetails []ProductDetail
	var productDetail ProductDetail

	rows, err := db.Query("SELECT p.productName, SUM(od.quantity), p.price FROM OrderDetails od" +
		" JOIN `Order` o ON o.orderId = od.orderId JOIN Product p ON p.productId=od.productId WHERE" +
		" o.transactionTime >= DATE(NOW() - INTERVAL 1 DAY) AND o.transactionTime < DATE(NOW()) GROUP BY p.productId")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the rekap order query"})
		return productDetails
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&productDetail.Name, &productDetail.Quantity, &productDetail.Price); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "product not found"})
		} else {
			productDetails = append(productDetails, productDetail)
		}
	}

	return productDetails

}

func GetValueFromRedis(productName string) string {
	// cek key-value nya masih ada atau sudah expire
	// kalau expirenya dalam waktu 3 menit atau kurang, refresh cache
	if rdb.TTL(ctx, productName).Val() < 180 {
		CacheProdukGambar()
	}
	res, err := rdb.Get(ctx, productName).Result()
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func SendEmail(c *gin.Context) {
	env_email := os.Getenv("EMAIL")
	env_password := os.Getenv("PASSWORD")

	investors := GetOwners(c)
	productDetails := RekapOrder(c)

	var wg sync.WaitGroup
	wg.Add(len(investors))
	d := gomail.NewDialer("smtp.gmail.com", 587, env_email, env_password)

	// bikin body string disini
	grandTotal := 0
	var strData string
	for _, productDetail := range productDetails {
		totalProduct := productDetail.Quantity * productDetail.Price
		grandTotal += totalProduct
		productPicture := GetValueFromRedis(productDetail.Name)

		strData += "<br><br>Product Name: " + productDetail.Name
		strData += "<br>Product Price: Rp" + strconv.Itoa(productDetail.Price)
		strData += "<br>Quantity of items bought: " + strconv.Itoa(productDetail.Quantity)
		strData += "<br>Total from Products: Rp" + strconv.Itoa(totalProduct)
		strData += "<br> <img src='" + productPicture + "' alt='" + productDetail.Name + "' width='200' height='300'/>"
	}

	stringBody := "<br>Rekap Harian Tanggal " + time.Now().AddDate(0, 0, -1).Format("02-01-2006")
	stringBody += strData
	stringBody += "<br><br><b>Grand Total : Rp" + strconv.Itoa(grandTotal) + "</b>"

	for i, investor := range investors {
		go func(i int, investor Investor) {
			defer wg.Done()
			m := gomail.NewMessage()
			m.SetHeader("From", "if-21020@students.ithb.ac.id")
			m.SetHeader("To", investor.Email)
			m.SetHeader("Subject", "Rekap Penjualan Harian Fore Cafe")
			m.SetBody("text/html", "Selamat pagi, "+investor.Username+stringBody)

			if err := d.DialAndSend(m); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}(i, investor)

	}
	wg.Wait()
	var response Response
	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = gin.H{"message": "All emails successfully sent"}
	c.JSON(http.StatusOK, response)
}

func CacheProdukGambar() {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT name, pictureurl FROM product")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var picture string
	var name string
	for rows.Next() {
		if err := rows.Scan(&name, &picture); err != nil {
			panic(err)
		}
		err = rdb.Set(ctx, name, picture, 24*time.Hour).Err()
		if err != nil {
			panic(err)
		}
	}
}

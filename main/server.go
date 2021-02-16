package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Fish struct {
	Habitat string `json:"Habitat"`
	Images  struct {
		Src   string `json:"src"`
		Alt   string `json:"alt"`
		Title string `json:"title"`
	} `json:"Species Illustration Photo"`
	Location   string `json:"Location"`
	Population string `json:"Population"`
	Species    string `json:"Species Name"`
	Color      string `json:"Color"`
	Fat        string `json:"Fat, Total"`
	Health     string `json:"Health Benefits"`
	Taste      string `json:"Taste"`
}

func getFish(c *gin.Context) {
	var fish []Fish
	name := c.Query("name")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.fishwatch.gov/api/species/"+name, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &fish)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": fish,
	})
}

func getRandomFish(c *gin.Context) {
	var fish []Fish
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.fishwatch.gov/api/species", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &fish)
	if err != nil {
		fmt.Println(err)
	}

	num1 := rand.Intn(111)
	num2 := rand.Intn(111)
	num3 := rand.Intn(111)
	fishes := []Fish{fish[num1], fish[num2], fish[num3]}
	c.JSON(200, gin.H{
		"message": fishes,
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/api/getRandomFish", getRandomFish)
	router.GET("/api/getFish", getFish)
	router.Run(":3000")
}

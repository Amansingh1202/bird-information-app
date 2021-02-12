package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"io/ioutil"
)


func getFish(c *gin.Context){
	res,err := http.Get("https://www.fishwatch.gov/api/species/red-snapper")
	if err != nil{
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      log.Fatalln(err)
   }
   sb := string(body)
	c.JSON(200, gin.H{
		"message": sb,
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
	router.GET("/api/getAll", getFish)
	router.Run(":3000")
}
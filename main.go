package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itokun12/face-maker/models"
	"github.com/itokun12/face-maker/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Warning: .env file not loading. " + err.Error())
	}
	r := gin.Default()
	models.Initialize()
	r.LoadHTMLGlob("views/parts/*")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/parts")
	})
	r.GET("/parts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Face-Maker",
		})
	})
	r.GET("/parts/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.tmpl", gin.H{
			"title": "Face-Maker",
		})
	})
	r.POST("/parts", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		partsTypeID, err := strconv.Atoi(c.PostForm("parts-type"))
		if err != nil {
			log.Fatal(err)
		}
		filename := header.Filename
		utils.SaveFile(file, filename, "original")
		utils.Resize(filename)
		models.PartsRepository.RegisterPart(filename, partsTypeID)
		c.Redirect(http.StatusMovedPermanently, "/parts")
	})
	r.Run(":" + listenPort())
}

func listenPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}

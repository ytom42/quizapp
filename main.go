package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Quiz struct {
	Id       int
	Question string `gorm:"not null"`
	Answer   string `gorm:"not null"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/resource", "./resource")
	DB, err := gorm.Open(mysql.Open("root:@/quizapp"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// quiz := Quiz{Question: "qqq", Answer: "aaa"}
	// DB.Create(&quiz)
	DB.AutoMigrate(&Quiz{})

	r.GET("/", func(c *gin.Context) {
		quizzes := make([]Quiz, 0)
		DB.Find(&quizzes)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"quizzes": quizzes,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		var quiz Quiz
		quiz.Question = c.PostForm("question")
		quiz.Answer = c.PostForm("answer")
		DB.Create(&quiz)
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":3000")
}

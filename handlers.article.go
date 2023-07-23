package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles, err := getAllArticles()
	if err != nil {
		log.Println("err")
	}

	// Вызов функции рендеринга с именем шаблона для рендеринга
	render(c, gin.H{
		"title":    "Home Page",
		"MyOutput": articles}, "index.html")
}

func showArticleCreationPage(c *gin.Context) {
	// Вызов функции рендеринга с именем шаблона для рендеринга
	render(c, gin.H{
		"title": "Create New NickName"}, "create-NickName.html")
}

func getArticle(c *gin.Context) {
	// Проверяем, действителен ли ID статьи
	if articleID, err := strconv.Atoi(c.Param("NickName_id")); err == nil {
		// Проверяем, существует ли статья
		if article, err := getArticleByID(articleID); err == nil {
			// Вызов функции рендеринга с заголовком, статьей и названием шаблона
			render(c, gin.H{
				"title":    article.Title,
				"MyOutput": article}, "NickName.html")

		} else {
			// Если статья не найдена, прерываем работу с ошибкой
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// Если в URL указан неверный ID статьи, прерываем работу с ошибкой
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {
	// Получаем значения POSTed title и content
	title := c.PostForm("title")
	content := c.PostForm("content")
	username, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
	}

	if a, err := createNewArticle(title, content, username); err == nil {
		// Если статья создана успешно, показать сообщение об успехе
		render(c, gin.H{
			"title":    "Submission Successful",
			"MyOutput": a}, "submission-successful.html")
	} else {
		// если при создании статьи произошла ошибка, прервать с ошибкой
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

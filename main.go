package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Обрабатываем шаблоны в начале, чтобы их не нужно было загружать
	// снова с диска. Это делает обслуживание HTML-страниц очень быстрым.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application   go run .
	router.Run()
}

// Визуализация одного из HTML, JSON или CSV на основе заголовка «Accept» запроса
// Если в заголовке это не указано, отображается HTML при условии, что
// имя шаблона присутствует
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["MyOutput"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["MyOutput"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

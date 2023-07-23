package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Это промежуточное ПО гарантирует, что запрос будет прерван с ошибкой
// если пользователь не авторизован
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Если произошла ошибка или токен пуст
		// пользователь не авторизован
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			//if token, err := c.Cookie("token"); err != nil || token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// Это промежуточное ПО гарантирует, что запрос будет прерван с ошибкой
// если пользователь уже вошел в систему
func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Если ошибки нет или токен не пустой
		// пользователь уже авторизован
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			// if token, err := c.Cookie("token"); err == nil || token != "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// Это промежуточное программное обеспечение устанавливает, вошел ли пользователь в систему или нет
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

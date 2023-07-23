package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
	// Вызов функции рендеринга с именем шаблона для рендеринга
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func performLogin(c *gin.Context) {
	// Получаем значения имени пользователя и пароля, отправленные POST
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Проверяем правильность комбинации логин/пароль
	if isUserValid(username, password) {
		// Если имя пользователя/пароль верно, установите токен в файле cookie
		token := username
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		// Если комбинация имени пользователя и пароля недействительна,
		// показать сообщение об ошибке на странице входа
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func logout(c *gin.Context) {

	// Очистить куки
	c.SetCookie("token", "", -1, "", "", false, true)

	// Перенаправление на главную страницу
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showRegistrationPage(c *gin.Context) {
	// Вызов функции рендеринга с именем шаблона для рендеринга
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	// Получаем значения имени пользователя и пароля, отправленные POST
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		// Если пользователь создан, устанавливаем токен в куки и регистрируем пользователя
		token := username
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		// Если комбинация имени пользователя и пароля недействительна,
		// показать сообщение об ошибке на странице входа
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}

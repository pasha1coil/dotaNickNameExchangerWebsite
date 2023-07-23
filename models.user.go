// models.user.go

package main

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// Для этой демонстрации мы сохраняем список пользователей в памяти
// У нас также есть предопределенные пользователи.
// В реальном приложении этот список, скорее всего, будет выбран
// из базы данных. Более того, в производственных настройках следует
// безопасно хранить пароли, вместо этого добавляя соль и хешируя их
// использовать их, как мы делаем в этой демонстрации
var userList = []user{}

func InsArr(username, password string) ([]user, error) {
	userList = nil

	db, err := sql.Open("mysql", "root:@/dota")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	if err != nil {
		log.Println("err")
	}
	rows, err := db.Query("SELECT Username, Password FROM Users")
	if err != nil {
		log.Println(err)
	}

	// Цикл по строкам,
	// используя Scan для назначения данных столбца полям структуры.
	for rows.Next() {
		var s user
		if err := rows.Scan(&s.Username, &s.Password); err != nil {
			return userList, err
		}
		userList = append(userList, s)
	}
	return userList, nil
}

// Проверяем правильность комбинации имени пользователя и пароля
func isUserValid(username, password string) bool {
	InsArr(username, password)
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// Зарегистрировать нового пользователя с заданным именем пользователя и паролем
// ПРИМЕЧАНИЕ. Для этой демонстрации мы
func registerNewUser(username, password string) (*user, error) {
	InsArr(username, password)
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}
	u := user{Username: username, Password: password}

	db, err := sql.Open("mysql", "root:@/dota")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt := `INSERT INTO Users (username,password)
    VALUES(?, ?)`

	_, err = db.Exec(stmt, username, password)
	if err != nil {
		return nil, errors.New("Не внеслось в БД")
	}

	return &u, nil
}

// Проверяем, доступно ли указанное имя пользователя
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

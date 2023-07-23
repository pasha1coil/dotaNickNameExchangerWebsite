// models.article.go

package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	Postedid int    `json:"Postedid"`
	Username string `json:"Username"`
	Title    string `json:"Title"`
	Content  string `json:"Content"`
}

// Для этой демонстрации мы сохраняем список статей в памяти
// В реальном приложении этот список, скорее всего, будет выбран
// из базы данных или из статических файлов
var articleList = []article{}

// Возвращаем список всех статей
func getAllArticles() ([]article, error) {

	db, err := sql.Open("mysql", "root:@/dota")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT Postedid, Username, Title, Content FROM Posted")
	if err != nil {
		log.Println(err)
	}
	articleList = nil
	// Цикл по строкам,
	// используя Scan для назначения данных столбца полям структуры.
	for rows.Next() {
		var s article
		if err := rows.Scan(&s.Postedid, &s.Username, &s.Title, &s.Content); err != nil {
			return articleList, err
		}
		articleList = append(articleList, s)
	}

	return articleList, nil
}

// Получить статью на основе предоставленного идентификатора
func getArticleByID(id int) (*article, error) {
	getAllArticles()
	for _, a := range articleList {
		if a.Postedid == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Создать новую статью с указанным заголовком и содержимым
func createNewArticle(title, content, username string) (*article, error) {
	db, err := sql.Open("mysql", "root:@/dota")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt := `INSERT INTO Posted (username,title, content)
    VALUES(?, ?, ?)`

	_, err = db.Exec(stmt, username, title, content)
	if err != nil {
		return nil, errors.New("Не внеслось в БД")
	}

	// Устанавливаем ID новой статьи на единицу больше, чем количество статей
	a := article{Postedid: len(articleList) + 1, Title: title, Content: content}

	return &a, nil
}

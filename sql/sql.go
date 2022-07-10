package sql

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Post struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var content string

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=user dbname=sample password=*** sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		post := Post{}
		posts := []*Post{}

		if err := Db.QueryRow("select id, name, content", id).Scan(&post.Id, &post.Name, &post.Content); err != nil {
			return errors.Wrapf(err, "connot connect SQL")
		}

		posts = append(posts, &Post{Id: post.Id, Name: post.Name, Content: post.Content})
		return c.JSON(http.StatusOK, posts)
	}
}

func GetPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		post := Post{}
		posts := []*Post{}

		rows, err := Db.Query("select id, name, content")
		if err != nil {
			return errors.Wrapf(err, "connot connect SQL")
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&post.Id, &post.Name, &post.Content); err != nil {
				return errors.Wrapf(err, "connot connect SQL")
			}
			posts = append(posts, &Post{Id: post.Id, Name: post.Name, Content: post.Content})
		}

		return c.JSON(http.StatusOK, posts)

	}
}

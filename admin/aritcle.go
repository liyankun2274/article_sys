package admin

import (
	"Article_sys/modules"
	"Article_sys/utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type AritcleParmas struct {
	Title    string `json:"title" form:"title" query:"title"`
	Content  string `json:"content" form:"content" query:"content"`
	ColumnId int    `json:"column_id" form:"column_id" query:"column_id"`
	UserId   int    `json:"user_id" form:"user_id" query:"user_id"`
	Author   string `json:"author" form:"author" query:"author"`
	Status   int8   `json:"int" form:"int" query:"int"`
}

func New_article(e *echo.Group, db *gorm.DB) error {
	e.POST("/add_article", func(context echo.Context) error {
		articleP := new(AritcleParmas)
		if err := context.Bind(articleP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		new_article := modules.Article{Title: articleP.Title, Content: articleP.Content, ColumnID: articleP.ColumnId, UserId: articleP.UserId, Author: articleP.Author, Status: 1, CreateAt: time.Now()}
		db.NewRecord(&new_article)
		db.Create(&new_article)
		return utils.Ok(context, nil)

	})
	e.DELETE("/del_article/:id", func(context echo.Context) error {
		id, _ := strconv.Atoi(context.Param("id"))
		article := &modules.Article{ID: id}
		db.Delete(&article)
		return utils.Ok(context, nil)
	})
	e.PUT("/modify_article/:id", func(context echo.Context) error {
		id, _ := strconv.Atoi(context.Param("id"))
		articleP := new(AritcleParmas)
		if err := context.Bind(articleP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		article := modules.Article{ID: id, Title: articleP.Title, Content: articleP.Content, ColumnID: articleP.ColumnId, UserId: articleP.UserId, Author: articleP.Author, Status: articleP.Status, UpdateAt: time.Now()}
		db.Model(&article).Updates(map[string]interface{}{"title": article.Title, "content": article.Content, "column_id": article.ColumnID, "user_id": article.UserId, "author": article.Author, "status": article.Status, "update_at": article.UpdateAt})
		return utils.Ok(context, nil)
	})
	e.GET("/article", func(context echo.Context) error {
		var article []modules.Article
		db.Find(&article)
		return utils.Ok(context, article)
	})
	e.GET("/article/:id", func(context echo.Context) error {
		var article modules.Article
		id, _ := strconv.Atoi(context.Param("id"))
		db.First(&article, id)
		if article.ID == 0 {
			return utils.Ok(context, nil)
		}
		return utils.Ok(context, article)
	})
	e.GET("article_bycond", func(context echo.Context) error {
		articleP := new(AritcleParmas)
		if err := context.Bind(articleP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		article := modules.Article{Title: articleP.Title, Content: articleP.Content, ColumnID: articleP.ColumnId, UserId: articleP.UserId, Author: articleP.Author, Status: 1, CreateAt: time.Now()}
		db.Where(map[string]interface{}{"title": article.Title, "content": article.Content, "column_id": article.ColumnID, "user_id": article.UserId, "author": article.Author, "status": article.Status, "update_at": article.UpdateAt}).Find(&article)
		return utils.Ok(context, article)
	})

	return nil
}

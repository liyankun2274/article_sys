package admin

import (
	"Article_sys/modules"
	"Article_sys/utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type ColumnParmas struct {
	Name string `json:"name" form:"name" query:"name"`
}

func New_column(e *echo.Group, db *gorm.DB) error {
	e.POST("/add_column", func(context echo.Context) error {
		columnP := new(ColumnParmas)
		if err := context.Bind(columnP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		new_colum := modules.Column{Name: columnP.Name, Status: 1, CreateAt: time.Now()}
		db.NewRecord(&new_colum)
		db.Create(&new_colum)
		return utils.Ok(context, nil)

	})
	e.DELETE("/del_column/:id", func(context echo.Context) error {
		id, _ := strconv.Atoi(context.Param("id"))
		column := &modules.Column{ID: id}
		db.Delete(&column)
		return utils.Ok(context, nil)
	})
	e.PUT("/modify_column/:id", func(context echo.Context) error {
		id, _ := strconv.Atoi(context.Param("id"))
		columnP := new(ColumnParmas)
		if err := context.Bind(columnP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		column := &modules.Column{ID: id, Name: columnP.Name}
		db.Model(&column).Update("name", column.Name)

		return utils.Ok(context, nil)
	})
	e.GET("/column", func(context echo.Context) error {
		var column []modules.Column
		db.Find(&column)
		return utils.Ok(context, column)
	})
	e.GET("/column/:id", func(context echo.Context) error {
		var column modules.Column
		id, _ := strconv.Atoi(context.Param("id"))
		db.First(&column, id)
		if column.ID == 0 {
			return utils.Ok(context, nil)
		}
		return utils.Ok(context, column)
	})

	return nil
}

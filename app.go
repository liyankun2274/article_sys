package main

import (
	"Article_sys/admin"
	"Article_sys/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
)

func main() {
	/*
	加载配置文件
	 */
	config, err := LoadConfig("config.toml")
	if err != nil {
		fmt.Println("load config file has a error: ", err.Error())
		return
	}
	/*
	初始化mysql
	 */
	db, err := gorm.Open("mysql", config.MysqlDb)
	//全局禁用表明为复数形式
	db.SingularTable(true)
	if err != nil {
		fmt.Println("connect database: ", err.Error())
		return
	}
	//打印sql语句
	db.LogMode(true)
	defer db.Close()
	/*
	连接redis
	 */
	redis_pool := utils.NewRedisPool(config.RedisUrl)
	redis_session, err := utils.GetRedisSession(redis_pool, config.RedisSessionPrefix)
	if err != nil {
		fmt.Println("connect redis: ", err.Error())
		return
	}
	//实例化 echo
	var e = echo.New()
	//静态资源
	e.Static("/", "assets")
	/*
	加载模板引擎
	 */
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	/*
	加载echo http
	 */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(redis_session)
	e.Use(Is_login)
	admin.New_user(e.Group("/api/admin"), db)
	admin.New_column(e.Group("/api/admin"), db)
	admin.New_article(e.Group("/api/admin"), db)
	e.Logger.Fatal(e.Start(config.Port))
}

//模板引擎
// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

//中间件函数(验证session)
func Is_login(e echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Path()
		login_api := utils.Substr(url,11,16)
		if login_api == "login"{
			return e(c)
		}
		s, _ := session.Get("session", c)
		if s.Values["user_id"] == nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		} else {
			return e(c)
		}
		//执行下一个中间件或者执行控制器函数, 然后返回执行结果
	}
}

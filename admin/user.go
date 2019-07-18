package admin

import (
	"Article_sys/modules"
	"Article_sys/utils"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type RegParams struct {
	Mobile   string `json:"mobile" form:"mobile" query:"mobile"`
	Password string `json:"password" form:"password" query:"password"`
	Nickname string `json:"nickname" form:"nickname" query:"nickname"`
}

type LoginParams struct {
	Mobile   string `json:"mobile" form:"mobile" query:"mobile"`
	Password string `json:"password" form:"password" query:"password"`
}


func New_user(e *echo.Group, db *gorm.DB) error {
	e.GET("/login_view", func(context echo.Context) error {
		//渲染hello.html模版文件，模版参数为world
		return context.Render(200, "login.html", "")
	})
	e.POST("/reg", func(context echo.Context) error {
		regP := new(RegParams)
		if err := context.Bind(regP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		var user modules.User
		db.Where("mobile = ?", regP.Mobile).First(&user)
		if user.ID != 0 {
			return utils.Fail(context, "1000", "该账号已经被注册!")
		}
		var new_user = modules.User{Mobile: regP.Mobile, Password: regP.Password, Nickname: regP.Nickname, Status: 1}
		db.NewRecord(&new_user)
		db.Create(&new_user)
		return utils.Ok(context, nil)
	})
	e.POST("/login", func(context echo.Context) error {
		loginP:=new(LoginParams)
		if err := context.Bind(loginP); err != nil {
			return utils.Fail(context, "-1000", err.Error())
		}
		var user modules.User
		db.Where(&modules.User{Mobile: loginP.Mobile, Password: loginP.Password}).First(&user)

		if user.ID == 0 {
			return utils.Fail(context, "1000", "请输入正确的用户名或密码!")
		}
		s, _ := session.Get("session", context)
		s.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		s.Values["user_id"] = user.ID
		s.Values["user_nickname"] = user.Nickname
		if err := s.Save(context.Request(), context.Response()); err != nil {
			fmt.Println("session save: ", err)
		}
		return utils.Ok(context, nil)
	})
	// 登出
	e.PUT("/logout", func(context echo.Context) error {
		s, _ := session.Get("session", context)
		s.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}
		s.Save(context.Request(), context.Response())
		return utils.Ok(context, nil)
	})
	return nil
}

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smf8/http-challange/db"
	"github.com/smf8/http-challange/handler"
	"github.com/smf8/http-challange/middleware"
	"github.com/smf8/http-challange/model"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	d := db.New()
	d.AutoMigrate(&model.Student{})
	err := model.LoadStudents("data/CEIT98List.txt", d)
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(d)
	g := e.Group("/api")
	g.GET("/", h.MainPage)
	g.GET("/get", h.HandleGET)          //sid ro to header bede
	g.GET("/param", h.HandleQueryParam) // sid va name ro ba query param bede
	g.POST("/post", h.HandlePOST)       // sid va password ro besoorat form-data POST kone
	v := middleware.New(d)
	g1 := g.Group("/private", v.ValidateHeader)
	g1.POST("", h.HeaderHandler)      // POST khali be /private ba header token ke to marhale POST migire
	g1.GET("/cookie", h.CookieSetter) // GET be /private/Cookie ke dobare bayad token dashte bashe

	g.POST("/decode", h.DecryptKey)
	return e
}

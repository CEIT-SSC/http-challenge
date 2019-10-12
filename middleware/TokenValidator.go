package middleware

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/smf8/http-challange/model"
	"net/http"
	"strings"
)

type Validator struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Validator {
	return &Validator{db: db}
}
func (v *Validator) ValidateHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		b, _ := base64.URLEncoding.DecodeString(token)
		token = strings.Split(string(b), "*")[0]
		s := model.Student{}
		v.db.Where("sid = ?", token).First(&s)
		if s.Sid != 0 {
			c.Set("sid", s.Sid)
			return next(c)
		} else {
			c.String(http.StatusUnauthorized, "توکن ات معیوبه. یعنی چی آخه ؟ فک کردی من خرم ؟ اگه اینقدر خر بودم که اینهمه اطلاعات شما رو پیش خودم نگه نمیداشتم.")
			return nil
		}
	}
}

package handler

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/smf8/http-challange/model"
	"github.com/smf8/http-challange/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(d *gorm.DB) *Handler {
	return &Handler{db: d}
}

func (h *Handler) MainPage(c echo.Context) error {
	return c.HTML(http.StatusOK, `<html dir="rtl"><body><p>`+model.MainPageResponse+`</p></body></html>`)
}

// in this step user must send a GET request to /param with query params "sid" and "name"
func (h *Handler) HandleQueryParam(c echo.Context) error {
	//handling query parameters
	name, sid := c.QueryParam("name"), c.QueryParam("sid")
	name, sid = strings.TrimSpace(name), strings.TrimSpace(sid)
	client := model.DetectClient(c.Request().UserAgent())
	fmt.Println(client + "   khd")
	if name != "" && sid != "" {
		num, _ := strconv.Atoi(sid)
		s := model.Student{}
		h.db.Model(&s).Where("sid = ?", num).First(&s)
		if s.Sid != 0 {
			// giving out 1/4 of the key
			h.db.Model(&s).Update("key", s.FinalKey[:len(s.FinalKey)/4])
			return c.HTML(http.StatusOK, `<html dir="rtl"><body>Save this text somewhere : `+s.FinalKey[:15]+`<p>`+model.SuccessQPMessage(name, "")+`</p></body></html>`)
		} else {
			return c.String(http.StatusBadRequest, model.FailedQueryParam)
		}
	}
	return c.String(http.StatusBadRequest, model.FailedQueryParam)
}

// in this step user must send a get request to / with "sid" header
func (h *Handler) HandleGET(c echo.Context) error {
	// header handling
	header := c.Request().Header.Get("sid")
	// check if header is a valid integer
	if num, err := strconv.Atoi(header); err == nil && len(header) == 7 {
		// non-empty header
		s := model.Student{}
		h.db.Model(&s).Where("sid = ?", num).First(&s)
		if s.Sid != 0 {
			return c.HTML(http.StatusOK, model.GreetingGetMessage(s.FirstName))
		} else {
			//Handle if user is not defined
			return c.HTML(http.StatusBadRequest, `<html dir="rtl"><body><p>`+model.FailedGetMessage+`</p></body></html>`)
		}
	} else {
		return c.HTML(http.StatusBadRequest, `<html dir="rtl"><body><p>`+model.FailedGetMessage+`</p></body></html>`)
	}
}

// handling POST requests to /post with "sid" and "password" form values.
func (h *Handler) HandlePOST(c echo.Context) error {
	sid := strings.TrimSpace(c.FormValue("sid"))
	psswd := c.FormValue("password")
	// checking if posted data are valid
	if _, err := strconv.Atoi(sid); err != nil || len(sid) != 7 || psswd == "" {
		return c.HTML(http.StatusBadRequest, `<html dir="rtl"><body><p>`+model.FailedPassword+`</p></body></html>`)
	}
	s := model.Student{}
	h.db.Model(&s).Where("sid = ?", sid).First(&s)
	if s.Sid != 0 {
		md := fmt.Sprintf("%x", md5.Sum([]byte(psswd)))
		// Validating student's password
		if md == s.Pass {
			h.db.Model(&s).Update("key", s.FinalKey[:len(s.FinalKey)/2])
		} else {
			// invalid password. returning 401
			return c.String(http.StatusUnauthorized, "password et ghalate, passwordi ke moghe jash behet dadim ro befrest")
		}
		token := base64.URLEncoding.EncodeToString([]byte(sid + "*"))
		c.Response().Header().Set("token", token)
		// giving out 2/4 of the key
		return c.String(http.StatusOK, "It might be of use someday, Save it. "+s.FinalKey[len(s.FinalKey)/4:len(s.FinalKey)/2]+"\n\n"+model.SuccessPOST)
	} else {
		// student not found
		return c.String(http.StatusBadRequest, model.FailedQueryParam)
	}
}

// handling GET request with key param. useless as there is already HandleQueryParam to handel query parameters
func (h *Handler) KeyHandler(c echo.Context) error {
	sid := c.Param("key")
	if len(sid) != 5 {
		return c.String(http.StatusBadRequest, "key ات نامعتبره برادر / خواهر")
	}
	s := model.Student{}
	h.db.Where("key LIKE ?", "%"+sid+"%").First(&s)
	if s.Sid != 0 {
		fmt.Println("Found User Hurray")
		key := md5.Sum([]byte("GET" + strconv.Itoa(s.Sid)))
		c.Response().Header().Set("CEIT-TOKEN", fmt.Sprintf("%x", key))
		h.db.Model(&s).Update("key", s.KEY+fmt.Sprintf("%x-", key))
		return c.String(http.StatusOK, `خب اینجا هم کارت به خوبی و درستی انجام شد. حالا وقتشه که یه چیز جدید یاد بگیری. Header ها :| تو این چیزی که من دارم بهت نشون میدم. علاوه بر این متن. یه چیز میزای دیگه ای هم هست که بهش میگن Response Header. از تو این برو یه گوگولی ای پیدا کن به دردت میخوره.

گوگولیه رو که پیدا کردی. گوگولیه پارت دوم پس کدت هست. نگهش دار برا روز مبادا.
راستی برای ادامه ماجرا. باید برای همه Request هایی که میفرستی هم Header بذاری. اگه علاقه داری تحقیق کن که Request Header ها چی ان. اگه هم علاقه نداری صرفا با Postman قسمت header یه هدر به اسم token و مقدار چیزی که اینجا پیدا کردی بفرست...`)
	} else {
		return c.String(http.StatusBadRequest, "یا میخوای سر منو کلاه بذاری و داری چرت و پرت میفرستی یا اینکه مرحله قبلی رو درست انجام ندادی. در هر صورت برو یه گوشه و به کار هات فکر کن")
	}
}

// Cookie setter is the last step of the challenge which sends the final part of the key as a cookie to the client
func (h *Handler) CookieSetter(c echo.Context) error {
	if sid := c.Get("sid"); sid != nil {
		// request is valid
		s := model.Student{}
		fmt.Println("fds")
		h.db.Where("sid = ?", sid).First(&s)
		h.db.Model(&s).Update("key", s.FinalKey)
		cookie := new(http.Cookie)
		cookie.Name = "pass"
		// giving out 4/4 of the key
		cookie.Value = s.FinalKey[3*len(s.FinalKey)/4:]
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.String(http.StatusOK, model.SuccessCookie)
	} else {
		return c.String(http.StatusBadRequest, "vala yeja ro fek konam eshtebah kardi. boro be eshtebahatet fekr kon")
	}
}
func (h *Handler) HeaderHandler(c echo.Context) error {
	if sid := c.Get("sid"); sid != nil {
		// request is valid
		s := model.Student{}
		h.db.Where("sid = ?", sid).First(&s)
		h.db.Model(&s).Update("key", s.FinalKey[:3*len(s.FinalKey)/4])
		// giving out 3/4 of the key
		return c.String(http.StatusOK, "Your final peace of shit... oh sorry, your final piece of code!!!: "+s.FinalKey[len(s.FinalKey)/2:3*len(s.FinalKey)/4]+"\n\n"+model.SuccessAuth)
	} else {
		return c.String(http.StatusBadRequest, "vala yeja ro fek konam eshtebah kardi. boro be eshtebahatet fekr kon")
	}
}

func (h *Handler) DecryptKey(c echo.Context) error {
	msg := c.FormValue("message")
	key := c.FormValue("key")
	if key == "Welcome2CE98:):)" {
		s := model.Student{}
		h.db.Where("final_key = ?", msg).First(&s)
		if s.Sid != 0 {
			msg, err := utils.Decrypt([]byte(key), msg)
			if err != nil {
				return err
			}
			return c.String(http.StatusOK, msg)
		} else {
			// msg ye chiz valid nist. cherto perte...
			return c.String(http.StatusBadRequest, "Aghaaaaaa. mano bikar gir Avordi ? midooni man chanta request ro bayad javab bedam ? bad to dari payam chert bara man mifresti man decode conam ? midooni har payami ke midi cheghadr vaght mano migire ?\nreAyat kon dige ghaaa")
		}
	} else {
		return c.String(http.StatusBadRequest, "key't ghalate baradaram / khaharam. ye chizi labela chert o pertaye jashn behet goftimaaaa. key az beyne hamoonas. age hamchenan nemidooni. boro az bache haye anjoman bepors.")
	}
}

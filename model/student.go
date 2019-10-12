package model

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smf8/http-challange/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

type Student struct {
	gorm.Model
	Sid       int `gorm:"unique;not null"`
	FirstName string
	LastName  string
	Pass      string
	KEY       string
	FinalKey  string
}

func New(sid int) *Student {
	smap := []string{"p", "q", "w", "e", "r", "t", "y", "u", "i", "o"}
	pass := ""
	d := sid
	for d > 0 {
		pass += smap[d%10]
		d = d / 10
	}
	pass = fmt.Sprintf("%x", md5.Sum([]byte(pass)))
	return &Student{Sid: sid,
		Pass: pass}
}

func LoadStudents(filePath string, db *gorm.DB) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(bytes.NewReader(file))
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		lineS := strings.Split(line, ":")
		sid, err := strconv.Atoi(strings.TrimSpace(lineS[0]))
		if err != nil {
			return err
		}
		s := New(sid)
		name := strings.Split(lineS[1], "-")
		s.FirstName = name[0]
		s.LastName = name[1]
		key := []byte("Welcome2CE98:):)")
		encryptMsg, _ := utils.Encrypt(key, "Salaaaam bache ye 98 i e gol, khoobi ? Aaaaaafarin... Saaaad Afarin. Hezaaaaro sisad Afarinnnnnn. hala ke inghadr chiz yad gerefti. bia to anjoman ma ham ye 2 seta jayeze behet bedim. faghat movazeb bash. ke ma hammmeeee chi ro midoonim. pas say nakon sar ma kolah bezari aghaye  "+name[0]+" "+name[1])
		s.FinalKey = encryptMsg
		db.Create(s)
	}
	return nil
}

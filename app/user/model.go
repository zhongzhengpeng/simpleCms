package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"simpleCms/app/common"

	"gorm.io/gorm"
)

type Work struct {
	Username  string       `gorm:"type:varchar(255);uniquendex" json:"username"`
	Password  string       `gorm:"type:varchar(255)" json:"password"`
	Token     string       `gorm:"type:varchar(4096)" json:"token"`
	TokenTime sql.NullTime `gorm:"" json:"token_time"`
	Name      string       `gorm:"type:varchar(255);index" json:"name"`
}

type User struct {
	gorm.Model
	Work
}

func init() {
	if err := common.DB.AutoMigrate(&User{}); err != nil {
		log.Fatalln(err)
	}
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func password(plain string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(plain))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

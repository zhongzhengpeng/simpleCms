package content

import (
	"log"
	"simpleCms/app/common"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Work
}

type Work struct {
	Subject string `gorm:"type:varchar(255);index" json:"subject"`
	Summary string `gorm:"type:varchar(255)" json:"summary"`
	Body    string `gorm:"type:text" json:"body"`
	UserID  uint   `gorm:"index" json:"user_id"`
	Status  string `gorm:"type:enum('new','draft','publish')" json:"status"`
}

func init() {
	if err := common.DB.AutoMigrate(&Content{}); err != nil {
		log.Println(err)
	}
}

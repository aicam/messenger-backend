package internal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Messages struct {
	gorm.Model
	SrcUsername  string `json:"src_username"`
	DestUsername string `json:"dest_username"`
	Text         string `json:"text"`
}

type FilesData struct {
	gorm.Model
	SrcUsername  string `json:"src_username"`
	DestUsername string `json:"dest_username"`
	FileName     string `json:"file_name"`
	Data         string `json:"data" gorm:"type:varchar(8192)"`
	Chunk        int    `json:"chunk"`
	HMAC         string `json:"hmac"`
	Key          string `json:"key"`
}

type Users struct {
	gorm.Model
	Username            string `json:"username"`
	Password            string `json:"password"`
	IdentificationToken string `json:"identification_token"`
}

func MakeMigrations(connectionString string) *gorm.DB {
	var db *gorm.DB
	var err error
	retryCount := 10
	for {
		db, err = gorm.Open("mysql", connectionString)
		if err != nil {
			log.Print("mysql connection error : ", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
		if retryCount == 0 {
			log.Print("maximum retry reached")
			break
		}
		retryCount--
	}
	db.AutoMigrate(&FilesData{})
	db.AutoMigrate(&Messages{})
	return db
}

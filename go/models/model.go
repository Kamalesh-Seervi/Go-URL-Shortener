package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Url struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Url      string `json:"url" gorm:"unique"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	conn := "host=172.23.0.2 port=5432 user=root password=changeme dbname=mydb sslmode=enable"
	fmt.Println(conn)
	var err error

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Url{})
	if err != nil {
		fmt.Println(err)
	}
}

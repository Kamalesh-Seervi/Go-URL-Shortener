package models

import "errors"

func GetAllUrl() ([]Url, error) {

	var urltable []Url
	tx := db.Find(&urltable)

	if tx.Error != nil {
		panic(tx.Error)
	}

	return urltable, nil
}

func GetOneUrl(id uint64) (Url, error) {
	var url Url

	tx := db.Where("id = ?", id).First(&url)
	if tx.Error != nil {
		return Url{}, tx.Error
	}

	if url.ID == 0 {
		return Url{}, errors.New("URL not found")
	}

	return url, nil
}

func CreateUrl(url Url) error {
	tx := db.Create(&url)
	return tx.Error
}

func UpdateUrl(url Url) error {
	tx := db.Save(&url)
	return tx.Error
}

func DeleteUrl(id uint64) error {
	tx := db.Unscoped().Delete(&Url{}, id)
	return tx.Error

}

func FindByGolyUrl(urll string) (Url, error) {
	var url Url
	tx := db.Where("url = ?", urll).First(&url)
	return url, tx.Error
}

package models

func GetAllUrl() ([]Url, error) {

	var urls []Url
	tx := db.Find(&urls)

	if tx.Error != nil {
		panic(tx.Error)
	}

	return urls, nil
}

func GetOneUrl(id uint64) (Url ,error) {
	var url Url

	tx := db.Where("id = ?", id).Find(&url)
	if tx.Error != nil {
		return Url{}, tx.Error
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

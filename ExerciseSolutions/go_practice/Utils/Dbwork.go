package Utils

import (
	"practice-bootcamp/Config"
	"practice-bootcamp/Models"
)

func IsPresent(id string, book *Models.Book) error {
	return Config.DB.Where("id = ?", id).First(book).Error
}

func DoCreate(book *Models.Book) {
	Config.DB.Create(book)
}

func DoFind(book *[]Models.Book) {
	Config.DB.Find(book)
}

func DoUpdate(book *Models.Book, newBook Models.UpdatedBook) {
	Config.DB.Model(book).Updates(newBook)
}

func DoDelete(book *Models.Book) error {
	return Config.DB.Delete(book).Error
}

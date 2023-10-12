package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
//The gorm.Model field is used to embed common fields that are often required 
//when working with a database, such as an ID, created timestamp, updated timestamp, and deleted timestamp.
//DeletedAt: A timestamp indicating when the record was soft-deleted, if applicable.
// Soft deletion is a technique where records are not physically deleted from the database but marked as deleted with a timestamp.
	Id string 
	Title string
	Author string
	Price  int
}

//create a user
func CreateUser(db *gorm.DB, book *Book) (err error) {
	err = db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, book *[]Book) (err error) {
	err = db.Find(book).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, book *Book, id string) (err error) {
	err = db.Where("id = ?", id).First(book).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, book *Book) (err error) {
	db.Save(book)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, book *Book, id string) (err error) {
	db.Where("id = ?", id).Delete(book)
	
	return nil
}

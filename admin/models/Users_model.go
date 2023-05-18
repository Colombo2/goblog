package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User_model struct {
	gorm.Model
	Username, Password string
}

func (user User_model) Migrate() {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&user)
}

func (user User_model) Add() {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&user)
}

func (user User_model) Get(where ...interface{}) (Post User_model) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}
	db.First(&user, where...)
	return user
}

func (user User_model) Get_All(where ...interface{}) (Post []User_model) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts_list []User_model
	db.Find(&posts_list, where...)
	return posts_list
}

func (user User_model) Update(column string, value interface{}) (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Model(&user).Update(column, value).Error; err == nil {
		return "Error"
	} else {
		return "OK"
	}
}

func (user User_model) Updates(data User_model) (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Model(&user).Updates(data).Error; err != nil {
		return "Error"
	} else {
		return "OK"
	}
}

func (user User_model) Delete() (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Delete(&user, user.ID).Error; err != nil {
		return "Error"
	} else {
		return "OK"
	}
}

package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) Migrate() {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&category)
}

func (category Category) Add() {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&category)
}

func (category Category) Get(where ...interface{}) (Post Category) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return category
	}
	db.First(&category, where...)
	return category
}

func (category Category) Get_All(where ...interface{}) (Post []Category) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts_list []Category
	db.Find(&posts_list, where...)
	return posts_list
}

func (category Category) Update(column string, value interface{}) (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Model(&category).Update(column, value).Error; err == nil {
		return "Error"
	} else {
		return "OK"
	}
}

func (category Category) Updates(data Category) (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Model(&category).Updates(data).Error; err != nil {
		return "Error"
	} else {
		return "OK"
	}
}

func (category Category) Delete() (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Delete(&category, category.ID).Error; err != nil {
		return "Error"
	} else {
		return "OK"
	}
}

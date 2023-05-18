package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post_model struct {
	gorm.Model
	Title, Slug, Description, Content, Pic_url string
	CategoryID                                 int
}

func (post Post_model) Migrate() {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&post)
}

func (post Post_model) Add() {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&post)
}

func (post Post_model) Get(where ...interface{}) (Post Post_model) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}
	db.First(&post, where...)
	return post
}

func (post Post_model) Get_All(where ...interface{}) (Post []Post_model) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts_list []Post_model
	db.Find(&posts_list, where...)
	return posts_list
}

func (post Post_model) Update(column string, value interface{}) (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Model(&post).Update(column, value).Error; err == nil {
		return "Error"
	} else {
		return "OK"
	}
}

func (post Post_model) Updates(data Post_model) (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Model(&post).Updates(data).Error; err != nil {
		return "Error"
	} else {
		return "OK"
	}
}

func (post Post_model) Delete() (err0r string) {
	db, err := gorm.Open(mysql.Open(db_c), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Delete(&post, post.ID).Error; err != nil {
		return "Error"
	} else {
		return "OK"
	}
}

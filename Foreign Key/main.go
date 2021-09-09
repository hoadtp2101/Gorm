package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Cate_id  uint
	Category Category `gorm:"foreignKey:Cate_id"`
}

type Category struct {
	gorm.Model
	Id   uint
	Name string
}

func main() {
	// Kết nối database
	dsn := "root:hoa@123@/testgo?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Tạo table
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})

	// Thêm mới
	cate := []Category{{Name: "Áo"}, {Name: "Quần"}, {Name: "Váy"}}
	// db.Create(cate)

	pro := []Product{{Name: "Áo ba lỗ", Cate_id: 1}, {Name: "Áo thun", Cate_id: 1}, {Name: "Quần đùi", Cate_id: 2}, {Name: "Váy công chúa", Cate_id: 3}}
	// db.Create(pro)

	for _, v := range cate {
		fmt.Println(v.ID)
		fmt.Println(v.Name)
		fmt.Println("===================")
	}
	fmt.Println("******************")
	for _, v := range pro {
		fmt.Println(v.ID)
		fmt.Println(v.Name)
		fmt.Println(v.Cate_id)
		fmt.Println("===================")
	}
}

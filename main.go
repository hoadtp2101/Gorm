package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:hoa@123@/testlab?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	// c1:
	// db.Create(&Product{Code: "H01", Price: 210})

	// c2:
	// pro := Product{Code: "H02", Price: 2000}
	// db.Create(&pro)

	// c3
	// var pro = []Product{{Code: "H03", Price: 2400}, {Code: "H04", Price: 2300}}
	// db.Create(&pro)

	// UPDATE
	// pro := Product{}
	// db.Find(&pro, 6)
	// pro.Price = 3000
	// db.Save(&pro)

	// UPDATE với where, update ....

	// delete
	// var product Product
	// db.Where("price = ?", "2000").Delete(&product)          // Xoá mềm
	// db.Unscoped().Where("price = ?", "200").Delete(&product) // xoá vĩnh viễn
}

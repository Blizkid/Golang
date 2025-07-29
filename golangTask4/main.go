package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func runDb() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	//run()
	RunBlog(db)

	//accounts := []*Account{
	//	{ID: 1, Balance: 180},
	//	{ID: 2, Balance: 190},
	//}
	//db.Create(accounts)

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	//users := []*User{
	//	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	//	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	//}

	//students := []*Student{
	//	{Name: "Jinzhu", Age: 18, Grade: "高一"},
	//	{Name: "Jackson", Age: 19, Grade: "高三"},
	//	{Name: "Zhangsan", Age: 15, Grade: "初二"},
	//	{Name: "Lisi", Age: 14, Grade: "初一"},
	//	{Name: "Wangwu", Age: 13, Grade: "初一"},
	//	{Name: "Zhaoliu", Age: 18, Grade: "高二"},
	//}

	//result := db.Create(users) // 通过数据的指针来创建
	//result := db.Create(students)
	//fmt.Println(result.RowsAffected)

	//db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	//var students []Student
	//db.Where("age >= ?", 15).Find(&students)
	//fmt.Println(students) // 或者自己for循环格式化输出

	//var students []Student

	//db.Model(&Student{}).Where("name = ?", "张三").Update("Grade", "四年级")
	//db.Where("age < ?", 15).Delete(&Student{Name: "张三"})
	//
	//fmt.Println(students) // 或者自己for循环格式化输出
	// 开始事务
	//tx := db.Begin()
	//
	//var accountA, accountB Account
	//if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&accountA, 1).Error; err != nil {
	//	tx.Rollback()
	//	return
	//}
	//if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&accountB, 2).Error; err != nil {
	//	tx.Rollback()
	//	return
	//}
	//
	//if accountA.Balance >= 100 {
	//	// 先减A
	//	if err := tx.Model(&accountA).Update("Balance", accountA.Balance-100).Error; err != nil {
	//		tx.Rollback()
	//		return
	//	}
	//	// 再加B
	//	if err := tx.Model(&accountB).Update("Balance", accountB.Balance+100).Error; err != nil {
	//		tx.Rollback()
	//		return
	//	}
	//	// 记流水
	//	if err := tx.Create(&Transaction{
	//		FromAccountID: accountA.ID,
	//		ToAccountID:   accountB.ID,
	//		Amount:        100,
	//	}).Error; err != nil {
	//		tx.Rollback()
	//		return
	//	}
	//	tx.Commit()
	//} else {
	//	tx.Rollback()
	//}

}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//type User struct {
//	gorm.Model
//	Name     string
//	Age      uint
//	Birthday time.Time
//}

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   uint
	Grade string
}

type Account struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"not null;default:0"`
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountID uint    `gorm:"not null"` // 这里指向 Account.ID
	ToAccountID   uint    `gorm:"not null"` // 这里也指向 Account.ID
	Amount        float64 `gorm:"not null"`

	// 结构体关联（可选，便于加载详细信息）
	FromAccount Account `gorm:"foreignKey:FromAccountID;references:ID"`
	ToAccount   Account `gorm:"foreignKey:ToAccountID;references:ID"`
}

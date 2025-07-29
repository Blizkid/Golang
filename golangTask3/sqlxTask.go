package golangTask3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func run() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?parseTime=true"
	db, _ := sqlx.Open("mysql", dsn)

	var employeesTe []Employee
	db.Select(&employeesTe, "SELECT * FROM employees WHERE department = ?", "技术部")
	fmt.Println(employeesTe)

	var maxSalaryEmployee Employee
	db.Get(&maxSalaryEmployee, "SELECT * FROM employees order by salary desc limit 1")
	fmt.Println(maxSalaryEmployee)

	books := []Book{
		{1, "我的奋斗", "希特勒", 100.0},
		{2, "活着", "余华", 80.0},
		{3, "战争的艺术", "麦克阿瑟", 90.0},
		{4, "时间简史", "霍金", 12.0},
	}

	//db.Exec(`CREATE TABLE IF NOT EXISTS books (
	//    id INTEGER PRIMARY KEY,
	//    title TEXT,
	//    author TEXT,
	//    price REAL
	//)`)

	for _, boo := range books {
		res, err := db.NamedExec(`INSERT INTO books (id, title, author, price)
           VALUES (:id, :title, :author, :price)`, &boo)
		if err != nil {
			fmt.Printf("插入失败: %v\n", err)
		}
		affect, _ := res.RowsAffected()
		fmt.Printf("插入影响行数: %v\n", affect)
	}

	var selectBooks []Book
	db.Select(&selectBooks, "SELECT * FROM books WHERE price > ?", 50)
	fmt.Println(selectBooks)
}

type Employee struct {
	Id         uint64 `gorm:"primary_key"`
	Name       string
	Department string
	Salary     float64
}

type Book struct {
	Id     uint64 `gorm:"primary_key"`
	Title  string
	Author string
	Price  float64
}

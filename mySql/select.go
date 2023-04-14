package main
//
//import (
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Person struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//type Place struct {
//	Country string `db:"country"`
//	City    string `db:"city"`
//	TelCode int    `db:"telcode"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
//	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/local")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//	Db = database
//	//defer database.Close()  // 注意这行代码要写在上面err判断的下面
//}
//
//func main() {
//	var person []Person
//	err := Db.Select(&person,"select user_id, username, sex, email from person where user_id = ? ", 2)
//	if err != nil {
//		fmt.Println("Select failed, ", err)
//		return
//	}
//
//	fmt.Println("select succ:", person)
//}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type UserInfo struct {
	ID        int          `db:"id"`
	Name      string       `db:"name"`
	Password  string       `db:"password"`
	Age       int          `db:"age"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func main() {
	db, err := sqlx.Open("sqlite3", "data/data.db")
	//con, err := sqlx.Connect("sqlite3", "data/data.db")
	if err != nil {
		log.Fatalln("database connent erro ", err)
		panic("database connent error")
	}
	defer db.Close()
	//defer con.Close()

	// sqlStmt := `drop table if exists UserInfo;
	// create table UserInfo (id integer not null primary key autoincrement, name text(64),password text(32) not null,age integer default 0, created_at date,updated_at date);
	// `

	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }
	// now := time.Now().Format("2006-01-02 15:04:05")
	// _, err = db.Exec("insert into UserInfo (name,password,age,created_at)values(?,?,?,?)", "张三", "123456", 14, now)
	// if err != nil {
	// 	fmt.Println("insert error= ", err)
	// }
	// rows, _ := db.Query("select id,name,password,age from UserInfo where id=?;", 1)
	// defer rows.Close()
	// for rows.Next() {
	// 	var user = UserInfo{}
	// 	rows.Scan(&user.ID, &user.Name, &user.Password, &user.Age)
	// 	fmt.Printf("id:%d name:%v passowrd:%v age:%v \n", user.ID, user.Name, user.Password, user.Age)
	// }

	// row := db.QueryRow("select id,name,password,age from UserInfo")
	// var user = UserInfo{}
	// row.Scan(&user.ID, &user.Name, &user.Password, &user.Age)
	// fmt.Printf("id:%d name:%v passowrd:%v age:%v \n", user.ID, user.Name, user.Password, user.Age)
	// row := db.QueryRow("select id,name,password,age from UserInfo")
	// var user = UserInfo{}
	// row.Scan(&user.ID, &user.Name, &user.Password, &user.Age)
	// fmt.Printf("id:%d name:%v passowrd:%v age:%v \n", user.ID, user.Name, user.Password, user.Age)
	// var user = UserInfo{}
	// db.Get(&user, "select id,name,password,age,created_at from UserInfo where id=?", 1)
	// fmt.Printf("id:%d name:%v passowrd:%v age:%v CreatedAt:%v \n", user.ID, user.Name, user.Password, user.Age, user.CreatedAt)
	// var users []UserInfo
	// db.Select(&users, "select id,name,password,age,created_at,updated_at from UserInfo;")
	// for _, user := range users {
	// 	fmt.Printf("id:%d name:%s passowrd:%s age:%d crated:%v updated:%v\n", user.ID, user.Name, user.Password, user.Age, user.CreatedAt.Format("2006-01-02 15:04:05"), user.UpdatedAt == sql.NullTime{})
	// }
	// result, _ := db.Exec("update UserInfo set password=?,updated_at=? where id=?", user.Password, time.Now(), 1)
	// id, _ := result.LastInsertId()
	// count, _ := result.RowsAffected()
	// fmt.Printf("result id:%d count:%d\n", id, count)

	// result, _ := db.NamedExec("update UserInfo set password=:Passowrd,updated_at=:UpdatedAt where id=:ID",
	// 	map[string]interface{}{
	// 		"Passowrd":  "走远",
	// 		"UpdatedAt": time.Now(),
	// 		"ID":        1,
	// 	})
	// id, _ := result.LastInsertId()
	// count, _ := result.RowsAffected()
	// fmt.Printf("result id:%d count:%d\n", id, count)

	// result, err := db.Exec("delete from UserInfo where id=?", 1)
	// if err != nil {
	// 	fmt.Printf("exec failed, err:%v\n", err)
	// 	return
	// }
	// count, _ := result.RowsAffected()
	// fmt.Printf("result count:%d\n", count)

	// result := db.MustExec("insert into UserInfo (id,name,password,age,created_at) values(?,?,?,?,?)", 1, "李极在", "123211", 87, time.Now())

	// id, _ := result.LastInsertId()
	// fmt.Printf("result id:%d\n", id)

	fmt.Println("hello sql")
}

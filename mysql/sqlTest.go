package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "12345678"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "test"
)

type User struct {
	Id         int           `json:"id"`
	UserName   string        `json:"username"`
	Password   string        `json:"password"`
	Status     sql.NullInt32 `json:"status"`
	CreateTime sql.NullInt64 `json:"createtime"`
}

func Sql() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	fmt.Println("connection to mysql succeed")
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数

	createTable(DB)

	//insertData(DB)

	queryOne(DB)

	queryMulti(DB)

	updateData(DB)
}

func createTable(DB *sql.DB) {
	sqlText := `create table if not exists users(
id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
username VARCHAR(64),
password VARCHAR(64),
status INT(4),
createtime INT(10)
);`

	if _, err := DB.Exec(sqlText); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table succeed")
}

func insertData(DB *sql.DB) {
	sqlText := `insert into users(username, password) values("test","123456")`
	result, err := DB.Exec(sqlText)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v\n", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return
	}
	fmt.Printf("Insert data id: %v\n", lastInsertID)

	rowsAffected, err := result.RowsAffected() //通过RowsAffected获取受影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("Affected rows:", rowsAffected)
}

func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Status, &user.CreateTime); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("Single row data: %v\n", *user)
}

func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from users where id=?", 1)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Status, &user.CreateTime) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Println("scan succeed:", *user)
	}
}

func updateData(DB *sql.DB) {
	result, err := DB.Exec("UPDATE users set password=? where id=?", "111111", 2)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	fmt.Println("update data succeed:", result)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Println("Affected rows:", rowsAffected)
}

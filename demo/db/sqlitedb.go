// package main

// import (
// 	"bytes"
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/mattn/go-sqlite3"
// 	"time"
// )

// func main2() {
// 	//可以安裝SQLite管理器(http://sqliteadmin.orbmu2k.de/)
// 	db, e := sql.Open("sqlite3", "./sqlite3Test.db")
// 	defer db.Close()
// 	throwIf(e)

// 	//insertUserInfo(db)
// 	statement, e := db.Prepare("Insert userinfo Set username=?, departname=?, created=?")
// 	throwIf(e)
// 	result, e := statement.Exec("Jagre Zhang", "技术研发department", time.Now().Format("2006-01-02"))
// 	throwIf(e)
// 	id, e := result.LastInsertId()
// 	throwIf(e)
// 	fmt.Printf("UserInfo.uid: %d\n", id)

// 	result, e = statement.Exec("Susan Liang", "毕加索教育班", time.Now().Format("2006-01-02"))
// 	throwIf(e)
// 	id, e = result.LastInsertId()
// 	throwIf(e)
// 	fmt.Printf("UserInfo.uid: %d\n", id)

// 	//update
// 	statement, e = db.Prepare("Update userinfo Set departname=? Where uid=?")
// 	throwIf(e)
// 	result, e = statement.Exec("毕加索教育班.初级", id)
// 	throwIf(e)
// 	affect, e := result.RowsAffected()
// 	throwIf(e)
// 	fmt.Printf("Update affected rows: %d\n", affect)

// 	//query
// 	rows, e := db.Query("Select * From userinfo")
// 	throwIf(e)

// 	var buffer bytes.Buffer
// 	var arr [4]string
// 	for rows.Next() {
// 		buffer.WriteString("{")
// 		rows.Scan(&arr[0], &arr[1], &arr[2], &arr[3])
// 		cols, e := rows.Columns()
// 		throwIf(e)
// 		for i, col := range cols {
// 			buffer.WriteString("\"")
// 			buffer.WriteString(col)
// 			buffer.WriteString("\"")
// 			buffer.WriteString(": ")
// 			buffer.WriteString("\"")
// 			buffer.WriteString(arr[i])
// 			buffer.WriteString("\", ")
// 		}
// 		buffer.WriteString("}")
// 	}
// 	fmt.Print("[" + buffer.String() + "]")

// 	//Delete
// 	// statement, e = db.Prepare("Delete From userinfo Where uid=?")
// 	// throwIf(e)
// 	// result, e = statement.Exec(id)
// }

// //redeclare
// // func throwIf(e error) {
// // 	if e != nil {
// // 		panic(e)
// // 	}

// // }

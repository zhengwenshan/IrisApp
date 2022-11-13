package controller

import (
	"database/sql"
	"fmt"
	"reflect"

	//mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Age  int32  `db:"age" json:"age"`
	Name string `db:"name" json:"name"`
}

func DbInvoke() {
	dsn := "root:root@tcp(localhost:3306)/test?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	db, err := ConnectMySQL(dsn)
	if err == nil {
		rows, err := db.Query("select * from person")
		if err == nil {
			if !rows.Next() {
				return
			}

			//types, err := rows.ColumnTypes()
			//if err == nil {
			//	fmt.Print(types)
			//}

			//columns, err := rows.Columns()
			//if err==nil {
			//	fmt.Print(columns)
			//
			//}


			p := new(Person)
			of := reflect.ValueOf(p)

			err := rows.Scan(of.FieldByName("Age"), of.FieldByName("Name"))
			if err == nil {
				fmt.Print(p)
			} else {
				fmt.Println(err)
			}
		}
	}
}

func ConnectMySQL(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

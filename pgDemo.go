package main

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"fmt"
)
var db *sql.DB

func sqlOpen()  {
	var err error
	db, err = sql.Open("postgres","port=5432 user=postgres password=root dbname=userdemo sslmode=disable")
	checkErr(err)
}
func sqlClose()  {
	db.Close()
}
func sqlInsert()  {
	//insert
	stmt, err := db.Prepare("INSERT INTO public.user(id,username) VALUES($1,$2)")
	fmt.Println(stmt)
	res,err := stmt.Exec("6","yang6")
	fmt.Println(res)
	affect , err := res.RowsAffected()
	checkErr(err)
	fmt.Println("insert row affected: ",affect)
}
func sqlQuery()  {
	rows, err := db.Query("SELECT * FROM public.user")
	checkErr(err)
	for rows.Next() {
		var id 			string
		var username 	string
		err = rows.Scan(&id,&username)
		checkErr(err)
		fmt.Println("id = ",id, "\nname = ",username)
	}
}
func sqlUpdate()  {
	stmt, err := db.Prepare("UPDATE public.user SET username=$1 WHERE id=$2")
	checkErr(err)
	res,err := stmt.Exec("yang666","6")
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("update rows affected: ",affect)
}
func sqlDelete()  {
	stmt,err := db.Prepare("DELETE FROM public.user WHERE id=$1")
	res,err := stmt.Exec("6")
	checkErr(err)
	affect,err := res.RowsAffected()
	checkErr(err)
	fmt.Println("delete rows affected: ",affect)
}

func main()  {
	sqlOpen()
	sqlInsert()
	sqlQuery()
	sqlUpdate()
	sqlDelete()
	sqlClose()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}
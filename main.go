package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "postgres"
)

type ListData struct {
	list_id int
	title   string
	is_comp bool
	date    string
}

func main() {
	GetConvItem()
	Show(4)
	Insert("new", true)
	Update("old", false, 3)
	Delete(5)
}

func GetConvItem() {
	rows, err := Connected().Query("select list_id,title,is_comp,date from list ")
	CheckError(err)
	// ConvItem := []ListData{}
	defer rows.Close()
	for rows.Next() {
		ConvItem := ListData{}
		err = rows.Scan(&ConvItem.list_id, &ConvItem.title, &ConvItem.is_comp, &ConvItem.date)
		CheckError(err)
		fmt.Println(ConvItem)
	}

}

func Insert(title string, is_comp bool) {
	db := Connected()
	now := time.Now()
	secs := now.Unix()
	insForm, err := db.Prepare("INSERT INTO list(title,is_comp,date) VALUES($1,$2,$3)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(title, is_comp, secs)

	defer db.Close()
	fmt.Println("success fully!!")
}

func Update(title string, is_comp bool, list_id int) {
	db := Connected()
	now := time.Now()
	secs := now.Unix()
	insForm, err := db.Prepare("UPDATE list SET title=$1, is_comp=$2 ,date=$3 WHERE list_id=$4")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(title, is_comp, secs, list_id)
	defer db.Close()
	fmt.Println("success fully!!")
}

func Delete(list_id int) {
	db := Connected()
	delForm, err := db.Prepare("DELETE FROM list WHERE list_id=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(list_id)
	log.Println("DELETE")
	defer db.Close()
}

func Show(list_id int) {
	rows, err := Connected().Query("select list_id,title,is_comp,date from list where list_id = $1", list_id)
	CheckError(err)
	// ConvItem := []ListData{}
	defer rows.Close()
	for rows.Next() {
		ConvItem := ListData{}
		err = rows.Scan(&ConvItem.list_id, &ConvItem.title, &ConvItem.is_comp, &ConvItem.date)
		CheckError(err)
		fmt.Println(ConvItem)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Connected() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

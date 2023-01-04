package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type Users struct {
	Id      int
	Login   string
	Surname string
	Name    string
	Role    int
}

var database *sql.DB

func GetUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("select * from toy_shop.users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	users := []Users{}

	for rows.Next() {
		p := Users{}
		err := rows.Scan(&p.Id, &p.Login, &p.Surname, &p.Name, &p.Role)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, users)
}

func main() {

	db, err := sql.Open("mysql", "root:root@/toy_shop")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", GetUsers)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

	//var inputNumber int
	//
	//for {
	//	fmt.Println("1. Add cart\n2. Add cart item to cart\n3. Remove item from cart\n4. View cart")
	//	fmt.Scanf("%d\n", &inputNumber)
	//
	//	switch inputNumber {
	//	case 1:
	//		service.AddCart()
	//	case 2:
	//		service.AddCartItem()
	//	case 3:
	//		service.DeleteCartItem()
	//	case 4:
	//		service.ShowCartItems()
	//	case 5:
	//		service.ShowCartItems()
	//	default:
	//		fmt.Println("This case is not exist")
	//	}
	//}
}

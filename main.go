package main

import (
	"database/sql"
	"example.com/m/v2/handlers"
	"example.com/m/v2/repository"
	"example.com/m/v2/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var database *sql.DB

func main() {

	db, err := sql.Open("mysql", "root:root@/toy_shop")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	usersRepository := repository.NewUserRepository(db)
	rolesRepository := repository.NewRoleRepository(db)

	usersService := service.NewUserService(usersRepository)
	rolesService := service.NewRoleService(rolesRepository)

	usersHandler := handlers.NewUserHandler(usersService)
	rolesHandler := handlers.NewRoleHandler(rolesService)

	//methods for baskets
	//methods for products
	//methods for reviews
	//methods for orders
	//methods for products_baskets
	//methods for products_orders
	//methods for statuses

	http.HandleFunc("/", usersHandler.GetUsers)
	http.HandleFunc("/roles", rolesHandler.GetRoles)

	fmt.Println("Server is listening...")
	fmt.Println("localhost:8181")
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

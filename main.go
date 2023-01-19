package main

import (
	"database/sql"
	"example.com/m/v2/handlers"
	"example.com/m/v2/repository"
	"example.com/m/v2/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	address string
	router  *mux.Router
	db      *sql.DB
}

func NewApp() App {
	address := fmt.Sprintf("%s:%d", "localhost", 8181)
	db, err := sql.Open("mysql", "root:root@/toy_shop")
	if err != nil {
		log.Println(err)
	}

	router := mux.NewRouter()
	app := App{address, router, db}
	app.setRouters()

	return app
}

func (app *App) setRouters() {
	usersRepository := repository.NewUserRepository(app.db)
	rolesRepository := repository.NewRoleRepository(app.db)
	basketRepository := repository.NewBasketRepository(app.db)
	productsRepository := repository.NewProductRepository(app.db)
	reviewsRepository := repository.NewReviewRepository(app.db)

	usersService := service.NewUserService(usersRepository)
	rolesService := service.NewRoleService(rolesRepository)
	basketService := service.NewBasketService(basketRepository, productsRepository)
	productsService := service.NewProductService(productsRepository)
	reviewsService := service.NewReviewService(reviewsRepository)

	usersHandler := handlers.NewUserHandler(usersService)
	rolesHandler := handlers.NewRoleHandler(rolesService)
	basketHandler := handlers.NewBasketHandler(basketService)
	productsHandler := handlers.NewProductHandler(productsService)
	reviewsHandler := handlers.NewReviewHandler(reviewsService)

	//methods for reviews
	//methods for orders
	//methods for statuses

	//Users
	app.router.HandleFunc("/api/user", usersHandler.Add).Methods("POST")
	app.router.HandleFunc("/api/user/{userId}", usersHandler.GetUser).Methods("GET")
	app.router.HandleFunc("/api/users", usersHandler.GetUsers).Methods("GET")
	app.router.HandleFunc("/api/user/{userId}", usersHandler.Edit).Methods("PUT")
	app.router.HandleFunc("/api/user/{userId}", usersHandler.Delete).Methods("DELETE")
	//Roles
	app.router.HandleFunc("/api/role/{roleId}", rolesHandler.GetRole).Methods("GET")
	app.router.HandleFunc("/api/roles", rolesHandler.GetRoles).Methods("GET")
	//Baskets
	app.router.HandleFunc("/api/basket/product", basketHandler.AddProductToBasket).Methods("POST")
	app.router.HandleFunc("/api/basket/product/{basketId}/{productId}", basketHandler.DecreaseQuantityProductToBasket).Methods("PUT")
	app.router.HandleFunc("/api/basket/product/{productId}", basketHandler.DeleteProductToBasket).Methods("DELETE")
	app.router.HandleFunc("/api/basket", basketHandler.GetBasket).Methods("GET")
	//Products
	app.router.HandleFunc("/api/product", productsHandler.GetProduct).Methods("GET")
	app.router.HandleFunc("/api/products", productsHandler.GetProducts).Methods("GET")
	//Reviews
	app.router.HandleFunc("/api/review", reviewsHandler.GetReview).Methods("GET")
	app.router.HandleFunc("/api/reviews-product", reviewsHandler.GetReviewsProduct).Methods("GET")
}

func (app *App) Run() {
	log.Println(http.ListenAndServe(app.address, app.router))
}

func (app *App) Stop() {
	err := app.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	api := NewApp()
	api.Run()
	defer api.Stop()
}

package main

import (
	"database/sql"
	handlers2 "example.com/m/v2/internal/handlers"
	repository2 "example.com/m/v2/internal/repository"
	service2 "example.com/m/v2/internal/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type App struct {
	address string
	router  *mux.Router
	db      *sql.DB
}

func NewApp() App {
	db, err := sql.Open("mysql", "root:root@/toy_shop")
	if err != nil {
		log.Println(err)
	}

	router := mux.NewRouter()
	app := App{":8181", router, db}
	app.setRouters()

	return app
}

func (app *App) setRouters() {
	usersRepository := repository2.NewUserRepository(app.db)
	rolesRepository := repository2.NewRoleRepository(app.db)
	basketRepository := repository2.NewBasketRepository(app.db)
	productsRepository := repository2.NewProductRepository(app.db)
	reviewsRepository := repository2.NewReviewRepository(app.db)
	ordersRepository := repository2.NewOrderRepository(app.db)
	statusesRepository := repository2.NewStatusRepository(app.db)

	usersService := service2.NewUserService(usersRepository)
	rolesService := service2.NewRoleService(rolesRepository)
	basketService := service2.NewBasketService(basketRepository, productsRepository)
	productsService := service2.NewProductService(productsRepository)
	reviewsService := service2.NewReviewService(reviewsRepository)
	ordersService := service2.NewOrderService(ordersRepository, basketRepository, productsRepository)
	satusesService := service2.NewStatusService(statusesRepository)

	usersHandler := handlers2.NewUserHandler(usersService)
	rolesHandler := handlers2.NewRoleHandler(rolesService)
	basketHandler := handlers2.NewBasketHandler(basketService)
	productsHandler := handlers2.NewProductHandler(productsService)
	reviewsHandler := handlers2.NewReviewHandler(reviewsService)
	ordersHandler := handlers2.NewOrderHandler(ordersService)
	statusesHandler := handlers2.NewStatusHandler(satusesService)

	// Users
	app.router.HandleFunc("/api/users", usersHandler.Add).Methods("POST")
	app.router.HandleFunc("/api/users/{userId}", usersHandler.GetUser).Methods("GET")
	app.router.HandleFunc("/api/users", usersHandler.GetUsers).Methods("GET")
	app.router.HandleFunc("/api/users/{userId}", usersHandler.Edit).Methods("PUT")
	app.router.HandleFunc("/api/users/{userId}", usersHandler.Delete).Methods("DELETE")
	// Roles
	app.router.HandleFunc("/api/roles/{roleId}", rolesHandler.GetRole).Methods("GET")
	app.router.HandleFunc("/api/roles", rolesHandler.GetRoles).Methods("GET")
	// Baskets
	app.router.HandleFunc("/api/baskets/product", basketHandler.AddProductToBasket).Methods("POST")
	app.router.HandleFunc("/api/baskets/product/{basketId}/{productId}", basketHandler.DecreaseQuantityProductToBasket).Methods("PUT")
	app.router.HandleFunc("/api/baskets/product/{productId}", basketHandler.DeleteProductToBasket).Methods("DELETE")
	app.router.HandleFunc("/api/baskets", basketHandler.GetBasket).Methods("GET")
	// Products
	app.router.HandleFunc("/api/products", productsHandler.AddProduct).Methods("POST")
	app.router.HandleFunc("/api/products/{productId}", productsHandler.GetProduct).Methods("GET")
	app.router.HandleFunc("/api/products/{productId}", productsHandler.EditProduct).Methods("PUT")
	app.router.HandleFunc("/api/products/{productId}", productsHandler.DeleteProduct).Methods("DELETE")
	app.router.HandleFunc("/api/products", productsHandler.GetProducts).Methods("GET")
	// Reviews
	app.router.HandleFunc("/api/reviews", reviewsHandler.AddReview).Methods("POST")
	app.router.HandleFunc("/api/reviews/{reviewId}", reviewsHandler.EditReview).Methods("PUT")
	app.router.HandleFunc("/api/reviews/{reviewId}", reviewsHandler.GetReview).Methods("GET")
	app.router.HandleFunc("/api/reviews/{reviewId}", reviewsHandler.DeleteReview).Methods("DELETE")
	app.router.HandleFunc("/api/reviews/{productId}", reviewsHandler.GetReviewsProduct).Methods("GET")
	app.router.HandleFunc("/api/reviews", reviewsHandler.GetReviews).Methods("GET")
	// Orders
	app.router.HandleFunc("/api/orders/{userId}", ordersHandler.AddOrder).Methods("POST")
	app.router.HandleFunc("/api/orders/{orderId}", ordersHandler.GetOrder).Methods("GET")
	app.router.HandleFunc("/api/orders/{orderId}", ordersHandler.DeleteOrder).Methods("DELETE")
	app.router.HandleFunc("/api/orders", ordersHandler.GetOrders).Methods("GET")
	// Statuses
	app.router.HandleFunc("/api/statuses/{statusId}", statusesHandler.GetStatus).Methods("GET")
	app.router.HandleFunc("/api/statuses", statusesHandler.GetStatuses).Methods("GET")

	app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Api.")
	})

	fmt.Println("Server listening!")
}

func (app *App) Run() {
	server := &http.Server{
		Addr:         app.address,
		Handler:      app.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	log.Println("Start server")
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

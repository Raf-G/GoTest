package main

import (
	"database/sql"
	_ "example.com/m/v2/docs"
	"example.com/m/v2/internal/handlers"
	"example.com/m/v2/internal/repository"
	"example.com/m/v2/internal/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"time"
)

//	@title			GoTest Swagger API
//	@version		1.0
//	@description	Swagger API for Golang GoTest.

//	@contact.name	Rovshan Gasanov
//	@contact.email	rovshan27121@gmail.com

//	@host		localhost:8181
//	@BasePath	/api/

type App struct {
	address     string
	router      *mux.Router
	db          *sql.DB
	redisClient *redis.Client
}

func NewApp() App {
	db, err := sql.Open("mysql", "root:root@/toy_shop")
	if err != nil {
		log.Println(err)
	}

	router := mux.NewRouter()

	//ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//err = rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}

	app := App{":8181", router, db, rdb}
	app.setRouters()

	return app
}

func (app *App) setRouters() {
	usersRepository := repository.NewUserRepository(app.db)
	rolesRepository := repository.NewRoleRepository(app.db, app.redisClient)
	basketRepository := repository.NewBasketRepository(app.db)
	productsRepository := repository.NewProductRepository(app.db)
	reviewsRepository := repository.NewReviewRepository(app.db)
	ordersRepository := repository.NewOrderRepository(app.db)
	statusesRepository := repository.NewStatusRepository(app.db)

	usersService := service.NewUserService(usersRepository)
	rolesService := service.NewRoleService(rolesRepository)
	basketService := service.NewBasketService(basketRepository, productsRepository)
	productsService := service.NewProductService(productsRepository)
	reviewsService := service.NewReviewService(reviewsRepository)
	ordersService := service.NewOrderService(ordersRepository, basketRepository, productsRepository)
	satusesService := service.NewStatusService(statusesRepository)

	usersHandler := handlers.NewUserHandler(usersService)
	rolesHandler := handlers.NewRoleHandler(rolesService)
	basketHandler := handlers.NewBasketHandler(basketService)
	productsHandler := handlers.NewProductHandler(productsService)
	reviewsHandler := handlers.NewReviewHandler(reviewsService)
	ordersHandler := handlers.NewOrderHandler(ordersService)
	statusesHandler := handlers.NewStatusHandler(satusesService)

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
	app.router.HandleFunc("/api/baskets/{basketId}", basketHandler.GetBasket).Methods("GET")
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
	app.router.HandleFunc("/api/reviews/product/{productId}", reviewsHandler.GetReviewsProduct).Methods("GET")
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

	app.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

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

package main

import (
	"database/sql"
	api_handlers "example.com/m/v2/handlers"
	"example.com/m/v2/routes"
	"example.com/m/v2/utils"
	_ "github.com/go-sql-driver/mysql"
	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Db *sql.DB

func main() {

	var err error
	dsn := "testuser:testpassword@tcp(localhost:3306)/testdb"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	//Check(Session)
	defer Db.Close()

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	myHandler := api_handlers.Handler{Db: Db}

	r := routes.ApiRouter{
		Router:  mux.NewRouter().StrictSlash(true),
		Handler: &myHandler,
	}
	//routes
	r.Home("/")
	r.Create("/create")
	r.GetOne("/get/{id}")
	r.GetAll("/getAll")
	r.Count("/count")
	r.Delete("/delete/{id}")
	r.DeleteAll("/deleteAll")
	r.Update("/update/{id}")

	headers := gorilla_handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := gorilla_handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := gorilla_handlers.AllowedOrigins([]string{"*"})

	//r.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	if serveErr := http.ListenAndServe(":4040", gorilla_handlers.CORS(headers, methods, origins)(r.Router)); serveErr != nil {
		utils.Panic(serveErr, "Failed to run a backend")
	}

}

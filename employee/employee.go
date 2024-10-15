package employee
import (
	"github.com/labstack/echo/v4"
	"os"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
     "database/sql"
	 _ "github.com/lib/pq" // Import the PostgreSQL driver
    "log"


)
var db *sql.DB
var e=echo.New()
var v=validator.New()

func Start(){

	connStr := "host=localhost dbname=postgres port=5432 user=postgres password=veameec121 sslmode=disable"
	var err error
	db,err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:",err)
	}
   
    port:=os.Getenv("MY_APP_PORT")
	if port==""{
		port="8080"
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.POST("/employee",createEmployees)
	e.GET("/employee",getEmployees)
	e.GET("/employee/:id",getEmployee)
	e.PUT("/employee/:id",updateEmployees)
	e.PATCH("/employee/:id",updateEmployee)
	e.DELETE("/employee/:id",deleteEmployee)
	e.Logger.Print(fmt.Sprintf("Listening to the port %v",port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v",port)))

}

package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ChanchalS7/golang_postgres/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/tools/go/analysis/passes/defers"
)

//response format
type response struct {
	ID		int64	`json:"id,omitempty"`
	Message	string	`json:"message,omitempty"`
}

//Create connection with postgres db
func createConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Open the connection

	db, err := sql.Open("postgres",os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//return the connection 
	return db 

}

//CreateStock create a stock in the postgres db
func CreateStock(){}
//GetStock will return a single stock by its id
func GetStock(){}
//GetAllStock will return all the stock
func GetAllStock(){}
//UpdateStock update stock's details in the postgres db
func UpdateStock(){}
//DeleteStock delete stock's detail in the postgres db
func DeleteStock(){}

// ------------------handler functions -------------------------
//func insertStock in the DB
func insertStock(stock models.Stock) int64 {
db := createConnection()
defer db.Close()
sqlStatement:=`INSERT INTO stocks (name,price,company) VALUES ($1, $2, $3) RETURNING stockid`

var id int64 

err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

if err != nil {
	log.Fatalf("Unable to execute the query. %v", err)
}
fmt.Printf("Inserted a single record. %v",id)
return id
}
// get one stock from the DB by its stockid
func getStock(id int64) (models.Stock,error){
	//Create the postgres db connection
	db := createConnection()

	//close the db connection
	defer db.Close()

	//create a stock of model.Stock type
	var stock models.Stock

	//create the select sql query
	sqlStatement := `SELECT * FROM stock WHERE stockid = $1`

	//Execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	//unmarshal the row object to stock
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
			fmt.Println("No rows were returned")
			return stock,nil
	case nil :
		return stock, nil	
	default:
		log.Fatalf("Unable to scan the row. %v",err)		
	}
	//return empty stock on error
	return stock, err
}
//get all stock from the db by its stockid 
func getAllStocks() ([]models.Stock, error){
	//Create the postgres db connection
	db := createConnection()

	//close the db connection
	defer db.Close()

	var stocks []models.Stock

	//Create the select sql query
	sqlStatement := `SELECT * FROM stocks`

	//Execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v",err)
	}
	//close the statement
	defer rows.Close()

	//iterate over the rows
	for rows.Next() {
		var stock models.Stock

		//unmarshall the row object to stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err !=nil {
			log.Fatal("Unable to scan the row. %v",err)
		}
		//append the stock in stocks slice
		stocks = append(stocks,stock)
	}
	//return empty stock on error
	return stocks, err 
}
//update stock in the db
func updateStock(){}
//delete stock in the db
func deletestock(){}

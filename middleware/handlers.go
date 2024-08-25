package middleware
import(
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ChanchalS7/golang_postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//response format
type response struct {
	ID		int64	`json:"id,omitempty"`
	Message	string	`json:"message,omitempty"`
}

//Create connection with postgres db
func CreateConnection() *sql.DB {
	
}
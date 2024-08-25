package main
import (
	"fmt"
	"github.com/ChanchalS7/golang_postgres/router"
	"log"
	"net/http"
)
func main(){
	r:= router.Router()
	fmt.Println("Start server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080",r))
}
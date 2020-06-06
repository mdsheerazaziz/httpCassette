package main
/*
	Incoming Request http://localhost:3000/getLatestUpdates
	Forward Request http://localhost:4000/api/v1/getLatestUpdates

*/
import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)
type  Request struct {
	incomingURL string
}
func Parse(w http.ResponseWriter, req *http.Request){
	fmt.Println("Incoming Request")
}
func main(){
	http.HandleFunc("/process", Parse)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
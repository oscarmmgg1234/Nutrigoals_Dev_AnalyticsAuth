package main

import (
	
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

var port string = ":5005"

var emails = [2]string{"oscarmmgg1234@gmail.com","edwardcordero@gmail.com"}
var passwords = [2]string{"oscarTester23","edwardTester25"}

var isValid = `{"authorized" : "true"}`
var notValid = `{"authorized" : "false"}`

func credentialValidation(username string, password string) bool { 
var result bool = false
if username != "" {
	for i := 0; i < len(emails); i++ {
		if username == emails[i] {
			for i := 0; i < len(passwords); i++ {
				if password == passwords[i]{
					result = true
				}
			}
		}
	}
	return result
} else{
	return result
}

}

func userValidation(w http.ResponseWriter, r* http.Request){
rUsername := r.Header.Get("username")
rPassword := r.Header.Get("password")
result := credentialValidation(rUsername, rPassword)
w.Header().Set("Content-Type", "application/json")

if result == true {
	w.Write([]byte(isValid))
} else {
	w.Write([]byte(notValid))
}


}


func main(){
	router := mux.NewRouter()
	
	router.HandleFunc("/auth", userValidation).Methods("GET")

		log.Fatal(http.ListenAndServe(port, router))

}
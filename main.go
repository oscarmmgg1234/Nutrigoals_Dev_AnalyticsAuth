package main

import (
	
	"net/http"
	"fmt"
	"log"
)

var port string = ":5005"

var emails = [2]string{"oscarmmgg1234@gmail.com","edwardcordero@gmail.com"}
var passwords = [2]string{"oscarTester23","edwardTester25"}
var jsonStr string = `{"title":"Buy cheese and bread for breakfast."}`

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
if result == true {
	fmt.Fprintf(w, jsonStr)
} else {
	fmt.Fprintf(w, jsonStr)
}


}

func handleRequest(){
	http.HandleFunc("/auth", userValidation)
}




func main(){
	
	handleRequest()
	log.Fatal(http.ListenAndServe(port, nil))

}
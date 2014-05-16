package main

import(
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
)

func PeopleHandler(w http.ResponseWriter, req *http.Request){
        if(req.Method == "GET") {
                fmt.Println("Hello")
        
        }

}

func main(){
  r := mux.NewRouter()
  r.HandleFunc("/people", PeopleHandler)
  http.Handle("/", r)
  http.ListenAndServe(":8003", nil) 
}

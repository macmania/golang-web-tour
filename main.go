/**
* Playing around with using Go as a server back-end service for API services
*/

package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "fmt"
)

//Data model
type Person struct {
  Name string
  Age int
}

var ryanne Person

func ryanneHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.Path)
  switch req.Method {
    case "GET":
      buf, _ := json.Marshal(&ryanne)
      w.Write(buf)

    //curl -v -H "Content-Type: application/json" -X PUT --data "stuff.json" http://localhost:8003/ryanne
    //this uploads
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      err := json.Unmarshal(buf, &ryanne)
      if err != nil {
        fmt.Println("error: ", err)
      }
      fmt.Printf("%s %d", ryanne.Name, ryanne.Age)

    default:
      w.WriteHeader(400)
  }
}


func handler(w http.ResponseWriter, req *http.Request){
  switch req.URL {
    case "/register":
      //go to this function handler
      //and handle the registration process
    case "/login":


  }
}


func main() {
  ryanne.Name = "Ryanne"
  ryanne.Age = 25

  http.HandleFunc("/ryanne", ryanneHandler)

  http.ListenAndServe(":8003", nil)
}

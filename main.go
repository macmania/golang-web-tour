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

type Person struct {
  UserName string
  Name string
  Age int
  Password string
  EmailAddress string
}
var personTemp Person
var manager PeopleManager

//a collection array that stores the list of people
type PeopleManager struct {
  peopleManager map[string]*Person
}


func peopleHandler(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.URL.Path)

  switch req.Method {
    case "GET":

      buf, _ := json.Marshal(&manager)
      w.Write(buf)

    //curl -v -H "Content-Type: application/json" -X PUT --data "stuff.json" http://localhost:8003/ryanne
    //this uploads
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      err := json.Unmarshal(buf, &personTemp)
      if printErr (&err) == 1{
        fmt.Printf("%s %d", personTemp.Name, personTemp.Age)

    }

    case "DELETE":
      buf, _ := ioutil.ReadAll(req.Body)
      err := json.Unmarshal(buf, &personTemp)
      printErr (&err)


    case "POST":
      //to-do stub

    default:
      w.WriteHeader(400)
  }
}

//Handles each specific person that is in the peopleHandler
func personHandler(w http.ResponseWriter, req *http.Request){

    exists := isPersonExists(req.URL.Path[1:])
    person := req.URL.Path[1:]
    switch req.Method{
      case "GET":

        if exists {
          buf, _ := json.Marshal(&personTemp)
          w.Write(buf)
        }
        else {
          w.Write("couldn't find ", person)
        }

      case "PUT":
        if exists {
          buf, err := json.Unmarshal(buf, manager[&person])
        }
        printErr (&err)
      case "DELETE":
        if exists {
          //delete the resource off the server
        }

      case "POST":
        if !exists {
          //
        }

      default:
        w.WriteHeader(400) //not found code

}

func printErr (err *string) error{
  if &err != nil{
    fmt.Println("error: ", &err)
    return nil
  }
  else return 1
}

func isPersonExists(person *string) error boolean {
  if val, ok := manager[&person]; ok {
    return true
  }
  else {
    return false
  }
}

func main() {

  //starting variables
  manager.peopleManager = make(map[string]*Person)
  //personTemp.UserName = ""

  http.HandleFunc("/people", peopleHandler)
  http.HandleFunc("/{people}", personHandler)

  http.ListenAndServe(":8003", nil)
}

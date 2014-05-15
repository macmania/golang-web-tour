/**
* Playing around with using Go as a server back-end service for API services
* Starter kit on using Go to implement routing services
* to-do, need to test the concurrency part for this project
*/

package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "fmt"
  "html/template"
  "log"
)

type Context struct {
    Title  string
    Static string
}

const STATIC_URL = "/"

//data model for each particular person to be saved in the map
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
  fmt.Println(req.URL.Path[1:])

  switch req.Method {
    case "GET":

      buf, _ := json.Marshal(&manager.peopleManager)
      w.Write(buf)

    //curl -v -H "Content-Type: application/json" -X PUT --data "stuff.json" http://localhost:8003/ryanne
    //this uploads
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      err := json.Unmarshal(buf, &manager.peopleManager)
      if err != nil{
        fmt.Printf("%s %d", personTemp.Name, personTemp.Age)

    }

    case "DELETE":
      buf, _ := ioutil.ReadAll(req.Body)
      err := json.Unmarshal(buf, &manager.peopleManager)
      printErr (err)

    default:
      w.WriteHeader(400)
  }
}

//Handles each specific person that is in the peopleHandler using their username
func personHandler(w http.ResponseWriter, req *http.Request){
    fmt.Print("HERE")
    exists := isPersonExists(req.URL.Path[1:])
    var person string
    person = req.URL.Path[1:]
    switch req.Method{
      case "GET":
        if exists {
          buf, _ := json.Marshal(manager.peopleManager[person])
          w.Write(buf)
        }else {
          fmt.Print("Cannot find the person")
        }

      case "PUT":
          /*Need to test if this is ok*/
        //if exists {
          buf, _ := ioutil.ReadAll(req.Body)
          err := json.Unmarshal(buf, manager.peopleManager[person])
          printErr(err)
        //}
/*        else { //creates a new person and add to the manager map
          manager[&person]
        }*/

      case "DELETE":
        if exists {
          //delete the resource off the server
        }

      //Post appends and updates a resource
      case "POST":
        buf, _ := ioutil.ReadAll(req.Body)
        if exists {
          err := json.Unmarshal(buf, manager.peopleManager[person])
          printErr(err)
        } else {
          fmt.Print("Cannot find the person")
        }

      default:
        w.WriteHeader(400) //not found code

    }
}

func printErr (err error){
  if err != nil {
    fmt.Println("error: ", err)
  }
}

func isPersonExists(person string) (isExists bool) {
  if _, ok := manager.peopleManager[person]; ok {
    return true
  } else {
    return false
  }
}
/*******HTML templating*******/
func Index(w http.ResponseWriter, req *http.Request){
  context := Context{Title: "Index"}
  render(w, "index.html", context)
}


//all credits go to http://www.reinbach.com/golang-webapps-1.html
func render(w http.ResponseWriter, tmpl string, context Context) {
    context.Static = STATIC_URL
    //not working right now, still need to fix the memory location for this portion
    //here
    tmpl_list := []string{"index.html",
        fmt.Sprintf("%s.html", tmpl)}
    t, err := template.ParseFiles(tmpl_list...)
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, context)
    if err != nil {
        log.Print("template executing error: ", err)
    }
}



/*******HTML templating*******/
func main() {

  //starting variables
  manager.peopleManager = make(map[string]*Person)
  //personTemp.UserName = ""
  /**
  UserName string
  Name string
  Age int
  Password string
  EmailAddress string
  **/
  //still need to encrypt this password to something else
  //might want to do some validation on the server side for the email address
  //need to use sendgrid to send the email address
  person1 := &Person{
                UserName: "jojofabe123",
                Name: "Jouella",
                Age: 15,
                Password: "somethingxyz",
                EmailAddress: "jojofabe@gmail.com",
            }

  manager.peopleManager["jojofabe123"] = person1
  http.HandleFunc("/people", peopleHandler)
  http.HandleFunc("/", Index)

  //need to know how to handle this case here
  http.HandleFunc("/person/", personHandler)

  http.ListenAndServe(":8003", nil)
}

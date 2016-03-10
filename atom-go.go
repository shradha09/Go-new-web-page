package main

import (
 "fmt"
 "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
 "bufio"
 "log"
 "html/template"
 "net/http"
 "os"
 "strings"

)
type Name struct {
 Id bson.ObjectId `bson:"_id"`
 MyName QueryName

}

type QueryName struct {
 FirstName string `bson:"FirstName"`
 LastName string `bson:"LastName"`
}

func (q *QueryName) ConvertToInterface() (interface{}) {
 return  map[string]interface{}{
   "FirstName": q.FirstName,
                        "LastName": q.LastName,
                }
}

func myfunction()






func login(w http.ResponseWriter, r *http.Request) {
fmt.Println("method:", r.Method) //get request method
if r.Method == "GET" {
    t, _ := template.ParseFiles("htmlcode.gtpl")
    t.Execute(w, nil)
} else {
    r.ParseForm()
    // logic part of log in
 fmt.Println("FirstNamename:", r.FormValue("FirstNamename"))
    fmt.Println("LastName:", r.FormValue("LastName"))
}

}


func main() {
  http.HandleFunc("/", login) // setting router rule

     http.HandleFunc("/login", login)
     err := http.ListenAndServe(":9090", nil) // setting listening port
     if err != nil {
         log.Fatal("ListenAndServe: ", err)
     }
 session, sessionErr := mgo.Dial("localhost")
 if sessionErr != nil {
  fmt.Println(sessionErr)
 } else {
  fmt.Println("Session created")

  r :=bufio.NewReader(os.Stdin)
   fmt.Print("First name: ")
   first, _ := r.ReadString('\n')
   fmt.Print("Last name: ")
   last, _ := r.ReadString('\n')
   first = strings.TrimSpace(first)
   last = strings.TrimSpace(last)

   database := session.DB("testdb")
   collection := database.C("names")
  /* collection.DropCollection()
   collection = database.C("names") */
  // nameForQuery := QueryName{FirstName: first, LastName: last}
  // query1 := collection.Find(bson.M{"myname": bson.M{"FirstName": first, "LastName": last}})
//   param := nameForQuery.ConvertToInterface()
  // query1 := collection.Find(bson.M{"myname": param})
  //count, _ := query1.Count()

   name := Name{Id: bson.NewObjectId(), MyName: QueryName{FirstName: first, LastName: last}}
    add_err := collection.Insert(name)
    if add_err != nil {
     fmt.Println("Error on add:", add_err)
    } else {
     fmt.Println("Name was successfully added")
    }
    var results []Name
collection.Find(nil).All(&results)
for _, name := range results {
 fmt.Println(name.Id, ":", name.MyName.FirstName, name.MyName.LastName)
}
var results2 []bson.M
              collection.Find(nil).All(&results2)
for _, obj := range results2 {
 fmt.Println(obj)
}
session.Close()
}
}

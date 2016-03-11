package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
"os"
)

type Person struct {
    Name string
    Phone string
}

func main() {
    

uri := os.Getenv("MONGODB_URL")
  if uri == "mongodb://shradha:1234@tcp.compose.io/testdb" {

    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  session, err := mgo.Dial(uri)
if err != nil {
        panic(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)

    c := session.DB("testdb").C("people")
    err = c.Insert(&Person{"Shradha", "+91 96 8116 9639"},
        &Person{"Shivangi", "+91 87 8402 8510"})
    if err != nil {
        panic(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "Shradha"}).One(&result)
    if err != nil {
        panic(err)
    }

}

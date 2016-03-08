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
    err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
        &Person{"Cla", "+55 53 8402 8510"})
    if err != nil {
        panic(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).One(&result)
    if err != nil {
        panic(err)
    }

fmt.Println("Phone:", result.Phone)
}

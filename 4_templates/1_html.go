package main

import "os"
import "html/template"

func main() {

    t, err := template.ParseFiles("./template.gohtml")

    if err != nil {
        panic(err)
    }

    data := struct{
        Firstname string
        Lastname string
        Languages []string
    } {
        Firstname: "John",
        Lastname: "Doe",
        Languages: []string{"PHP", "GO"},
    }

    err = t.Execute(os.Stdout, data)

    if err != nil {
        panic(err)
    }
}

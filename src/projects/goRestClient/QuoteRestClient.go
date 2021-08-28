package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

    response, err := http.Get("http://quotes.rest/qod.json")

    if err != nil {
        fmt.Println(err)
        return
    }
    defer response.Body.Close() // this executes when function is exiting
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return // this will execute any defer statement that is stated in the code
    }
    fmt.Println(string(contents))
}

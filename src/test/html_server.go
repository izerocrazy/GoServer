package main

import (
    "fmt"
    "net/http"
    //"html"
)

func main() {
    HtmlServer := http.FileServer(http.Dir("."))
    http.Handle("/", HtmlServer)

    http.HandleFunc("/test", testFun)

    err := http.ListenAndServe(":8000", nil)
    CheckErr(err)
}

func testFun(w http.ResponseWriter, r *http.Request){
    //fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.RawQuery))
    fmt.Fprintf(w, "hello, %q", r.URL.RawQuery)
    fmt.Println(r.URL.Query())
}

func CheckErr(e error) {
    if e != nil {
        fmt.Println("error :", e.Error())
    }
}

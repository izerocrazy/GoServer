package main

import (
    "fmt"
    "net/http"
    //"Step"
    //"html/template"
    //"strconv"
    //"time"
    "encoding/json"
)

type HttpBaseRequest struct {
    retcode int
    msg string
    datetime int
}

func testFun(w http.ResponseWriter, r *http.Request){
    /* fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.RawQuery)) */
    fmt.Fprintf(w, "hello, %q", r.URL.RawQuery)
    fmt.Println(r.URL.Query())
}

type RegUserData struct {
    userid int
}

type RegUserStruct struct {
    /* HttpBaseRequest */
    retcode int
    msg string
    datetime int
    data RegUserData
}

func RegUser(w http.ResponseWriter, r *http.Request){
    szUserName := r.URL.Query()["username"][0]
    fmt.Fprintf(w, "hello, %s", szUserName)
    
    reg := RegUserStruct {
        retcode: 200,
        msg: "ok",
        datetime: 10,
        data: RegUserData { userid : 100 },
    }

    fmt.Println(reg)

    encode := json.NewEncoder(w)
    encode.Encode(reg) 
}

func main() {
    HtmlServer := http.FileServer(http.Dir("."))
    http.Handle("/", HtmlServer)
    http.HandleFunc("/test", testFun)

    http.HandleFunc("/user/reg", RegUser)
    //http.HandleFunc("/friend/add", AddFriend)
    //http.HandleFunc("/friend/list", GetFriendList)
    //http.HandleFunc("/yo/sendyo", SendYO)
    //http.HandleFunc("/yo/getyo", GetYO)    

    // 缺少一个默认的 404 

    err := http.ListenAndServe(":8000", nil)
    CheckErr(err)
}

func CheckErr(e error) {
    if e != nil {
        fmt.Println("error :", e.Error())
    }
}

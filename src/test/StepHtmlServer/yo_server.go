package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type HttpBaseRequest struct {
    retcode int
    msg string
    datetime int
}

func testFun(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "hello, %q", r.URL.RawQuery)
    fmt.Println(r.URL.Query())
}

type RegUserData struct {
    Userid int `json:"username"`
}

-- 思考怎么实现数据上的继承
type RegUserStruct struct {
    Retcode int `json:"retconde"`
    Msg string `json:"msg"`
    Datetime int `json:"datetime"`
    Data RegUserData `json:"data"`  
}

func RegUser(w http.ResponseWriter, r *http.Request){
    /* szUserName := r.URL.Query()["username"][0] */
    /* fmt.Fprintf(w, "hello, %s", szUserName) */
    
    reg := RegUserStruct {
        Retcode: 200,
        Msg: "ok",
        Datetime: 10,
        Data: RegUserData { Userid : 100 },
    }
    
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

package main

import (
    "Base"
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

///////////////////////////////////////////////////////
/* RegUse */
// 使用方法：客户端请求网址 /user/reg ，务必带上 username 参数
// 使用例子：http://localhost:8000/user/reg?username=1
type RegUserData struct {
    Userid int `json:"username"`
}

// 思考怎么实现数据上的继承
type IRegUser struct {
    Retcode int `json:"retconde"`
    Msg string `json:"msg"`
    Datetime int `json:"datetime"`
    Data RegUserData `json:"data"`  
}

func RegUser(w http.ResponseWriter, r *http.Request){
    /* szUserName := r.URL.Query()["username"][0] */
    /* fmt.Fprintf(w, "hello, %s", szUserName) */
    
    reg := IRegUser {
        Retcode: 200,
        Msg: "ok",
        Datetime: 10,
        Data: RegUserData { Userid : 100 },
    }
    
    encode := json.NewEncoder(w)
    encode.Encode(reg) 
}

///////////////////////////////////////////////////////
/* AddFriend */
// 使用方法：客户端请求网址 /friend/add ，务必带上 friendname 参数
// 使用例子：http://localhost:8000/friend/add?friendname=1
type AddFriendData struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

// 思考怎么实现数据上的继承
type IAddFriend struct {
    Retcode int `json:"retconde"`
    Msg string `json:"msg"`
    Datetime int `json:"datetime"`
    Data AddFriendData `json:"data"`  
}

func AddFriend(w http.ResponseWriter, r *http.Request){
    szUserName := r.URL.Query()["friendname"][0]
    /* fmt.Fprintf(w, "hello, %s", szUserName) */
    
    reg := IAddFriend {
        Retcode: 200,
        Msg: "ok",
        Datetime: 10,
        Data: AddFriendData { Id : 100 , Name : szUserName},
    }
    
    encode := json.NewEncoder(w)
    encode.Encode(reg) 
}


func main() {
    HtmlServer := http.FileServer(http.Dir("."))
    http.Handle("/", HtmlServer)
    http.HandleFunc("/test", testFun)

    http.HandleFunc("/user/reg", RegUser)
    http.HandleFunc("/friend/add", AddFriend)
    //http.HandleFunc("/friend/list", GetFriendList)
    //http.HandleFunc("/yo/sendyo", SendYO)
    //http.HandleFunc("/yo/getyo", GetYO)    

    // 缺少一个默认的 404 

    err := http.ListenAndServe(":8000", nil)
    Base.CheckErr(err)
}

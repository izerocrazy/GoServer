package HttpG

import (
        "net/http"
        "fmt"
        "os"
        "time"
        "strings"
        "net/url"
        "github.com/p/mahonia"
        "code.google.com/p/go.net/html"
        "encoding/json"
)

// because do not make ,waste my time 
var c chan int = make(chan int)

func CheckError(err error) {
    if err != nil {
        fmt.Println("Check Fatal error ", err.Error())
        os.Exit(1)
        //c <- 1
    }
}

func GetChannel() int{
    fmt.Println("waiting for channel...")
    nRetCode := <-c
    return nRetCode
}

func SendChannel(nRetCode int) {
    fmt.Println("send to channel", nRetCode)
    c <- nRetCode
}

func GetCharset(response *http.Response) string {
    contentType := response.Header.Get("Content-Type")
    if contentType == "" {
        // guess
        return "UTF-8"
    }
    idx := strings.Index(contentType, "charset:")
    if idx == -1 {
        // guess
        return "UTF-8"
    }
    return strings.Trim(contentType[idx:], " ")
}

type CompanyBaseInfo struct{
    SzCompanyName string
    ArrQylx []CompanyQylx
    SzZczb string
    ArrNswh []CompanyNswh
    ArrQyzz []CompanyQyzz
    ArrQyzzInfo [][]CompanyQyzzInfo
}

type CompanyQylx struct{
    SzName string
    SzEndTime string
}

type CompanyNswh struct {
    SzYear string
    SzMoney string
}

type CompanyQyzz struct {
    Qyzzid string
}

type CompanyQyzzInfo struct {
    Zzdj string
    ZznrName string
}

type XmyjBaseJson struct {
    Qymc string
    Xmmc string
    Zbtzsrq string
    Zbj string
    Htj string
    Jgysrq string
}

type XmyjQyzzJson struct{
    Zzmc string
    Zzdj string
}

type XmyjHjqk struct {
    Nd string
    Hjmc string
    Bjsj string
    Bjdw string
}

type Xmyj struct {
    Base XmyjBaseJson
    ArrQyzz []XmyjQyzzJson
    Plus map[string]string
    ArrHjqk []XmyjHjqk
}

func GetHttpResp(szUrl string) (*http.Response) {
    client := &http.Client{}

    request, err := http.NewRequest("GET", szUrl, nil)
    // only accept UTF-8
    request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
    CheckError(err)

	//CheckError(err)
	var response *http.Response
	for {
		response, err = client.Do(request)
		if err != nil {
            fmt.Println("Get url err wait 10 Second....")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

    if response.Status != "200 OK" {
        fmt.Println(response.Status)
        os.Exit(2)
    }

    chSet := GetCharset(response)
    //fmt.Printf("got charset %s\n", chSet)
    if chSet != "UTF-8" {
        fmt.Println("Cannot handle", chSet)
        os.Exit(4)
    }

    return response
}

func PostHttpResp(szUrl string, szPost *strings.Reader) (*http.Response) {
    client := &http.Client{}
    request, err := http.NewRequest("POST", szUrl, szPost);
    CheckError(err)

    request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    request.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
    request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,nl;q=0.2,zh-TW;q=0.2")
    request.Header.Add("Host", "www.gzzb.gd.cn")
    request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.149 Safari/537.36")

    var resp *http.Response
    for {
        resp, err = client.Do(request)
        if err != nil {
            fmt.Println("post url err wait 10 second")
            time.Sleep(10 * time.Second)
        } else {
            break
        }
    }
    //CheckError(err)

    chSet := GetCharset(resp)
    //fmt.Printf("got charset %s\n", chSet)
    if chSet != "UTF-8" {
        fmt.Println("Cannot handle", chSet)
        os.Exit(4)
    }

    return resp
}

func PostGzHttpJson(szUrl string, szService string, szArguments string, szFunc string) (*http.Response) {
    values := make(url.Values)

    values.Set("service", szService)
    values.Set("arguments", szArguments)
    values.Set("method", szFunc)

    szPost := strings.NewReader(values.Encode())
    //fmt.Println(szPost)
    return PostHttpResp(szUrl, szPost)
}

func GetCompanyQylxInfo(resp* http.Response) (CompanyBaseInfo){
    var cb CompanyBaseInfo

    r := resp.Body
    defer r.Close()
    doc, err := html.Parse(r)
    CheckError(err)

    var szTempQylxmc string
    var bGetQylxmc bool
    var szTempYxqz string
    var bGetYxqz bool

    var f func(*html.Node, bool, bool)
    f = func(n *html.Node, bFindDiv1 bool, bFindDiv2 bool) {
        if (n.Type == html.DocumentNode || n.Type == html.ElementNode) && n.Data == "div" {
            for _, a := range n.Attr {
                if a.Val == "qylxmc" {
                    bFindDiv1 = true
                } else if a.Val == "yxqz" {
                    bFindDiv2 = true
                }
            }
        }

        for c:= n.FirstChild; c != nil; c = c.NextSibling {
            if c.Type == html.TextNode {
                if bFindDiv1== true {
                    enc:=mahonia.NewDecoder("gbk")
                    //converts a  string from UTF-8 to gbk encoding.
                    szGbk := enc.ConvertString(c.Data) 

                    szTempQylxmc = szGbk
                    bGetQylxmc = true
                } else if bFindDiv2 == true {
                    szTempYxqz = c.Data
                    bGetYxqz = true
                }

                if bGetQylxmc && bGetYxqz {
                    var q CompanyQylx
                    q.SzName = szTempQylxmc
                    q.SzEndTime = szTempYxqz
                    cb.ArrQylx = append(cb.ArrQylx, q)

                    bGetQylxmc = false
                    bGetYxqz = false
                }
            }

            f(c, bFindDiv1, bFindDiv2)
        }

        if bFindDiv1 == true {
            bFindDiv1 = false
        }

        if bFindDiv2 == true {
            bFindDiv2 = false
        }
    }

    f(doc, false, false)

    return cb
}

func GetCompanyJczl(resp* http.Response) (string, string){
    r := resp.Body
    defer r.Close()

    type Cjson struct {
        Czzb    string
        Qymc    string
    }
    dec := json.NewDecoder(r)
    var c Cjson
    err := dec.Decode(&c)
    CheckError(err)

    return c.Qymc, c.Czzb
}

func GetCompanyNswh(resp* http.Response) ([]CompanyNswh) {
    r := resp.Body
    defer r.Close()
    type QylxData struct {
        Nsze string
        Nd  string
    }

    type Djson struct {
        Data []QylxData
    }

    dec := json.NewDecoder(r)
    var d Djson
    err := dec.Decode(&d)
    CheckError(err)

    var arrCn []CompanyNswh
    for _, a := range d.Data{
        if a.Nd == "2010" || a.Nd == "2011" || a.Nd == "2012"{
            //fmt.Println(a.Nd, a.Nsze)
            var cn CompanyNswh
            cn.SzYear = a.Nd
            cn.SzMoney = a.Nsze

            arrCn = append(arrCn, cn)
        }
    }

    return arrCn
}

func GetCompanyQyzz(resp* http.Response) ([]CompanyQyzz) {
    r := resp.Body
    defer r.Close()

    type QyzzData struct {
        Qyzzid string
    }

    type Djson struct {
        Data []QyzzData
    }

    dec := json.NewDecoder(r)
    var d Djson
    err := dec.Decode(&d)
    CheckError(err)

    var arrCn []CompanyQyzz
    for _, a := range d.Data {
        var cn CompanyQyzz
        cn.Qyzzid = a.Qyzzid

        arrCn = append(arrCn, cn)
    }

    return arrCn
}

func GetCompanyQyzzInfo(resp* http.Response) ([]CompanyQyzzInfo) {
    r := resp.Body
    defer r.Close()

    type Djson struct {
        Data []CompanyQyzzInfo
    }

    dec := json.NewDecoder(r)
    var d Djson
    err := dec.Decode(&d)
    CheckError(err)

    return d.Data
}

func CreateFileWithNameAddTitle(szFileName string, szTitleLine string) (file *os.File) {
    //file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777);
    file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777);
    CheckError(err)

    file.WriteString(szTitleLine)

    return file
}

func GetZzdj(n string) string {
    //StrMap := map[string]string{"01":"特级","02":"一级","03":"二级","04":"三级","05":"不分等级"}
    StrMap := map[string]string{"01":"特级","02":"一级","03":"二级","04":"三级","06":"甲","07":"乙","08":"丙","09":"暂乙级","12":"暂定级","13":"暂二级","14":"暂三级","17":"暂一级","21":"丁","10":"暂五级","05":"临时资质","20":"五级","19":"四级","11":"暂四级","15":"不分等级","16":"暂甲级","18":"暂丙级","22":"预备级"}


    return StrMap[n]
}


package main

import (
        "fmt"
        "net/http"
        "net/url"
        "os"
        "strings"
        "code.google.com/p/go.net/html"
        "github.com/p/mahonia"
        "time"
        "strconv"
)

var nStringMap = make(map[int][]string)
var szQybhUrlMap = make(map[string]string)
var szCompanyChenxinMap = make(map[string][]string)

func main() {
    gzgcjg := "http://www.gzgcjg.com/gzqypjtx/Login.aspx"
    gzgcjg2 := "http://www.gzgcjg.com/gzqypjtx/common/LoginYbhnt.aspx"
    gzgcjg3 := "http://www.gzgcjg.com/gzqypjtx/common/LoginYllh.aspx"

    FilterBody(GetHttpResp(gzgcjg), false, "")
    FilterBody(GetHttpResp(gzgcjg2), true, "div_2")
    FilterBody(GetHttpResp(gzgcjg3), true, "div_yllh")

	fmt.Println("[INFO] Get All Campany Name !")
	
    //ShowReader(PostHttpResp(url4.String(), 1, nStringMap[1][1]))
    for key, value := range nStringMap{
        fmt.Println(key)
        for key2, value2 := range value {
            fmt.Println(key2, "Get Company ", value2, "Base Info Url")
            gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"
            FilterBody2(PostHttpResp(gzzb, key, value2), value2)
            time.Sleep(10 * time.Second)
        }
    }

    for key, value := range szQybhUrlMap {
        fmt.Println("Get Company", key, "Data Info")
        s := []string{"http://www.gzzb.gd.cn/", value}
        szUrl := strings.Join(s, "");
        FilterBody3(GetHttpResp(szUrl), key);
        time.Sleep(10 * time.Second)
    }

    //fmt.Println(szCompanyChenxinMap)
    file, err := os.OpenFile("result.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777);
    checkError(err)
    defer file.Close()
    for key, value := range szCompanyChenxinMap {
        s := []string{key}
        for _, value2 := range value {
            s = append(s, value2)
        }
        szLine := strings.Join(s, "\t");
        szLine = szLine + "\r\n"
        file.WriteString(szLine)
    }

    //fmt.Println(nStringMap)
}

func GetHttpResp(szUrl string) (*http.Response) {
    client := &http.Client{}

    request, err := http.NewRequest("GET", szUrl, nil)
    // only accept UTF-8
    request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
    checkError(err)

    
	//checkError(err)
	var response *http.Response
	for {
		response, err = client.Do(request)
		if err != nil {
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}
	
    if response.Status != "200 OK" {
        fmt.Println(response.Status)
        os.Exit(2)
    }

    chSet := getCharset(response)
    //fmt.Printf("got charset %s\n", chSet)
    if chSet != "UTF-8" {
        fmt.Println("Cannot handle", chSet)
        os.Exit(4)
    }

    return response
}

func PostHttpResp(szUrl string, nSelTypeId int, szQymc string) (*http.Response) {
    client := &http.Client{}
    values := make(url.Values)

    //cd, err := iconv.Open("gbk", "utf-8")
    //checkError(err)
    //defer cd.Close()
	//szGbk := cd.ConvString(szQymc)
	
	enc:=mahonia.NewEncoder("gbk")
	//converts a  string from UTF-8 to gbk encoding.
	szGbk := enc.ConvertString(szQymc) 

    values.Set("qyxx_qymc", szGbk)
    szSelTypeId := "0"
    if nSelTypeId > 9 {
        szSelTypeId = strconv.Itoa(nSelTypeId)
    } else {
        s := []string{"0", strconv.Itoa(nSelTypeId)}
        szSelTypeId = strings.Join(s, "");
    }
    values.Set("qyxx_qylx", szSelTypeId)

    request, err := http.NewRequest("POST", szUrl, strings.NewReader(values.Encode()))
    checkError(err)
    //fmt.Println(values.Encode())

    request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    request.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
    request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,nl;q=0.2,zh-TW;q=0.2")
    request.Header.Add("Host", "www.gzzb.gd.cn")
    request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.149 Safari/537.36")

    resp, err := client.Do(request)
    checkError(err)

    chSet := getCharset(resp)
    fmt.Printf("got charset %s\n", chSet)
    if chSet != "UTF-8" {
        fmt.Println("Cannot handle", chSet)
        os.Exit(4)
    }

    return resp
}

func FilterDivValue(s string) bool {
    arr := []string{"myTab_div1","myTab_div2","myTab_div3","myTab_divRight1","myTab_divRight2","myTab_divRight3","myTab_divRight4","div_zjzx"}
    for _, a := range arr {
        if s == a {
            //fmt.Println(s)
            return true
        }
    }

    return false
}

func GetStringToInt(s string) int {
    StrMap := map[string]int{"myTab_div1": 1, "myTab_div2": 1, "myTab_div3": 1, "myTab_divRight1": 2, "myTab_divRight2": 2, "myTab_divRight3": 2, "myTab_divRight4": 5, "div_yllh": 12, "div_2": 6, "div_zjzx": 9}

    return StrMap[s]
}

func FilterBody3(resp* http.Response, szCompanyName string) {
    r := resp.Body
    defer r.Close()
    doc, err := html.Parse(r)
    checkError(err)

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
                    //cd, err := iconv.Open("utf-8", "gbk")
                    //checkError(err)
                    //defer cd.Close()
                    //szGbk := cd.ConvString(c.Data)
					
					enc:=mahonia.NewDecoder("gbk")
					//converts a  string from UTF-8 to gbk encoding.
					szGbk := enc.ConvertString(c.Data) 
					
                    szCompanyChenxinMap[szCompanyName] = append(szCompanyChenxinMap[szCompanyName], szGbk)
					
					fmt.Println("Name: ", szGbk)
                } else if bFindDiv2 == true {
                    szCompanyChenxinMap[szCompanyName] = append(szCompanyChenxinMap[szCompanyName], c.Data)
					fmt.Println("Data: ", c.Data)
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
}

func FilterBody2(resp *http.Response, szCompanyName string) {
    r := resp.Body
    defer r.Close()
    doc, err := html.Parse(r)
    checkError(err)

    var f func(*html.Node, bool)
    f = func(n *html.Node, bFindDiv bool) {
        if (n.Type == html.DocumentNode || n.Type == html.ElementNode) && n.Data == "div" {
            for _, a := range n.Attr {
                if a.Val == "bszn_right_table" {
                    bFindDiv = true
                }
            }
        }

        for c:= n.FirstChild; c != nil; c = c.NextSibling {
            if (bFindDiv == true && c.Data == "a") {
                for _, a := range c.Attr {
                    if a.Key == "href"{
                        bFindDiv = false;
                        szQybhUrlMap[szCompanyName] = a.Val
                        
						fmt.Println("url:", a.Val)
                    }
                }
            }

            f(c, bFindDiv)
        }

        if bFindDiv == true {
            bFindDiv = false
        }
    }

    f(doc, false)
}

func FilterBody(resp *http.Response, bFindDiv bool, szDivName string) {
    r := resp.Body
    defer r.Close()
    doc, err := html.Parse(r)
    checkError(err)

    var f func(*html.Node, bool, string)
    f = func(n *html.Node, bFindDiv bool, szDivName string) {
        bFind1 := false
        if (n.Type == html.DocumentNode || n.Type == html.ElementNode) && n.Data == "td" {
            for _, a := range n.Attr {
                if a.Val == "gridview_itemStyle" {
                    //fmt.Println(a.Key)
                    bFind1 = true
                }
            }
        }else if (n.Type == html.ElementNode && n.Data == "div") {
            for _, a := range n.Attr {
                if FilterDivValue(a.Val) {
                    szDivName = a.Val
                    bFindDiv = true;
                }
            }
        }

        for c:= n.FirstChild; c != nil; c = c.NextSibling {
            if bFind1 == true && bFindDiv == true && len(c.Data) > 6 && c.Type == html.TextNode {
                bFind1 = false

                nDivId := GetStringToInt(szDivName)
                nStringMap[nDivId] = append(nStringMap[nDivId], c.Data)
				
				fmt.Println("Get Company Name: ", c.Data)
            }

            f(c, bFindDiv, szDivName)
        }

        if bFindDiv == true {
            bFindDiv = false
        }
    }

    f(doc, bFindDiv, szDivName)
}

func ShowReader(resp *http.Response) {
    r := resp.Body
    defer resp.Body.Close()

    var buf [512]byte
    reader := r
    //fmt.Println("got body")
    for {
        n, err := reader.Read(buf[0:])
        if err != nil {
            os.Exit(0)
        }

        //cd, err := iconv.Open("gbk", "utf-8")
        //checkError(err)
        //defer cd.Close()

        //szGbk := cd.ConvString(string(buf[0:n]))

        //fmt.Print(szGbk)
        fmt.Println(string(buf[0:n]))
    }

    os.Exit(0)
}

func getCharset(response *http.Response) string {
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

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}

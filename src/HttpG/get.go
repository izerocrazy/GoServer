package main

import (
        "fmt"
        "net/http"
        "net/url"
        "os"
        "strings"
        "io"
        "code.google.com/p/go.net/html"
        "github.com/qiniu/iconv"
)

func main() {
    gzgcjg := "http://www.gzgcjg.com/gzqypjtx/Login.aspx"
    //gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"

    url, err := url.Parse(gzgcjg)
    checkError(err)

    //ShowReader(GetHttpBady(url.String()))
    FilterBody(GetHttpBady(url.String()))
    //ShowReader(PostHttpBady(url.String()))
}

func GetHttpBady(szUrl string) (io.ReadCloser) {
    client := &http.Client{}

    request, err := http.NewRequest("GET", szUrl, nil)
    // only accept UTF-8
    request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
    checkError(err)

    response, err := client.Do(request)
    if response.Status != "200 OK" {
        fmt.Println(response.Status)
        os.Exit(2)
    }

    chSet := getCharset(response)
    fmt.Printf("got charset %s\n", chSet)
    if chSet != "UTF-8" {
        fmt.Println("Cannot handle", chSet)
        os.Exit(4)
    }

    return response.Body
}

func PostHttpBady(szUrl string, szQymc string) (io.ReadCloser) {
    client := &http.Client{}
    values := make(url.Values)
    values.Set("qyxx_qymc", szQymc)
    values.Set("qyxx_qylx", "02")

    request, err := http.NewRequest("POST", szUrl, strings.NewReader(values.Encode()))
    fmt.Println(values.Encode())
    checkError(err)

    request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    request.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
    request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,nl;q=0.2,zh-TW;q=0.2")
    request.Header.Add("Host", "www.gzzb.gd.cn")
    request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.149 Safari/537.36")

    resp, err := client.Do(request)
    checkError(err)

    //resp.Body.Close()
    chSet := getCharset(resp)
    fmt.Printf("got charset %s\n", chSet)
    if chSet != "UTF-8" {
        fmt.Println("Cannot handle", chSet)
        os.Exit(4)
    }

    return resp.Body
}

func FilterBody(r io.ReadCloser) {
    doc, err := html.Parse(r)
    checkError(err)

    var f func(*html.Node, bool)
    f = func(n *html.Node, bFindDiv bool) {
        bFind1 := false
        bFind2 := true
        if (n.Type == html.DocumentNode || n.Type == html.ElementNode) && n.Data == "td" {
            for _, a := range n.Attr {
                if a.Val == "gridview_itemStyle" {
                    //fmt.Println(a.Key)
                    bFind1 = true
                } else if(a.Val == "width:242px;") {
                    //fmt.Println(a.Key)
                    bFind2 = true
                }
            }
        }else if (n.Type == html.ElementNode && n.Data == "div") {
            for _, a := range n.Attr {
                if a.Val == "myTab_divRight1" {
                    bFindDiv = true;
                }
            }
        }

        for c:= n.FirstChild; c != nil; c = c.NextSibling {
            if bFind1 == true && bFind2 == true && bFindDiv == true && len(c.Data) > 6 && c.Type == html.TextNode {
                bFind1 = false
                bFind2 = false
                //fmt.Println(len(c.Data))
                fmt.Println(c.Data)

                gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"
                url, err := url.Parse(gzzb)
                checkError(err)

                ShowReader(PostHttpBady(url.String(), c.Data))

                os.Exit(0)
            }

            f(c, bFindDiv)
        }

        if bFindDiv == true {
            bFindDiv = false
        }
    }

    f(doc,false)
}

func ShowReader(r io.ReadCloser) {
    var buf [512]byte
    reader := r
    fmt.Println("got body")
    for {
        n, err := reader.Read(buf[0:])
        if err != nil {
            os.Exit(0)
        }
        fmt.Print(string(buf[0:n]))
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

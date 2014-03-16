package main

import (
        "fmt"
        "net/http"
        "net/url"
        "os"
        "strings"
        "io"
        "code.google.com/p/go.net/html"
)

func main() {
    gzgcjg := "http://www.gzgcjg.com/gzqypjtx/Login.aspx"

    url, err := url.Parse(gzgcjg)
    checkError(err)

    //ShowReader(GetHttpBady(url.String()))
    FilterBody(GetHttpBady(url.String()))
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

func FilterBody(r io.ReadCloser) {
    doc, err := html.Parse(r)
    checkError(err)

    var f func(*html.Node)
    f = func(n *html.Node) {
        bFind := false
        if (n.Type == html.DocumentNode || n.Type == html.ElementNode) && n.Data == "td" {
            for _, a := range n.Attr {
                if a.Val == "gridview_itemStyle" {
                    //fmt.Println(a.Key)
                    bFind = true
                    break
                }
            }
        }
        for c:= n.FirstChild; c != nil; c = c.NextSibling {
            if bFind == true && c.Type == html.TextNode {
                bFind = false
                fmt.Println(c.Data)
            }
            f(c)
        }
    }

    f(doc)
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

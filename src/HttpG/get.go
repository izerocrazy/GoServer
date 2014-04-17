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
        "encoding/json"
)

var nStringMap = make(map[int][]string)
var szQybhUrlMap = make(map[string][]string)
var szCompanyChenxinMap = make(map[string]CompanyBaseInfo)
var szXmyjMap = make(map[string]Xmyj)

type CompanyBaseInfo struct{
    szCompanyName string
    arrQylx []CompanyQylx
    szZczb string
    arrNswh []CompanyNswh
}

type CompanyQylx struct{
    szName string
    szEndTime string
}

type CompanyNswh struct {
    szYear string
    szMoney string
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

func ShowIntStringTable() (int){
    var id int
    for {
        fmt.Println("请选择你需要查询企业类别对应的数字：\n")
        fmt.Println("0：全取")
        fmt.Println("1：施工企业排名")
        fmt.Println("2：施工企业-市政排名")
        fmt.Println("3：施工企业-房建排名")
        fmt.Println("4：监理单位排名")
        fmt.Println("5：监理单位-市政排名")
        fmt.Println("6：监理单位-房建排名")
        fmt.Println("7：招标代理排名")
        fmt.Println("8：园林绿化")
        fmt.Println("9：预拌混凝土")
        fmt.Println("10：造价咨询")
        fmt.Print("请输入：")

        fmt.Scanf("%d", &id)

        if id > -1 && id < 11 {
            break
        }
    }

    return id
}

func GetCodeForTest() {
    s := "http://www.gzzb.gd.cn/qyww/json";
    szArguments := fmt.Sprintf("[\"[{\\\"tableName\\\":\\\"basic_code\\\",\\\"valueField\\\":\\\"code\\\",\\\"textField\\\":\\\"name\\\",\\\"where\\\":\\\"code_type=\\\\\\'QY_YJ_ZZDJ\\\\\\'\\\",\\\"eleid\\\":\\\"zzdj\\\"}]\"]")
    ShowReader(PostHttpResp2(s, "CodeBS", szArguments, "findCode"))
}

func main() {
    gzgcjg := "http://www.gzgcjg.com/gzqypjtx/Login.aspx"
    gzgcjg2 := "http://www.gzgcjg.com/gzqypjtx/common/LoginYbhnt.aspx"
    gzgcjg3 := "http://www.gzgcjg.com/gzqypjtx/common/LoginYllh.aspx"
    //gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"

    //ShowReader(GetHttpResp(url.String()))
    fmt.Println("Program Init...")
    for {
        FilterBody(GetHttpResp(gzgcjg), false, "")
        fmt.Println("Wait 5 Second...", gzgcjg)
        time.Sleep(5 * time.Second)
        FilterBody(GetHttpResp(gzgcjg2), true, "div_2")
        fmt.Println("Wait 5 Second...", gzgcjg2)
        time.Sleep(5 * time.Second)
        FilterBody(GetHttpResp(gzgcjg3), true, "div_yllh")
        fmt.Println("Wait 5 Second...", gzgcjg3)
        time.Sleep(5 * time.Second)

        id := ShowIntStringTable()

        //ShowReader(PostHttpResp(url4.String(), 1, nStringMap[1][1]))
	//fmt.Println(nStringMap)
        for key, value := range nStringMap {
            //fmt.Println(key)
            if (id == 0 || key == id){
                for _, value2 := range value {
                    fmt.Println("已载入 "+ value2)
                    gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"
                    selKey := GetIndexKey(key)
                    FilterBody2(PostHttpResp(gzzb, selKey, value2), value2)
                    fmt.Println("Wait 10 Second...")
                    time.Sleep(10 * time.Second)
                    //break // __________ debug
                }
                //break //__________ debug
            }
        }

        for key, value := range szQybhUrlMap {
            fmt.Println("已载入 "+key)
            arrS := []string{"http://www.gzzb.gd.cn/", value[0]}
            szUrl := strings.Join(arrS, "");

            cb := FilterBody3(GetHttpResp(szUrl), key);

            fmt.Println("Wait 10 Second...")
            time.Sleep(10 * time.Second)

            //value[1] = "10020" // __________ debug

            s := "http://www.gzzb.gd.cn/qyww/json";
            szArguments := fmt.Sprintf("[\"%s\"]", value[1])
            cb.szZczb = FilterJson(PostHttpResp2(s, "TQyQyjczlBS", szArguments, "findQyjczl"))

            szArguments = fmt.Sprintf("[0,10,\"%s\"]", value[1])
            cb.arrNswh = FilterJson2(PostHttpResp2(s, "TQyQynswhBS", szArguments, "findTQyQynswhInfo"))

            // only can be set at the end of process 
            szCompanyChenxinMap[key] = cb

            szArguments = fmt.Sprintf("[0,10,\"%s\"]", value[1])
            arrId := FilterJson3(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findTQyXmyjInfo"), value[1], true)  // _______________ debug
            //fmt.Println(arrId)
            //arrId := FilterJson3(PostHttpResp2(s, "JyBlzbtzsBS", szArguments, "findJyBlzbtzsInfox"), value[1], true)

            for _, szId := range arrId {
                //szId = "yj_80f11ebf-be47-4d80-9ea8-e30c3d20831c" // ____________ debug
                fmt.Println("Wait 10 Second...")
                time.Sleep(5 * time.Second)
                fmt.Println("已载入 "+szId)
                szArguments := fmt.Sprintf("[{\"xmyjid\":\"%s\"}]", szId)
                qBase := FilterJson_XmyjBase(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findQyyj"))
                fmt.Println(qBase)
                //ShowReader(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findQyyj"))

                szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szId)
                //ShowReader(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findQyzz"))
                arrQyzz := FilterJson_XmyjQyzz(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findQyzz"))
                //fmt.Println(arrQyzz)

                szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szId)
                qXmgm := FilterJson_XmyjPlus(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findXmgm"))
                //fmt.Println(qXmgm)

                szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szId)
                qHjqk := FilterJson_XmyjHjqk(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findHjqk"))
                //fmt.Println(qHjqk)

                szXmyjMap[szId] = Xmyj{
                    Base: qBase,
                    ArrQyzz: arrQyzz,
                    Plus: qXmgm,
                    ArrHjqk: qHjqk,
                }
                //break // _____________ debug
            }
            //break // ________________ debug
        }
        //fmt.Println(szCompanyChenxinMap)
        //fmt.Println(szXmyjMap)

        SaveChenxin()
        SaveXmyj()
        SaveHjqk()
    }
}

func SaveChenxin(){
    file, err := os.OpenFile("1.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777);
    checkError(err)
    defer file.Close()

    szTitleLine := "企业名称\t注册资本(万元) \t2010纳税\t2011纳税\t2012纳税\t企业类型\t有效期\r\n"
    file.WriteString(szTitleLine)

    for key, value := range szCompanyChenxinMap {
        //fmt.Println(value)
        s := []string{key}
        s = append(s, value.szZczb)

        szNs2010 := ""
        szNs2011 := ""
        szNs2012 := ""
        for _, szNs := range value.arrNswh {
            if szNs.szYear == "2010" {
                szNs2010 = szNs.szMoney
            } else if szNs.szYear == "2011" {
                szNs2011 = szNs.szMoney
            } else if szNs.szYear == "2012" {
                szNs2012 = szNs.szMoney
            }
        }
        s = append(s, szNs2010)
        s = append(s, szNs2011)
        s = append(s, szNs2012)

        if len(value.arrQylx) > 0 {
            for _, value2 := range value.arrQylx {
                s1 := append(s, value2.szName)
                s1 = append(s1, value2.szEndTime)

                szLine := strings.Join(s1, "\t")
                szLine = szLine + "\r\n"
                file.WriteString(szLine)
            }
        } else {
            szLine := strings.Join(s, "\t")
            szLine = szLine + "\r\n"
            file.WriteString(szLine)
        }
    }
}

func SaveXmyj() {
    file, err := os.OpenFile("2.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
    checkError(err)
    defer file.Close()

    szTitleLine := "企业名称\t项目名称\t中标日期\t中标价（万元）\t合同价（万元）\t竣工验收日期\t工程资质\t工程等级\r\n"
    file.WriteString(szTitleLine)

    for _, value := range szXmyjMap{
        //s := []string{key}
        s := []string{}
        s = append(s, value.Base.Qymc)
        s = append(s, value.Base.Xmmc)
        s = append(s, value.Base.Zbtzsrq)
        s = append(s, value.Base.Zbj)
        s = append(s, value.Base.Htj)
        s = append(s, value.Base.Jgysrq)

        szPlus := ""
        for key, value := range value.Plus {
            szPlus = szPlus + key + "\t"
            szPlus = szPlus + value + "\t"
        }

        if len(value.ArrQyzz) > 0 {
            for _, value2 := range value.ArrQyzz{
                s1 := append(s, value2.Zzmc)
                s1 = append(s1, getZzdj(value2.Zzdj))
                szLine := strings.Join(s1, "\t");
                szLine = szLine + "\t" + szPlus
                szLine = szLine + "\r\n"
                file.WriteString(szLine)
            }
        } else {
            szLine := strings.Join(s, "\t");
            szLine = szLine + "\t" + szPlus
            szLine = szLine + "\r\n"
            file.WriteString(szLine)
        }
    }
}

func SaveHjqk() {
    file, err := os.OpenFile("3.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
    checkError(err)
    defer file.Close()

    szTitleLine := "企业名称\t项目名称\t年度\t奖项\t颁奖时间\t颁奖单位\r\n"
    file.WriteString(szTitleLine)
    
    for _, value := range szXmyjMap{
        s := []string{}
        s = append(s, value.Base.Qymc)
        s = append(s, value.Base.Xmmc)

        for _, value2 := range value.ArrHjqk{
            s1 := append(s, value2.Nd)
            s1 = append(s1, getHjmc(value2.Hjmc))
            s1 = append(s1, value2.Bjsj)
            s1 = append(s1, value2.Bjdw)

            szLine := strings.Join(s1, "\t");
            szLine = szLine + "\r\n"
            file.WriteString(szLine)
        }
    }
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
            fmt.Println("Get url err wait 10 Second....")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}
	
    if response.Status != "200 OK" {
        //fmt.Println(response.Status)
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

func PostHttpResp2(szUrl string, szService string, szArguments string, szFunc string) (*http.Response) {
    values := make(url.Values)

    values.Set("service", szService)
    values.Set("arguments", szArguments)
    values.Set("method", szFunc)

    szPost := strings.NewReader(values.Encode())
    //fmt.Println(szPost)
    return _PostHttpResp(szUrl, szPost)
}

func PostHttpResp(szUrl string, nSelTypeId int, szQymc string) (*http.Response) {
	enc:=mahonia.NewEncoder("gbk")
	//converts a  string from UTF-8 to gbk encoding.
	szGbk := enc.ConvertString(szQymc) 
    //cd, err := iconv.Open("gbk", "utf-8")
    //checkError(err)
    //defer cd.Close()
	//szGbk := cd.ConvString(szQymc)
	
	values := make(url.Values)
    values.Set("qyxx_qymc", szGbk)
    szSelTypeId := "0"
    if nSelTypeId > 9 {
        szSelTypeId = strconv.Itoa(nSelTypeId)
    } else {
        s := []string{"0", strconv.Itoa(nSelTypeId)}
        szSelTypeId = strings.Join(s, "");
    }
    values.Set("qyxx_qylx", szSelTypeId)

    szPost := strings.NewReader(values.Encode())
    return _PostHttpResp(szUrl, szPost);
}

func _PostHttpResp(szUrl string, szPost *strings.Reader) (*http.Response) {
    client := &http.Client{}
    request, err := http.NewRequest("POST", szUrl, szPost);
    checkError(err)

    request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
    request.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
    request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,nl;q=0.2,zh-TW;q=0.2")
    request.Header.Add("Host", "www.gzzb.gd.cn")
    request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.149 Safari/537.36")

    resp, err := client.Do(request)
    checkError(err)

    chSet := getCharset(resp)
    //fmt.Printf("got charset %s\n", chSet)
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

func GetDivNameIndex(s string) int {
    StrMap := map[string]int{"myTab_div1": 1, "myTab_div2": 2, "myTab_div3": 3, "myTab_divRight1": 4, "myTab_divRight2": 5, "myTab_divRight3": 6, "myTab_divRight4": 7, "div_yllh": 8, "div_2": 9, "div_zjzx": 10}

    return StrMap[s]
}

func GetIndexKey(nIndex int) int {
    IndexKeyMap := map[int]int{1:1, 2:1, 3:1, 4:2, 5:2, 6:2, 7:5, 8:12, 9:6, 10:9}

    return IndexKeyMap[nIndex]
}

func FilterJson_XmyjHjqk(resp* http.Response) ([]XmyjHjqk) {
    r := resp.Body
    defer r.Close()

    type XmyjHjqkJson struct {
        Data []XmyjHjqk
    }

    dec := json.NewDecoder(r)
    var d XmyjHjqkJson
    err := dec.Decode(&d)
    checkError(err)

    return d.Data
}

func FilterJson_XmyjPlus(resp* http.Response) (map[string]string) {
    r := resp.Body
    defer r.Close()

    type XmyjPlusOne struct {
        Gmzb string
        Sl string
        Dw string
    }

    type XmyjPlusJson struct {
        Data []XmyjPlusOne
    }

    dec := json.NewDecoder(r)
    var d XmyjPlusJson
    err := dec.Decode(&d)
    checkError(err)

    ret := make(map[string]string)
    for _, value := range d.Data {
        /*szGbk := value.Gmzb
        if szGbk == "工程造价" {
            ret.Gczj = value.Sl
        } else if szGbk == "地上层数" {
            ret.Dscs = value.Sl
        } else if szGbk == "地下层数" {
            ret.Dxcs = value.Sl
        } else if szGbk == "建筑面积" {
            ret.Jzmj = value.Sl
        }*/
        ret[value.Gmzb] = value.Sl
    }

    return ret
}

func FilterJson_XmyjQyzz(resp* http.Response) ([]XmyjQyzzJson){
    r := resp.Body
    defer r.Close()

    type QyyjQyzzJson struct {
        Data []XmyjQyzzJson
    }

    dec := json.NewDecoder(r)
    var d QyyjQyzzJson
    err := dec.Decode(&d)
    checkError(err)

    return d.Data
}

func FilterJson_XmyjBase(resp* http.Response) (XmyjBaseJson){
    r := resp.Body
    defer r.Close()

    dec := json.NewDecoder(r)
    var d XmyjBaseJson
    err := dec.Decode(&d)
    checkError(err)

    return d
}

func FilterJson3(resp* http.Response, szId string, bNeedPage bool) ([]string) {
    r := resp.Body
    defer r.Close()

    type QyzbInfoJson struct {
        Xmyjid string
    }

    type QyzbInfoJsons struct {
        Total string
        Data []QyzbInfoJson
    }

    dec := json.NewDecoder(r)
    var d QyzbInfoJsons
    err := dec.Decode(&d)
    checkError(err)

    var arrId []string

    for _, a := range d.Data{
        //fmt.Println(a.Xmyjid)
        arrId = append(arrId, a.Xmyjid)
    }

    if bNeedPage == true {
        nTotal, err := strconv.Atoi(d.Total)
        checkError(err)

        nPage := nTotal / 10
        if nPage > 1 {
            for i := 1; i <= nPage; i ++ {
                fmt.Println("Wait 10 Second...")
                time.Sleep(10 * time.Second)
                s := "http://www.gzzb.gd.cn/qyww/json";
                szArguments := fmt.Sprintf("[%d,10,\"%s\"]", i * 10, szId)
                arrId2 := FilterJson3(PostHttpResp2(s, "TQyXmyjBS", szArguments, "findTQyXmyjInfo"), szId, false)
                for _, v := range arrId2 {
                    arrId = append(arrId, v)
                }

                //break //______ debug
            }
        }
    }
    return arrId
}

func FilterJson2(resp* http.Response) ([]CompanyNswh) {
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
    checkError(err)

    var arrCn []CompanyNswh
    for _, a := range d.Data{
        if a.Nd == "2010" || a.Nd == "2011" || a.Nd == "2012"{
            //fmt.Println(a.Nd, a.Nsze)
            var cn CompanyNswh
            cn.szYear = a.Nd
            cn.szMoney = a.Nsze

            arrCn = append(arrCn, cn)
        }
    }

    return arrCn
}

func FilterJson(resp* http.Response) (string){
    r := resp.Body
    defer r.Close()

    type Cjson struct {
        Czzb    string
    }
    dec := json.NewDecoder(r)
    var c Cjson
    err := dec.Decode(&c)
    checkError(err)

    return c.Czzb
}

func FilterBody3(resp* http.Response, szCompanyName string) (CompanyBaseInfo){
    var cb CompanyBaseInfo
    cb.szCompanyName = szCompanyName

    r := resp.Body
    defer r.Close()
    doc, err := html.Parse(r)
    checkError(err)

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
                    q.szName = szTempQylxmc
                    q.szEndTime = szTempYxqz
                    cb.arrQylx = append(cb.arrQylx, q)

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
                    if a.Key == "href" {
                        bFindDiv = false;
                        szQybhUrlMap[szCompanyName] = append(szQybhUrlMap[szCompanyName], a.Val)
                        //fmt.Println(szQybhUrlMap)
                        arr := strings.Split(a.Val, "=")
                        szQybhUrlMap[szCompanyName] = append(szQybhUrlMap[szCompanyName], arr[1])
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

                nDivId := GetDivNameIndex(szDivName)
                nStringMap[nDivId] = append(nStringMap[nDivId], c.Data)
		//fmt.Println("Get Company Name: ", c.Data)
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
            break
        }

        //cd, err := iconv.Open("gbk", "utf-8")
        //checkError(err)
        //defer cd.Close()

        //szGbk := cd.ConvString(string(buf[0:n]))

        //fmt.Print(szGbk)
        fmt.Println(string(buf[0:n]))
    }

    //os.Exit(0)
}

func getHjmc(n string) string {
    StrMap := map[string]string{"01":"中国建设工程鲁班奖（国家优质工程）","02":"全国市政金杯示范工程","03":"国家优质工程（金质奖）","04":"国家优质工程（银质奖）","05":"广东省建设工程金匠奖","06":"全国建筑工程装饰奖","07":"广州地区建设工程质量“五羊杯”","08":"广州市优良样板工程","09":"广州市安全文明施工样板工地（市双优）","10":"广东省房屋市政工程安全生产文明施工示范工地（原广东省安全文明施工样板工地）","11":"广东省建设工程优质奖（原省优良样板工程）","12":"广州市市政优良样板工程","13":"中国土木工程詹天佑奖","14":"全国建筑业新技术应用示范工程执行单位","15":"广东省建筑业新技术应用示范工程执行单位","16":"广东省优秀建筑装饰工程奖","17":"广州市建筑装饰优质工程奖","18":"广州市建设工程（市政）质量“五羊杯”"}
    
    return StrMap[n]            
}

func getZzdj(n string) string {
    StrMap := map[string]string{"01":"特级","02":"一级","03":"二级","04":"三级","05":"不分等级"}

    return StrMap[n]
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

package main

import (
        "HttpG"
        "fmt"
        "time"
        "os"
        "strings"
        "net/http"
        "net/url"
        "code.google.com/p/go.net/html"
        "github.com/p/mahonia"
        "strconv"
        "encoding/json"
)

var nStringMap = make(map[int][]string)
var szQybhUrlMap = make(map[string][]string)
var szCompanyChenxinMap = make(map[string]HttpG.CompanyBaseInfo)
var szXmyjMap = make(map[string]HttpG.Xmyj)

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
    // QY_YJ_ZZDJ
    s := "http://www.gzzb.gd.cn/qyww/json";
    szArguments := fmt.Sprintf("[\"[{\\\"tableName\\\":\\\"basic_code\\\",\\\"valueField\\\":\\\"code\\\",\\\"textField\\\":\\\"name\\\",\\\"where\\\":\\\"code_type=\\\\\\'QY_ZZDJ\\\\\\'\\\",\\\"eleid\\\":\\\"zzdj\\\"}]\"]")
    ShowReader(HttpG.PostGzHttpJson(s, "CodeBS", szArguments, "findCode"))
}

func _main() {
    GetCodeForTest()
}

func main() {
    gzgcjg := "http://www.gzgcjg.com/gzqypjtx/Login.aspx"
    gzgcjg2 := "http://www.gzgcjg.com/gzqypjtx/common/LoginYbhnt.aspx"
    gzgcjg3 := "http://www.gzgcjg.com/gzqypjtx/common/LoginYllh.aspx"
    //gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"

    //ShowReader(HttpG.GetHttpResp(url.String()))
    fmt.Println("Program Init...")
    for {
        FilterBody(HttpG.GetHttpResp(gzgcjg), false, "")
        fmt.Println("Wait 5 Second...", gzgcjg)
        time.Sleep(5 * time.Second)
        FilterBody(HttpG.GetHttpResp(gzgcjg2), true, "div_2")
        fmt.Println("Wait 5 Second...", gzgcjg2)
        time.Sleep(5 * time.Second)
        FilterBody(HttpG.GetHttpResp(gzgcjg3), true, "div_yllh")
        fmt.Println("Wait 5 Second...", gzgcjg3)
        time.Sleep(5 * time.Second)

        id := ShowIntStringTable()

        //ShowReader(PostHttpResp1(url4.String(), 1, nStringMap[1][1]))
        //fmt.Println(nStringMap)
        for key, value := range nStringMap {
            //fmt.Println(key)
            if (id == 0 || key == id){
                for _, value2 := range value {
                    fmt.Println("已载入 "+ value2)
                    gzzb := "http://www.gzzb.gd.cn/cms/wz/view/sccx/QyxxServlet?siteId=1"
                    selKey := GetIndexKey(key)
                    FilterBody2(PostHttpResp1(gzzb, selKey, value2), value2)
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

            cb := HttpG.GetCompanyQylxInfo(HttpG.GetHttpResp(szUrl));
            cb.SzCompanyName = key;

            fmt.Println("Wait 10 Second...")
            time.Sleep(10 * time.Second)

            //value[1] = "10020" // __________ debug

            s := "http://www.gzzb.gd.cn/qyww/json";
            szArguments := fmt.Sprintf("[\"%s\"]", value[1])
            _, cb.SzZczb = HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))

            szArguments = fmt.Sprintf("[0,10,\"%s\"]", value[1])
            cb.ArrNswh = HttpG.GetCompanyNswh(HttpG.PostGzHttpJson(s, "TQyQynswhBS", szArguments, "findTQyQynswhInfo"))

            // only can be set at the end of process 
            szCompanyChenxinMap[key] = cb

            szArguments = fmt.Sprintf("[0,10,\"%s\"]", value[1])
            arrId := GetCompanyJczl3(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findTQyXmyjInfo"), value[1], true)  // _______________ debug
            //fmt.Println(arrId)
            //arrId := GetCompanyJczl3(HttpG.PostGzHttpJson(s, "JyBlzbtzsBS", szArguments, "findJyBlzbtzsInfox"), value[1], true)

            for _, szId := range arrId {
                //szId = "yj_80f11ebf-be47-4d80-9ea8-e30c3d20831c" // ____________ debug
                fmt.Println("Wait 10 Second...")
                time.Sleep(5 * time.Second)
                fmt.Println("已载入 "+szId)
                szArguments := fmt.Sprintf("[{\"xmyjid\":\"%s\"}]", szId)
                qBase := GetCompanyJczl_XmyjBase(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyyj"))
                fmt.Println(qBase)
                //ShowReader(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyyj"))

                szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szId)
                //ShowReader(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyzz"))
                arrQyzz := GetCompanyJczl_XmyjQyzz(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyzz"))
                //fmt.Println(arrQyzz)

                szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szId)
                qXmgm := GetCompanyJczl_XmyjPlus(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findXmgm"))
                //fmt.Println(qXmgm)

                szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szId)
                qHjqk := GetCompanyJczl_XmyjHjqk(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findHjqk"))
                //fmt.Println(qHjqk)

                szXmyjMap[szId] = HttpG.Xmyj{
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
    HttpG.CheckError(err)
    defer file.Close()

    szTitleLine := "企业名称\t注册资本(万元) \t2010纳税\t2011纳税\t2012纳税\t企业类型\t有效期\r\n"
    file.WriteString(szTitleLine)

    for key, value := range szCompanyChenxinMap {
        //fmt.Println(value)
        s := []string{key}
        s = append(s, value.SzZczb)

        szNs2010 := ""
        szNs2011 := ""
        szNs2012 := ""
        for _, szNs := range value.ArrNswh {
            if szNs.SzYear == "2010" {
                szNs2010 = szNs.SzMoney
            } else if szNs.SzYear == "2011" {
                szNs2011 = szNs.SzMoney
            } else if szNs.SzYear == "2012" {
                szNs2012 = szNs.SzMoney
            }
        }
        s = append(s, szNs2010)
        s = append(s, szNs2011)
        s = append(s, szNs2012)

        if len(value.ArrQylx) > 0 {
            for _, value2 := range value.ArrQylx {
                s1 := append(s, value2.SzName)
                s1 = append(s1, value2.SzEndTime)

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
    HttpG.CheckError(err)
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
                s1 = append(s1, HttpG.GetZzdj(value2.Zzdj))
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
    HttpG.CheckError(err)
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

func PostHttpResp1(szUrl string, nSelTypeId int, szQymc string) (*http.Response) {
	enc:=mahonia.NewEncoder("gbk")
	//converts a  string from UTF-8 to gbk encoding.
	szGbk := enc.ConvertString(szQymc) 
    //cd, err := iconv.Open("gbk", "utf-8")
    //HttpG.CheckError(err)
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
    return HttpG.PostHttpResp(szUrl, szPost);
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

func GetCompanyJczl_XmyjHjqk(resp* http.Response) ([]HttpG.XmyjHjqk) {
    r := resp.Body
    defer r.Close()

    type XmyjHjqkJson struct {
        Data []HttpG.XmyjHjqk
    }

    dec := json.NewDecoder(r)
    var d XmyjHjqkJson
    err := dec.Decode(&d)
    HttpG.CheckError(err)

    return d.Data
}

func GetCompanyJczl_XmyjPlus(resp* http.Response) (map[string]string) {
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
    HttpG.CheckError(err)

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

func GetCompanyJczl_XmyjQyzz(resp* http.Response) ([]HttpG.XmyjQyzzJson){
    r := resp.Body
    defer r.Close()

    type QyyjQyzzJson struct {
        Data []HttpG.XmyjQyzzJson
    }

    dec := json.NewDecoder(r)
    var d QyyjQyzzJson
    err := dec.Decode(&d)
    HttpG.CheckError(err)

    return d.Data
}

func GetCompanyJczl_XmyjBase(resp* http.Response) (HttpG.XmyjBaseJson){
    r := resp.Body
    defer r.Close()

    dec := json.NewDecoder(r)
    var d HttpG.XmyjBaseJson
    err := dec.Decode(&d)
    HttpG.CheckError(err)

    return d
}

func GetCompanyJczl3(resp* http.Response, szId string, bNeedPage bool) ([]string) {
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
    HttpG.CheckError(err)

    var arrId []string

    for _, a := range d.Data{
        //fmt.Println(a.Xmyjid)
        arrId = append(arrId, a.Xmyjid)
    }

    if bNeedPage == true {
        nTotal, err := strconv.Atoi(d.Total)
        HttpG.CheckError(err)

        nPage := nTotal / 10
        if nPage > 1 {
            for i := 1; i <= nPage; i ++ {
                fmt.Println("Wait 10 Second...")
                time.Sleep(10 * time.Second)
                s := "http://www.gzzb.gd.cn/qyww/json";
                szArguments := fmt.Sprintf("[%d,10,\"%s\"]", i * 10, szId)
                arrId2 := GetCompanyJczl3(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findTQyXmyjInfo"), szId, false)
                for _, v := range arrId2 {
                    arrId = append(arrId, v)
                }

                //break //______ debug
            }
        }
    }
    return arrId
}


func FilterBody2(resp *http.Response, SzCompanyName string) {
    r := resp.Body
    defer r.Close()
    doc, err := html.Parse(r)
    HttpG.CheckError(err)

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
                        szQybhUrlMap[SzCompanyName] = append(szQybhUrlMap[SzCompanyName], a.Val)
                        //fmt.Println(szQybhUrlMap)
                        arr := strings.Split(a.Val, "=")
                        szQybhUrlMap[SzCompanyName] = append(szQybhUrlMap[SzCompanyName], arr[1])
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
    HttpG.CheckError(err)

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
        //HttpG.CheckError(err)
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


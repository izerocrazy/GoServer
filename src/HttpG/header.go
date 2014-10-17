package HttpG

import (
	"Base"
	"code.google.com/p/go.net/html"
	"encoding/json"
	"fmt"
	"github.com/p/mahonia"
	"net/http"
	"net/url"
	"os"
	"strings"
	"strconv"
	"time"
	"io/ioutil"
)

// because do not make ,waste my time
var c chan int = make(chan int)

func ShowReader(resp *http.Response) {
	r := resp.Body
	//defer resp.Body.Close()

	var buf [1024]byte
	reader := r
	//fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		fmt.Println(string(buf[0:n]))
		if err != nil {
			break
		}

		//cd, err := iconv.Open("gbk", "utf-8")
		//HttpG.Base.CheckErr(err)
		//defer cd.Close()
		//szGbk := cd.ConvString(string(buf[0:n]))
		//fmt.Print(szGbk)
	}

	//os.Exit(0)
}

func GetChannel() int {
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

type ProjectBaseInfo struct {
	Zbtzsrq string
	Zbj     string
	Xmjlxm  string
	Jgysrq  string
	Htj     string
}

type CompanyBaseInfo struct {
	SzCompanyName string
	ArrQylx       []CompanyQylx
	SzZczb        string
	ArrNswh       []CompanyNswh
	ArrQyzz       []CompanyQyzz
	ArrQyzzInfo   [][]CompanyQyzzInfo
}

type CompanyQylx struct {
	SzName    string
	SzEndTime string
}

type CompanyNswh struct {
	SzYear  string
	SzMoney string
}

type CompanyQyzz struct {
	Qyzzid string
}

type CompanyQyzzInfo struct {
	Zzdj     string
	ZznrName string
}

type ProjectZz struct {
	Zzmc string
	Zzdj string
}

type ProjectSize struct {
	Gclb string
	Gmzb string
	Sl   string
	Dw   string
}

type ProjectPrice struct {
	Nd   string
	Hjmc string
	Bjsj string
	Bjdw string
}

type Xmyj struct {
	Base    ProjectBaseInfo
	ArrQyzz []ProjectZz
	ArrXmgm []ProjectSize
	ArrHjqk []ProjectPrice
}

type QyyjSample struct {
	Name string
	Url  string
}

var channel chan *http.Response

func GetHttpResp(szUrl string) *http.Response {
	channel = make(chan *http.Response)
	var response *http.Response = nil
	for response == nil {
		go GetHttpResp2(szUrl)
		response = <-channel
	}
	return response
}

func GetHttpResp2(szUrl string) { //*http.Response {
	client := &http.Client{}

	request, err := http.NewRequest("GET", szUrl, nil)
	// only accept UTF-8
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	Base.CheckErr(err)

	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		response = nil
		fmt.Println("Get url err wait 10 Second....")
		time.Sleep(10 * time.Second)
		channel <- response
		return
	}
	// for {
	// 	response, err = client.Do(request)
	// 	if err != nil {
	// 		fmt.Println("Get url err wait 10 Second....")
	// 		time.Sleep(10 * time.Second)
	// 	} else {
	// 		break
	// 	}
	// }

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

	// return response
	channel <- response
}

func PostHttpResp(szUrl string, szPost *strings.Reader) *http.Response {
	channel = make(chan *http.Response)
	var response *http.Response = nil
	for response == nil {
		go PostHttpResp2(szUrl, szPost)
		response = <-channel
	}
	return response
}

func PostHttpResp2(szUrl string, szPost *strings.Reader) { //*http.Response {
	//client := &http.Client{}
	request, err := http.NewRequest("POST", szUrl, szPost)
	Base.CheckErr(err)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,nl;q=0.2,zh-TW;q=0.2")
	request.Header.Add("Host", "www.gzzb.gd.cn")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.149 Safari/537.36")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		resp = nil
		fmt.Println("post url err wait 10 second")
		time.Sleep(10 * time.Second)
		channel <- resp
		return
	}
	// for {
	// 	resp, err = http.DefaultClient.Do(request)
	// 	if err != nil {
	// 		fmt.Println("post url err wait 10 second")
	// 		time.Sleep(10 * time.Second)
	// 	} else {
	// 		break
	// 	}
	// }
	//Base.CheckErr(err)

	chSet := GetCharset(resp)
	// fmt.Printf("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	// return resp
	channel <- resp
}

func PostGzHttpJson(szUrl string, szService string, szArguments string, szFunc string) *http.Response {
	values := make(url.Values)

	values.Set("service", szService)
	values.Set("arguments", szArguments)
	values.Set("method", szFunc)

	szPost := strings.NewReader(values.Encode())
	// fmt.Println(szUrl, szPost)
	return PostHttpResp(szUrl, szPost)
}

func FindDivNodeByName(node *html.Node, szName string) []*html.Node {
	var retList []*html.Node

	if (node.Type == html.DocumentNode || node.Type == html.ElementNode) && node.Data == "div" {
		for _, a := range node.Attr {
			if a.Key == "class" && a.Val == szName {
				retList = append(retList, node)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		rList := FindDivNodeByName(c, szName)

		for _, n := range rList {
			retList = append(retList, n)
		}
	}

	return retList
}

func FindNodeByTypeName(node *html.Node, szTypeName string) []*html.Node {
	//fmt.Println("find node by type name", szTypeName, "this node's data is", node.Data)
	var retList []*html.Node

	if (node.Type == html.DocumentNode || node.Type == html.ElementNode) && node.Data == szTypeName {
		retList = append(retList, node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		rList := FindNodeByTypeName(c, szTypeName)

		for _, n := range rList {
			retList = append(retList, n)
		}
	}

	return retList
}

func GetNodeText(node *html.Node) (text string) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			return html.UnescapeString(strings.Replace(child.Data, "\u00a0", "",-1))
		}
	}

	return ""
}

func GetCompanyQyyjInfos(resp *http.Response, nCompanyId int) []QyyjSample {
	var retList []QyyjSample

	r := resp.Body
	defer r.Close()
	szAll, err := ioutil.ReadAll(r)
	
	doc, err := html.Parse(strings.NewReader(strings.Replace(string(szAll),  "&nbsp;", "",-1)))
	Base.CheckErr(err)

	divList := FindDivNodeByName(doc, "bszn_right_table")
	for _, div := range divList {
		trList := FindNodeByTypeName(div, "tr")
		for n, tr := range trList {
			if n == 0 {
				continue
			}

			tdList := FindNodeByTypeName(tr, "td")
			m := len(tdList)
			if m == 1 {	// 是为空
				continue
			} else{
				if m < 3 {	// Parse 可能为空（由于字符异常引起）
					html.Render(os.Stdout, div)
					Base.PrintErrExit("end")
				}
			}
			
			// 验证是否是对应的 Id 号
			tbChildId := GetNodeText(tdList[3])
			// fmt.Println(strconv.Atoi(tbChildId))
			nChildId, _ := strconv.Atoi(tbChildId)
			if nChildId != nCompanyId {
				continue
			}

			var qyyjSample QyyjSample
			//tbChild := td.FirstChild   //Table Index
			//fmt.Println(tbChild.Data)
			//tbChild = tbChild.NextSibling  //Id: YJ201103170297
			//fmt.Println(tbChild.Data)
			tbChild := tdList[2] //a and name

			for tempChild := tbChild.FirstChild; tempChild != nil; tempChild = tempChild.NextSibling {
				if tempChild.Data == "a" {
					szText := GetNodeText(tempChild)
					enc := mahonia.NewDecoder("gbk")
					szGbk := enc.ConvertString(szText)
					qyyjSample.Name = szGbk

					for _, a := range tempChild.Attr {
						if a.Key == "href" {
							qyyjSample.Url = a.Val
						}
					}
				}
			}

			retList = append(retList, qyyjSample)
		}
	}

	//fmt.Println(retList)
	return retList
}

func GetCompanyQylxInfo(resp *http.Response) CompanyBaseInfo {
	var cb CompanyBaseInfo

	r := resp.Body
	defer r.Close()
	doc, err := html.Parse(r)
	Base.CheckErr(err)

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

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				if bFindDiv1 == true {
					enc := mahonia.NewDecoder("gbk")
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

func GetCompanyJczl(resp *http.Response) (string, string) {
	r := resp.Body
	defer r.Close()

	type Cjson struct {
		Czzb string
		Qymc string
	}
	dec := json.NewDecoder(r)
	var c Cjson
	err := dec.Decode(&c)
	Base.CheckErr(err)

	return c.Qymc, c.Czzb
}

func GetCompanyNswh(resp *http.Response) []CompanyNswh {
	r := resp.Body
	defer r.Close()
	type QylxData struct {
		Nsze string
		Nd   string
	}

	type Djson struct {
		Data []QylxData
	}

	dec := json.NewDecoder(r)
	var d Djson
	err := dec.Decode(&d)
	Base.CheckErr(err)

	var arrCn []CompanyNswh
	for _, a := range d.Data {
		if a.Nd == "2010" || a.Nd == "2011" || a.Nd == "2012" || a.Nd == "2013" {
			//fmt.Println(a.Nd, a.Nsze)
			var cn CompanyNswh
			cn.SzYear = a.Nd
			cn.SzMoney = a.Nsze

			arrCn = append(arrCn, cn)
		}
	}

	return arrCn
}

func GetCompanyQyzz(resp *http.Response) []CompanyQyzz {
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
	Base.CheckErr(err)

	var arrCn []CompanyQyzz
	for _, a := range d.Data {
		var cn CompanyQyzz
		cn.Qyzzid = a.Qyzzid

		arrCn = append(arrCn, cn)
	}

	return arrCn
}

func GetCompanyQyzzInfo(resp *http.Response) []CompanyQyzzInfo {
	r := resp.Body
	defer r.Close()

	type Djson struct {
		Data []CompanyQyzzInfo
	}

	dec := json.NewDecoder(r)
	var d Djson
	err := dec.Decode(&d)
	Base.CheckErr(err)

	return d.Data
}

func GetProjectBaseInfo(resp *http.Response) ProjectBaseInfo {
	//ShowReader(resp)
	r := resp.Body
	defer r.Close()

	dec := json.NewDecoder(r)
	var t ProjectBaseInfo
	err := dec.Decode(&t)
	Base.CheckErr(err)

	// fmt.Println(t)

	return t
}

func GetProjectQyzz(resp *http.Response) []ProjectZz {
	//ShowReader(resp)
	r := resp.Body
	defer r.Close()

	type QyyjQyzzJson struct {
		Data []ProjectZz
	}

	dec := json.NewDecoder(r)
	var d QyyjQyzzJson
	err := dec.Decode(&d)
	Base.CheckErr(err)

	// fmt.Println(d.Data)

	return d.Data
}

func GetProjectSize(resp *http.Response) []ProjectSize {
	r := resp.Body
	defer r.Close()

	type XmyjHjqkJson struct {
		Data []ProjectSize
	}

	dec := json.NewDecoder(r)
	var d XmyjHjqkJson
	err := dec.Decode(&d)
	Base.CheckErr(err)

	return d.Data
}

func GetProjectPrice(resp *http.Response) []ProjectPrice {
	r := resp.Body
	defer r.Close()

	type XmyjHjqkJson struct {
		Data []ProjectPrice
	}

	dec := json.NewDecoder(r)
	var d XmyjHjqkJson
	err := dec.Decode(&d)
	Base.CheckErr(err)

	return d.Data
}

func CreateFileWithNameAddTitle(szFileName string, szTitleLine string) (file *os.File) {
	file = Base.CreateOrAppendFile(szFileName)

	file.WriteString(szTitleLine)

	return file
}

func GetZzdj(n string) string {
	//StrMap := map[string]string{"01":"特级","02":"一级","03":"二级","04":"三级","05":"不分等级"}
	StrMap := map[string]string{"01": "特级", "02": "一级", "03": "二级", "04": "三级", "06": "甲", "07": "乙", "08": "丙", "09": "暂乙级", "12": "暂定级", "13": "暂二级", "14": "暂三级", "17": "暂一级", "21": "丁", "10": "暂五级", "05": "临时资质", "20": "五级", "19": "四级", "11": "暂四级", "15": "不分等级", "16": "暂甲级", "18": "暂丙级", "22": "预备级"}

	return StrMap[n]
}

func GetHjmc(n string) string {
	StrMap := map[string]string{"01": "中国建设工程鲁班奖（国家优质工程）", "02": "全国市政金杯示范工程", "03": "国家优质工程（金质奖）", "04": "国家优质工程（银质奖）", "05": "广东省建设工程金匠奖", "06": "全国建筑工程装饰奖", "07": "广州地区建设工程质量“五羊杯”", "08": "广州市优良样板工程", "09": "广州市安全文明施工样板工地（市双优）", "10": "广东省房屋市政工程安全生产文明施工示范工地（原广东省安全文明施工样板工地）", "11": "广东省建设工程优质奖（原省优良样板工程）", "12": "广州市市政优良样板工程", "13": "中国土木工程詹天佑奖", "14": "全国建筑业新技术应用示范工程执行单位", "15": "广东省建筑业新技术应用示范工程执行单位", "16": "广东省优秀建筑装饰工程奖", "17": "广州市建筑装饰优质工程奖", "18": "广州市建设工程（市政）质量“五羊杯”"}

	return StrMap[n]
}

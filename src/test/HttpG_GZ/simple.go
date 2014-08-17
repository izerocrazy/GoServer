package main

import (
	"HttpG"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	DoForOneCompany(10020)
	// DoForOneQyyj("x")
}

func DoForOneCompany(nCompanyId int) {
	var sampleList []HttpG.QyyjSample
	nOffset := 0
	for {
		szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn/cms/wz/view/sccx/QyyjServlet?qyyj_qybh=%d&qyyj_qymc=&qyyj_xmbh=&qyyj_xmmc=&siteId=1&channelId=29&pager.offset=%d", nCompanyId, nOffset)
		fmt.Println(szHtmlUrl)
		//s := "http://www.gzzb.gd.cn/qyww/json";
		//szArguments := fmt.Sprintf("[\"%d\"]", nCompanyId)

		bIsEnd := true
		retList := HttpG.GetCompanyQyyjInfos(HttpG.GetHttpResp(szHtmlUrl))
		for n, sample := range retList {
			sampleList = append(sampleList, sample)

			if n == 14 {
				bIsEnd = false
			}
		}

		if bIsEnd == true {
			break
		} else {
			nOffset = nOffset + 15
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println(sampleList)

	for _, a := range sampleList {
		DoForOneQyyj(a.Url)
	}
	//cb := HttpG.GetCompanyQylxInfo(HttpG.GetHttpResp(szHtmlUrl));
	//cb.SzCompanyName, cb.SzZczb = HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))
}

func DoForOneQyyj(szUrl string) {
	szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn%s", szUrl)
	fmt.Println(szHtmlUrl)

	strList := strings.Split(szUrl, "=")
	szCompanyId := strList[1]
	// szCompanyId := "402828ac2f522638012f76cb6c841030"

	var xmInfo HttpG.Xmyj
	s := "http://www.gzzb.gd.cn/qyww/json"
	szArguments := fmt.Sprintf("[{\"xmyjid\":\"%s\"}]", szCompanyId)
	xmInfo.Base = HttpG.GetProjectBaseInfo(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findQyyj"))
	// HttpG.ShowReader(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findQyyj"))

	szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szCompanyId)
	xmInfo.ArrQyzz = HttpG.GetProjectQyzz(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findQyzz"))
	// HttpG.ShowReader(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findQyzz"))

	szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szCompanyId)
	xmInfo.ArrHjqk = HttpG.GetProjectPrice(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findHjqk"))
	// HttpG.ShowReader(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findHjqk"))

	fmt.Println(xmInfo)
	SaveToFile(xmInfo)
}

func SaveToFile(xmInfo HttpG.Xmyj) {
	f, err := os.Create("XMYJ.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	// w.Write([]string{"1", "张三", "23"})
	// Base
	var data []string
	data = append(data, xmInfo.Base.Zbtzsrq)
	data = append(data, xmInfo.Base.Zbj)
	data = append(data, xmInfo.Base.Xmjlxm)
	data = append(data, xmInfo.Base.Jgysrq)
	data = append(data, xmInfo.Base.Htj)

	// zz
	var szQyzz string
	for _, a := range xmInfo.ArrQyzz {
		szQyzz = szQyzz + a.Zzmc + "\t" + a.Zzdj + "\r\n"
	}
	data = append(data, szQyzz)

	// hjqk
	var szHjqk string
	for _, a := range xmInfo.ArrHjqk {
		szHjqk = szHjqk + a.Nd + "\t" + a.Hjmc + "\t" + a.Bjsj + "\t" + a.Bjdw + "\r\n"
	}
	data = append(data, szHjqk)

	w.Write(data)
	w.Flush()
}

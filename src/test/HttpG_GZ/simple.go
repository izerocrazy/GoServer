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
	f, err := os.Create("XMYJ.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"企业编号", "业绩名称", "中标日期", "中标价", "项目经理", "合同价", "竣工验收时间", "资质名称和等级", "特征", "获奖"})
	w.Flush()

	DoForOneCompany(10020, f)
	// DoForOneQyyj("x")
}

func DoForOneCompany(nCompanyId int, file *os.File) {
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
		xmInfo := DoForOneQyyj(a.Url)
		SaveToFile(nCompanyId, a.Name, xmInfo, file)
	}
	//cb := HttpG.GetCompanyQylxInfo(HttpG.GetHttpResp(szHtmlUrl));
	//cb.SzCompanyName, cb.SzZczb = HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))
}

func DoForOneQyyj(szUrl string) HttpG.Xmyj {
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
	xmInfo.ArrXmgm = HttpG.GetProjectSize(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findXmgm"))

	szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szCompanyId)
	xmInfo.ArrHjqk = HttpG.GetProjectPrice(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findHjqk"))
	// HttpG.ShowReader(HttpG.PostGzHttpJson(s, "XmyjBS", szArguments, "findHjqk"))

	fmt.Println(xmInfo)

	return xmInfo
}

func SaveToFile(nCompanyId int, szName string, xmInfo HttpG.Xmyj, file *os.File) {
	w := csv.NewWriter(file)
	// w.Write([]string{"1", "张三", "23"})
	// Base
	var data []string
	szCompanyId := fmt.Sprintf("%d", nCompanyId)
	data = append(data, szCompanyId)
	data = append(data, szName)

	data = append(data, xmInfo.Base.Zbtzsrq)
	data = append(data, xmInfo.Base.Zbj)
	data = append(data, xmInfo.Base.Xmjlxm)
	data = append(data, xmInfo.Base.Htj)
	data = append(data, xmInfo.Base.Jgysrq)

	// zz
	var szQyzz string
	for _, a := range xmInfo.ArrQyzz {
		szQyzz = szQyzz + a.Zzmc + " " + a.Zzdj + "\r\n"
	}
	data = append(data, szQyzz)

	// xmgm
	var szXmgm string
	for _, a := range xmInfo.ArrXmgm {
		szXmgm = szXmgm + a.Gclb + " " + a.Gmzb + " " + a.Sl + " " + a.Dw + "\r\n"
	}
	data = append(data, szXmgm)

	// hjqk
	var szHjqk string
	for _, a := range xmInfo.ArrHjqk {
		szHjqk = szHjqk + a.Nd + " " + a.Hjmc + " " + a.Bjsj + " " + a.Bjdw + "\r\n"
	}
	data = append(data, szHjqk)

	w.Write(data)
	w.Flush()
}

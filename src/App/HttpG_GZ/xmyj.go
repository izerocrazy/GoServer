package main

import (
	"Base"
	"HttpG"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var nBeginId int
	var nEndId int

	f := Base.CreateOrAppendFile("XMYJ.xls")
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"企业编号", "企业名称", "业绩名称", "中标日期", "中标价", "项目经理", "合同价", "竣工验收时间", "资质名称和等级", "特征", "获奖"})
	w.Flush()

	fmt.Print("温馨提示>> : 如果你希望进行新一轮的信息选取，请在输入前删除上次的信息文件。\r\n")
	fmt.Print("请输入开始企业ID（建议：网站默认第一个企业 ID 为10002）：")
	fmt.Scanf("%d", &nBeginId)
	var szStr string
	fmt.Scanf("%s", &szStr)
	fmt.Print("请输入结束企业ID（建议：目前最后一个企业 ID 至少大于20000）：")
	fmt.Scanf("%d", &nEndId)

	for i := nBeginId; i < nEndId+1; i++ {
		fmt.Println("正在载入 ID：", i)
		DoForOneCompany(i, f)

		if i < nEndId {
			time.Sleep(2 * time.Second)
		}
	}
}

func DoForOneCompany(nCompanyId int, file *os.File) {
	fmt.Println(">>>>>>>>>>> 读取公司 ID：", nCompanyId)

	s := "http://www.gzzb.gd.cn/qyww/json"
	szArguments := fmt.Sprintf("[\"%d\"]", nCompanyId)
	SzCompanyName, _ := HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))

	var sampleList []HttpG.QyyjSample
	nOffset := 0
	for {
		szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn/cms/wz/view/sccx/QyyjServlet?qyyj_qybh=%d&qyyj_qymc=&qyyj_xmbh=&qyyj_xmmc=&siteId=1&channelId=29&pager.offset=%d", nCompanyId, nOffset)

		bIsEnd := true
		retList := HttpG.GetCompanyQyyjInfos(HttpG.GetHttpResp(szHtmlUrl), nCompanyId)
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

		// time.Sleep(2 * time.Second)
	}

    //fmt.Println(sampleList);
	for i, a := range sampleList {
		if a.Name == "" {
            continue;
        }
		fmt.Println(">>>>>>>>>>> 读取公司 ID：", nCompanyId, "，其项目：", a.Name)
		xmInfo := DoForOneQyyj(a.Url)
		SaveToFile(nCompanyId, SzCompanyName, a.Name, xmInfo, file)
		if i%30 == 0 {
			fmt.Println("暂停 1 s")
			time.Sleep(1 * time.Second)
		}
	}
}

func DoForOneQyyj(szUrl string) HttpG.Xmyj {
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

	// fmt.Println(xmInfo)

	return xmInfo
}

func SaveToFile(nCompanyId int, szCompanyName string, szName string, xmInfo HttpG.Xmyj, file *os.File) {
	w := csv.NewWriter(file)
	// w.Write([]string{"1", "张三", "23"})
	// Base
	var data []string
	szCompanyId := fmt.Sprintf("%d", nCompanyId)
	data = append(data, szCompanyId)
	data = append(data, szCompanyName)
	data = append(data, szName)

	data = append(data, xmInfo.Base.Zbtzsrq)
	data = append(data, xmInfo.Base.Zbj)
	data = append(data, xmInfo.Base.Xmjlxm)
	data = append(data, xmInfo.Base.Htj)
	data = append(data, xmInfo.Base.Jgysrq)

	// zz
	var szQyzz string
	for _, a := range xmInfo.ArrQyzz {
		szQyzz = szQyzz + a.Zzmc + " " + HttpG.GetZzdj(a.Zzdj) + "\r\n"
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
		szHjqk = szHjqk + a.Nd + " " + HttpG.GetHjmc(a.Hjmc) + " " + a.Bjsj + " " + a.Bjdw + "\r\n"
	}
	data = append(data, szHjqk)

	w.Write(data)
	w.Flush()
}

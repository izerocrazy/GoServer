package main

import (
	"Base"
	"HttpG"
	"encoding/csv"
	"fmt"
	"os"
	// "strings"
	"time"
)

func main() {
	var nBeginId int
	var nEndId int

	fmt.Print("温馨提示>> : 如果你希望进行新一轮的信息选取，请在输入前删除上次的信息文件（文件『1.txt』和『2.txt』)。\r\n")
	fmt.Print("请输入开始企业ID（建议：网站默认第一个企业 ID 为10002）：")
	fmt.Scanf("%d", &nBeginId)
	var szStr string
	fmt.Scanf("%s", &szStr)
	fmt.Print("请输入结束企业ID（建议：目前最后一个企业 ID 至少大于20000）：")
	fmt.Scanf("%d", &nEndId)

	// szTitleLine :=
	// file1 := HttpG.CreateFileWithNameAddTitle("1.txt", szTitleLine)
	// defer file1.Close()

	file1 := Base.CreateOrAppendFile("1.xls")
	defer file1.Close()

	file1.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(file1)
	w.Write([]string{"企业名称", "注册资本(万元) ", "2010纳税", "2011纳税", "2012纳税", "2013纳税", "企业类型", "有效期"})
	w.Flush()

	file2 := Base.CreateOrAppendFile("2.xls")
	defer file2.Close()

	file2.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w2 := csv.NewWriter(file2)
	w2.Write([]string{"企业名称", "资质等级 ", "资质内容"})
	w2.Flush()

	for i := nBeginId; i < nEndId+1; i++ {
		fmt.Println("正在载入 ID：", i)
		DoForOneCompany(i, file1, file2)

		if i < nEndId {
			time.Sleep(2 * time.Second)
		}
	}
}

func DoForOneCompany(nCompanyId int, file *os.File, file2 *os.File) {
	szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn/qyww/sccx/basicInfoview.jsp?qybh=%d", nCompanyId)
	s := "http://www.gzzb.gd.cn/qyww/json"
	szArguments := fmt.Sprintf("[\"%d\"]", nCompanyId)

	cb := HttpG.GetCompanyQylxInfo(HttpG.GetHttpResp(szHtmlUrl))
	cb.SzCompanyName, cb.SzZczb = HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))

	szArguments = fmt.Sprintf("[0,10,\"%d\"]", nCompanyId)
	cb.ArrNswh = HttpG.GetCompanyNswh(HttpG.PostGzHttpJson(s, "TQyQynswhBS", szArguments, "findTQyQynswhInfo"))

	szArguments = fmt.Sprintf("[0,100,\"%d\"]", nCompanyId)
	cb.ArrQyzz = HttpG.GetCompanyQyzz(HttpG.PostGzHttpJson(s, "TQyQyzzBS", szArguments, "findTQyQyzzInfo"))

	for _, a := range cb.ArrQyzz {
		szArguments = fmt.Sprintf("[0,100,\"%s\"]", a.Qyzzid)
		cb.ArrQyzzInfo = append(cb.ArrQyzzInfo, HttpG.GetCompanyQyzzInfo(HttpG.PostGzHttpJson(s, "TQyZzxxBS", szArguments, "findZzxxInfo")))
	}

	if cb.SzCompanyName != "" {
		SaveToFile(nCompanyId, cb, file, file2)
		fmt.Println("已读取完公司", cb.SzCompanyName, "，请稍候。\r\n")
	} else {
		fmt.Println("此 ID 暂无对应公司信息。\r\n")
	}

	//HttpG.SendChannel(2)
}

func SaveToFile(nCompanyId int, cb HttpG.CompanyBaseInfo, file *os.File, file2 *os.File) {
	w := csv.NewWriter(file)
	s := []string{}

	szCompanyId := fmt.Sprintf("%d", nCompanyId)
	s = append(s, szCompanyId)
	s = append(s, cb.SzCompanyName)
	szNs2010 := ""
	szNs2011 := ""
	szNs2012 := ""
	szNs2013 := ""
	for _, szNs := range cb.ArrNswh {
		if szNs.SzYear == "2010" {
			szNs2010 = szNs.SzMoney
		} else if szNs.SzYear == "2011" {
			szNs2011 = szNs.SzMoney
		} else if szNs.SzYear == "2012" {
			szNs2012 = szNs.SzMoney
		} else if szNs.SzYear == "2013" {
			szNs2013 = szNs.SzMoney
		}
	}
	s = append(s, szNs2010)
	s = append(s, szNs2011)
	s = append(s, szNs2012)
	s = append(s, szNs2013)

	if len(cb.ArrQylx) > 0 {
		var NameList string
		var EndTimeList string
		for _, cb2 := range cb.ArrQylx {
			NameList = NameList + cb2.SzName + "\r\n"
			EndTimeList = EndTimeList + cb2.SzEndTime + "\r\n"
		}
		s = append(s, NameList)
		s = append(s, EndTimeList)
	}

	w.Write(s)
	w.Flush()

	//////////////////////////////
	w2 := csv.NewWriter(file2)
	for _, arr := range cb.ArrQyzzInfo {
		for _, cq := range arr {
			s2 := []string{}
			s2 = append(s2, cb.SzCompanyName)
			s2 = append(s2, HttpG.GetZzdj(cq.Zzdj))
			s2 = append(s2, cq.ZznrName)

			w2.Write(s2)
			w2.Flush()
		}
	}
}

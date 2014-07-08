package main

import (
        "fmt"
        "os"
        "time"
        "strings"
        "HttpG"
)


func main() {
    //DoForOneCompany(10020);
    DoForOneQyyj("x");
}

func DoForOneCompany(nCompanyId int) {
    var sampleList []HttpG.QyyjSample
    nOffset := 0;
    for {
        szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn/cms/wz/view/sccx/QyyjServlet?qyyj_qybh=%d&qyyj_qymc=&qyyj_xmbh=&qyyj_xmmc=&siteId=1&channelId=29&pager.offset=%d", nCompanyId, nOffset);
        fmt.Println(szHtmlUrl);
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

    fmt.Println(sampleList);

    DoForOneQyyj(sampleList[0].Url)
    //cb := HttpG.GetCompanyQylxInfo(HttpG.GetHttpResp(szHtmlUrl));
    //cb.SzCompanyName, cb.SzZczb = HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))
}

func DoForOneQyyj(szUrl string) {
    szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn%s", szUrl)
    fmt.Println(szHtmlUrl)

    //strList := strings.Split(szUrl, "=")
    //szCompanyId := strList[1]
    szCompanyId := "yj_0bee9f63-007e-4321-b06f-f106dfbca77f" 

    s := "http://www.gzzb.gd.cn/qyww/json";
    szArguments := fmt.Sprintf("[{\"xmyjid\":\"%s\"}]", szCompanyId)
    //HttpG.GetProjectBaseInfo(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyyj"))

    szArguments = fmt.Sprintf("[0,500,{\"xmyjid\":\"%s\"}]", szCompanyId)
    //HttpG.GetProjectQyzz(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyzz"))
    //HttpG.ShowReader(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyzz"))
    HttpG.ShowReader(HttpG.PostGzHttpJson(s, "TQyXmyjBS", szArguments, "findQyzz"))
}

func SaveToFile(nCompanyId int, cb HttpG.CompanyBaseInfo, file *os.File, file2 *os.File) {
    s := []string{}

    szCompanyId := fmt.Sprintf("%d", nCompanyId)
    s = append(s, szCompanyId)
    s = append(s, cb.SzCompanyName)
    szNs2010 := ""
    szNs2011 := ""
    szNs2012 := ""
    for _, szNs := range cb.ArrNswh {
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

    if len(cb.ArrQylx) > 0 {
        for _, cb2 := range cb.ArrQylx {
            s1 := append(s, cb2.SzName)
            s1 = append(s1, cb2.SzEndTime)

            szLine := strings.Join(s1, "\t")
            szLine = szLine + "\r\n"
            file.WriteString(szLine)
            //fmt.Println(szLine)
        }
    } else {
        szLine := strings.Join(s, "\t")
        szLine = szLine + "\r\n"
        file.WriteString(szLine)
        //fmt.Println(szLine)
    }

    //////////////////////////////
    for _, arr := range cb.ArrQyzzInfo {
        for _, cq := range arr {
            s2 := []string{}
            s2 = append(s2, cb.SzCompanyName)
            s2 = append(s2, HttpG.GetZzdj(cq.Zzdj))
            s2 = append(s2, cq.ZznrName)
            szLine := strings.Join(s2, "\t")
            szLine = szLine + "\r\n"
            file2.WriteString(szLine)

            //fmt.Println(szLine)
        }
    }
}

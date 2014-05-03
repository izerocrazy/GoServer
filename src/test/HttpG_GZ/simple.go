package main

import (
        "fmt"
        "HttpG"
)

func main() {
    for i := 10002; i < 10003; i++ {
        DoForOneCompany(i)
    }
}

func DoForOneCompany(nCompanyId int) {
    szHtmlUrl := fmt.Sprintf("http://www.gzzb.gd.cn/qyww/sccx/basicInfoview.jsp?qybh=%d", nCompanyId);
    s := "http://www.gzzb.gd.cn/qyww/json";
    szArguments := fmt.Sprintf("[\"%d\"]", nCompanyId)

    cb := HttpG.GetCompanyQylxInfo(HttpG.GetHttpResp(szHtmlUrl));
    cb.SzCompanyName, cb.SzZczb = HttpG.GetCompanyJczl(HttpG.PostGzHttpJson(s, "TQyQyjczlBS", szArguments, "findQyjczl"))

    szArguments = fmt.Sprintf("[0,10,\"%d\"]", nCompanyId)
    cb.ArrNswh = HttpG.GetCompanyNswh(HttpG.PostGzHttpJson(s, "TQyQynswhBS", szArguments, "findTQyQynswhInfo"))

    szArguments = fmt.Sprintf("[0,100,\"%d\"]", nCompanyId)
    cb.ArrQyzz = HttpG.GetCompanyQyzz(HttpG.PostGzHttpJson(s, "TQyQyzzBS", szArguments, "findTQyQyzzInfo"))

    for _, a := range cb.ArrQyzz {
       szArguments = fmt.Sprintf("[0,100,\"%s\"]", a.Qyzzid)
       cb.ArrQyzzInfo = append(cb.ArrQyzzInfo, HttpG.GetCompanyQyzzInfo(HttpG.PostGzHttpJson(s, "TQyZzxxBS", szArguments, "findZzxxInfo")))
    }

    fmt.Println(cb)
}

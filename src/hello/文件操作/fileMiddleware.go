package main

import (
	"fmt"
	//"hello/server"
	"hello/array"
)

//const(
//	AUTHOR string = "Hsu Lay"
//	SINCE = "2020-06-01"
//	LINK = "http://www.haqh.com"
//	VERSION = "1.0.0"
//	PARAM = ""
//	DESC = "NAS.exe is the one of way to pass commands manipulate the file syst"
//)

//var infos map[string]string
//
//infos = make(map[string]string)
//infos["AUTHOR"] = "Hsu Lay"

type SYS struct{
	AUTHOR string
	SINCE string
	VERSION string
	LINK string
	PARAM string
	DESC string
}





func printLogo(){
	fmt.Print(`
———————————————————————————————————————————
	。。。。。。。。。。。。。。。。。。。。。。。。。
	。。。。。。。。。。。。。。。。。。。。。。。。。
	。。         __   _       ___   _____         。。
	。。        |  \ | |     /   | /  ___/        。。
	。。        |   \| |    / /| | | |___         。。
	。。        | |\   |   / / | | \___  \        。。
	。。        | | \  |  / /  | |  ___| |        。。
	。。        |_|  \_| /_/   |_| /_____/        。。
	。。                                          。。
	。。。。。。。。。。。。。。。。。。。。。。。。。
	。。。。。。。。。。。。。。。。。。。。。。。。。
`)

	//Infos := map[string]string{
	//	"AUTHOR" : "Hsu Lay",
	//	"SINCE" : "2020-06-01",
	//	"VERSION" : "1.0.0",
	//	"LINK" : "http://www.haqh.com",
	//	"PARAM" : "",
	//	"DESC" : "NAS.exe is the one of commands to manipulate the file system",
	//}

	Infos := new(SYS)
	Infos.AUTHOR = "Hsu Lay"
	Infos.SINCE = "2020-06-01"
	Infos.VERSION = "1.0.0"
	Infos.LINK = "http://www.haqh.com"
	Infos.PARAM = ""
	Infos.DESC = "NAS.exe is the one of commands to manipulate the file system"

	//type Infos struct {
	//	"AUTHOR" : "Hsu Lay",
	//	"SINCE" : "2020-06-01",
	//	"VERSION" : "1.0.0",
	//	"LINK" : "http://www.haqh.com",
	//	"PARAM" : "",
	//	"DESC" : "NAS.exe is the one of commands to manipulate the file system",
	//}

	//for k, v := range Infos{
	//	fmt.Printf("* %s: %s\n",k,v)
	//}

}

func main1(){
	array.Pratice()

	//printLogo()
	//server.Start()
}

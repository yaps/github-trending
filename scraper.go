package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	//loop
	for {
		scrapeObjC()
		scrapeGo()
		scrapeJS()
		createMarkDown()
		gitAddAll()
		gitCommit()
		gitPush()
		time.Sleep(time.Duration(24) * time.Hour)
	}
}

func scrapeObjC() {

}
func scrapeGo()       {}
func scrapeJS()       {}
func createMarkDown() {}
func gitAddAll() {
	app := "git"
	arg0 := "add"
	arg1 := "."
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitCommit(date string, language string) {
	app := "git"
	arg0 := "commit"
	arg1 := "-am"
	arg2 := fmt.Sprintf("[%s] %s", language, date)
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}
func gitPush() {
	app := "git"
	arg0 := "push"
	arg1 := "origin"
	arg2 := "master"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

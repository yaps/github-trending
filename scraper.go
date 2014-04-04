package main

import (
	_ "bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "os"
	"os/exec"
	_ "strconv"
	"time"
)

func main() {
	//loop
	for {
		dateString := dateString()
		fmt.Println(dateString)
		scrape("objective-c")
		scrape("go")
		scrape("javascript")
		createMarkDown(dateString)
		gitAddAll()
		gitCommit(dateString)
		gitPush()
		time.Sleep(time.Duration(24) * time.Hour)
	}
}

func dateString() string {
	y, m, d := time.Now().Date()
	return fmt.Sprint("%s-%i-%i", m.String(), d, y)
}

func scrape(language string) {
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument(fmt.Sprintf("https://github.com/trending?l=%s", language)); e != nil {
		panic(e.Error())
	}

	doc.Find("li.repo-leaderboard-list-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("div h2 a").Text()
		owner := s.Find("span.owner-name").Text()
		repoName := s.Find("strong").Text()
		fmt.Println(title, owner, repoName)
	})
}

func createMarkDown(date string) {}
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

func gitCommit(date string) {
	app := "git"
	arg0 := "commit"
	arg1 := "-am"
	arg2 := date
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

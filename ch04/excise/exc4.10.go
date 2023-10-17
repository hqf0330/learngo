// 4.10 修改issues实例，按照时间来输出结果，比如一个月以内，一年以内或者超过一年

package main

import (
	"log"
	"os"
	"time"
)
import "../book/github"

func main() {
	monthBefore := time.Now().AddDate(0, -1, 0)
	yearBefore := time.Now().AddDate(-1, 0, 0)

	var monthBuf []*github.Issue
	var yearBuf []*github.Issue
	var yearAfter []*github.Issue

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		if monthBefore.Before(item.CreatedAt) {
			monthBuf = append(monthBuf, item)
		} else if yearBefore.Before(item.CreatedAt) {
			yearBuf = append(yearBuf, item)
		} else if yearBefore.After(item.CreatedAt) {
			yearAfter = append(yearAfter, item)
		}
	}

}

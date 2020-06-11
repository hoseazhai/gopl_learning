package main

import (
	"ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%d issues:\n", result.Items[:len(result.Items)])

	nowTime := time.Now()
	getTime := nowTime.AddDate(0, -1, 0)
	fmt.Println(getTime)
	resTime := getTime.Format("2006-01-02 15:04:05")

	for _, item := range result.Items {
		mTime, err := time.Parse("2006-01-02 15:04:05", resTime)
		cTime := item.CreatedAt.Format("2006-01-02 15:04:05")
		c1Time, err := time.Parse("2006-01-02 15:04:05", cTime)
		if err == nil && mTime.After(c1Time) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}

	}
}

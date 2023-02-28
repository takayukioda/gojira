package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	jira := NewJira(os.Getenv("JIRA_DOMAIN"), os.Getenv("JIRA_EMAIL"), os.Getenv("JIRA_TOKEN"))
	boardId, err := strconv.ParseInt(os.Getenv("JIRA_BOARD_ID"), 10, 0)
	if err != nil {
		log.Fatalln("fatal", err)
	}

	starts, err := time.Parse(time.RFC3339, "2023-03-06T09:00:00+09:00")
	if err != nil {
		log.Fatalln("fatal:", err)
	}

	days := time.Hour * 24 * 14 // 2 weeks
	for ; starts.Month() < time.June; starts = starts.Add(days) {
		if err != nil {
			log.Fatalln("fatal:", err)
		}
		board, err := jira.GetBoard(boardId)
		if err != nil {
			log.Fatalln("fatal:", err)
		}
		fmt.Printf("%#v\n", board)
	}

}

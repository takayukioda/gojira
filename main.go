package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	jira := NewJira(os.Getenv("JIRA_DOMAIN"), os.Getenv("JIRA_EMAIL"), os.Getenv("JIRA_TOKEN"))
	boardId, err := strconv.ParseInt(os.Getenv("JIRA_BOARD_ID"), 10, 0)
	if err != nil {
		log.Fatalln("fatal:", err)
	}

	board, err := jira.GetBoard(boardId)
	if err != nil {
		log.Fatalln("fatal:", err)
	}

	fmt.Println(board)
}

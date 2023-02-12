package data

import (
	"fmt"
	"time"
)

func printHealth() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + "Container is healthy")
}
package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	var found int
	for i:=156218;i<652527;i++ {
		var adjacent bool
		var decreases bool
		password := strconv.Itoa(i)
		for j:=0;j<5;j++ {
			if password[j] == password[j+1] &&
				!strings.Contains(password, string(password[j]) + string(password[j]) + string(password[j])) &&
				!strings.Contains(password, string(password[j]) + string(password[j]) + string(password[j]) + string(password[j])) {
				adjacent = true
			}
			if password[j] > password[j+1] {
				decreases = true
			}
		}

		if adjacent && !decreases {
			found++
		}
	}

	log.Printf("found %d", found)
}

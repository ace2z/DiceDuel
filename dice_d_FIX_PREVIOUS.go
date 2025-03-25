package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	//. "local/INGEST_ENGINE"

	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

var LETTERS = []string{"e", "d", "f", "r"}

func Need_to_FIX_PREVIOUS(input string) bool {

	var found = false
	for _, letter := range LETTERS {
		if strings.Contains(input, letter) {
			found = true
			break
		}
	}

	if found == false {
		return false
	}

	// otherwise, Delete the LAST item from history
	if len(HISTORY) > 0 {
		HISTORY = HISTORY[:len(HISTORY)-1]
	}
	Y.Println("")
	MW.Print(" *** LAST ITEM REMOVED ***")
	Y.Println("")
	Y.Println("")
	return true

}

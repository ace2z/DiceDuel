package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	//. "local/INGEST_ENGINE"

	// "bufio"
	// "os"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

/*
arrows:

⇒
⇢
⇨

↪

➔
➜

➡
➜
➤
*/
var arrow_char = "➤"

func Show_Record(IND int) {

	var DL = HISTORY[IND]

	Y.Print("(", IND, ") ")
	if DL.WINNER != "" {
		W.Print(" ")
		DL.WIN_COLOR.Print(DL.WINNER)

		if strings.Contains(DL.WINNER, "TIE") {
			W.Print("  ")
		} else {
			W.Print("    ")
		}

	} else {
		W.Print("        ")

	}

	M.Print(DL.RED_A)
	G.Print(arrow_char)
	M.Print(DL.RED_B)
	W.Print(" , ")
	C.Print(DL.BLUE_A)
	G.Print(arrow_char)
	C.Print(DL.BLUE_B)
	W.Print(" | ")

	//2. Show the DIf between Red Blue for last two games

	BLUE_YELLOW.Print(" ", DL.RB_DIFF, " ")
	W.Print(" ")

	//3. Iterate through the events. Show them with the colors specified
	for _, evt := range DL.EVENTS {
		evt.COLOR.Print(" ", evt.NAME, " ")
		BOLD_WHITE.Print(" | ")
	}

	W.Println("")
}

package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	. "local/DLOGIC"

	// "bufio"
	// "os"

	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
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

var red_wins_COL = WHITE_RED
var blue_wins_COL = WHITE_BLUE
var tie_wins_COL = BLUE_GREEN

func show_winner(win_tmp string, use_arrows bool) {

	// Show the winner of the hand
	if strings.Contains(win_tmp, "R") {
		if use_arrows {
			W.Print(" ")
			BOLD_MAGENTA.Print(arrow_char + arrow_char)
		}
		WR.Print(" R ")
		W.Print(" ")
	} else if strings.Contains(win_tmp, "B") {
		if use_arrows {
			W.Print(" ")
			BOLD_CYAN.Print(arrow_char + arrow_char)
		}
		blue_wins_COL.Print(" B ")
		W.Print(" ")
	} else if strings.Contains(win_tmp, "T") {
		if use_arrows {
			W.Print(" ")
			BOLD_GREEN.Print(arrow_char + arrow_char)
		}
		tie_wins_COL.Print(" T ")
		W.Print(" ")
	}

}

var arrow_char = "➤"

var even_COLOR = color.New(color.FgHiRed, color.BgHiYellow)
var odd_COLOR = color.New(color.FgHiBlack, color.BgHiWhite, color.Bold)

func Show_HAND(IND int, ACCENT_COLOR *color.Color) {

	var DL = HISTORY[IND]

	pretty_IND := IND + 1

	if pretty_IND < 10 {
		ACCENT_COLOR.Print("( ", pretty_IND, ")")
	} else {
		ACCENT_COLOR.Print("(", pretty_IND, ")")
	}

	// First show the ACTUAL winner of the hand
	W.Print(" ")
	show_winner(DL.WINNER, false)

	//2. Show the DIf between Red Blue for last two games
	W.Print(" ")
	if IS_EVEN(DL.RB_DIFF) || DL.RB_DIFF == 0 {
		even_COLOR.Print(" ", DL.RB_DIFF, " ")
	} else {
		odd_COLOR.Print(" ", DL.RB_DIFF, " ")
	}
	W.Print(" ")
	if IS_EVEN(DL.PREV_RBDIFF) || DL.PREV_RBDIFF == 0 {
		even_COLOR.Print(" ", DL.PREV_RBDIFF, " ")
	} else {
		odd_COLOR.Print(" ", DL.PREV_RBDIFF, " ")
	}

	//3. Now determine if there is a FUTURE hand winner

	// This shows FUTURE WINNER (of the hand following this one ..if it is avaiable)
	if DL.NEXT_WINNER != "" {
		//W.Print(" " + arrow_char + arrow_char)
		show_winner(DL.NEXT_WINNER, true)
	} else {
		W.Print("       ")
	}

	M.Print(DL.RED_A)
	G.Print(arrow_char)
	M.Print(DL.RED_B)
	W.Print(",")
	C.Print(DL.BLUE_A)
	G.Print(arrow_char)
	C.Print(DL.BLUE_B)
	W.Print(" | ")

	//3. Iterate through the events. Show them with the colors specified
	for _, evt := range DL.EVENTS {

		// Search through the EVENTS list, and find the one that matches the name
		// This is a hack we need because of the behavior from LOAD from Disk and SAVE to DISK (via struct as JSON)
		// We have to do it this way because coolors arent not saved to the struct (color.Color is NOT exportable)
		// nor does it have a JSON representation
		// So we have to use the EVENT_LIST to find and print the right color

		for _, k := range EVENT_LIST {
			if k.NAME == evt.NAME {
				k.COLOR.Print(evt.NAME)

				BOLD_WHITE.Print(" ")
				break
			}
		}
	}

	W.Println("")
	RESET_TERM()
}

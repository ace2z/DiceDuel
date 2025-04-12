package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"

	. "local/EVENTS"

	"fmt"

	// "bufio"
	// "os"

	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	"github.com/charmbracelet/lipgloss"
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

// Shows the event in the color/style using LIP GLOSS
func LG_showEvent(evt EVENT_OBJ) {
	var COLOR lipgloss.Style
	var event string

	var lookfor = evt.NAME
	for _, mevl := range EVENT_LIST {

		// Skip if not match
		if CONTAINS(mevl.NAME, lookfor) == false {
			continue
		}

		// Otherwise
		COLOR = mevl.LG_COLOR
		event = mevl.NAME
		break
	}

	//2. If we didnt find a match for the color in the existing EVENT_LIST, lets search the COLOR_MATRIX using COLOR_ID
	if event == "" {
		for _, cm := range COLOR_MATRIX {
			if evt.COLOR_ID == cm.ID {
				COLOR = cm.LG_COLOR
				event = evt.NAME
				break
			}
		}
	}

	// Finally, if event isnt blank, we can display the color

	if event != "" {
		// Otherwise we have a match. Exit when done
		fmt.Print(COLOR.Render(event))
		W.Print(" ")
		return
	}

}

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
	} else if strings.Contains(win_tmp, "B") {
		if use_arrows {
			W.Print(" ")
			BOLD_CYAN.Print(arrow_char + arrow_char)
		}
		blue_wins_COL.Print(" B ")

	} else if strings.Contains(win_tmp, "T") {
		if use_arrows {
			W.Print(" ")
			BOLD_GREEN.Print(arrow_char + arrow_char)
		}
		tie_wins_COL.Print(" T ")
	}

}

func show_RedBlue_DIFF(HND HAND_OBJ) {

	// Search through the events for the RB_DIFF
	var rbdiff = HND.META.RB_DIFF

	W.Print(" ")
	if IS_EVEN(rbdiff) {
		even_COLOR.Print(" ", rbdiff, " ")
	} else {
		odd_COLOR.Print(" ", rbdiff, " ")
	}

	// For the PREVIOUS RB_DIFF
	//var prev_diff = HND.META.PREV_RB_DIFF
	// W.Print(" ")
	// if IS_EVEN(prev_diff) {
	// 	even_COLOR.Print(" ", prev_diff, " ")
	// } else {
	// 	odd_COLOR.Print(" ", prev_diff, " ")
	// }
}

var arrow_char = "➤"
var arrow2 = "↩"
var even_COLOR = color.New(color.FgHiRed, color.BgHiYellow)
var odd_COLOR = color.New(color.FgHiBlack, color.BgHiWhite, color.Bold)

func Show_HAND(IND int, ACCENT_COLOR *color.Color) {

	var HND = HISTORY[IND]

	pretty_IND := IND + 1

	if pretty_IND < 10 {
		ACCENT_COLOR.Print("( ", pretty_IND, ")")
	} else {
		ACCENT_COLOR.Print("(", pretty_IND, ")")
	}

	// First show the ACTUAL winner of the hand
	W.Print(" ")
	show_winner(HND.WINNER, false)

	//2. Show the DIf between Red Blue for last two games
	show_RedBlue_DIFF(HND)

	//3. Now determine if there is a FUTURE hand winner
	// This shows FUTURE WINNER (of the hand following this one ..if it is avaiable)
	// if HND.NEXT_WINNER != "" {

	// 	show_winner(HND.NEXT_WINNER, true)
	// } else {
	// 	W.Print("      ")
	// }

	W.Print(" ")
	M.Print(HND.RED_B)
	G.Print(arrow2)
	M.Print(HND.RED_A)
	W.Print(",")
	C.Print(HND.BLUE_B)
	G.Print(arrow2)
	C.Print(HND.BLUE_A)
	W.Print(" | ")

	//3. Iterate through the events. Show them with the colors specified
	for _, evt := range HND.EVENTS {
		// Dont display events that are hidden
		if evt.SHOW_ME == false {
			continue
		}

		// Search through the EVENTS list, and find the one that matches the name
		// This is a hack we need because of the behavior from LOAD from Disk and SAVE to DISK (via struct as JSON)
		// We have to do it this way because coolors arent not saved to the struct (color.Color is NOT exportable)
		// nor does it have a JSON representation
		// So we have to use the EVENT_LIST or the EVT_COLOR_MATRIX....to find and print the right color/style
		//find_EVENT_ColorStyle_and_PRINT(evt).Print(evt.NAME)

		// Alt using LIP GLOSS
		LG_showEvent(evt)

	}

	W.Println("")
	RESET_TERM()
}

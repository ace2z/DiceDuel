package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"

	. "local/EVENTS"

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

// Gets the color associated with Event. Searches the master EVENTS LIST and
// the COLOR_MATRIX (if applicable)
func find_EVENT_ColorStyle_and_PRINT(evt EVENT_OBJ) *color.Color {
	for _, me := range EVENT_LIST {

		// Skip if not match
		if CONTAINS(evt.NAME, me.NAME) == false {
			continue
		}

		// Otherwise we have a match. Exit when done

		return me.COLOR

	}

	// If we get this far.. means we did NOT find a color define din the EVENT_LIST
	// So now we need to search the COLOR_MATRIX
	for _, cm := range EVT_COLOR_MATRIX {
		// Skip if not match
		if CONTAINS(evt.NAME, cm.NAME) == false {
			continue
		}

		// Otherwise we have a match. Exit when done
		return cm.COLOR
	}

	// If no color is defined, return a default color
	return color.New(color.FgWhite)
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

func show_RedBlue_DIFF(HND HAND_OBJ) {

	// Search through the events for the RB_DIFF
	var rbdiff = NULLV
	var prev_diff = NULLV
	for _, evt := range HND.EVENTS {
		if CONTAINS(evt.NAME, "RB_DIFF") == false {
			continue
		}

		// Otherwise we have a match. Exit when done
		rbdiff = evt.VAL
		prev_diff = evt.PREVVAL
		break
	}
	//error handling
	if rbdiff == NULLV {
		return
	}

	W.Print(" ")
	if IS_EVEN(rbdiff) {
		even_COLOR.Print(" ", rbdiff, " ")
	} else {
		odd_COLOR.Print(" ", rbdiff, " ")
	}
	W.Print(" ")
	if IS_EVEN(prev_diff) {
		even_COLOR.Print(" ", prev_diff, " ")
	} else {
		odd_COLOR.Print(" ", prev_diff, " ")
	}
}

var arrow_char = "➤"

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
	if HND.NEXT_WINNER != "" {
		//W.Print(" " + arrow_char + arrow_char)
		show_winner(HND.NEXT_WINNER, true)
	} else {
		W.Print("       ")
	}

	M.Print(HND.RED_A)
	G.Print(arrow_char)
	M.Print(HND.RED_B)
	W.Print(",")
	C.Print(HND.BLUE_A)
	G.Print(arrow_char)
	C.Print(HND.BLUE_B)
	W.Print(" |")

	//3. Iterate through the events. Show them with the colors specified
	for _, evt := range HND.EVENTS {
		// Dont display events that are hidden
		if evt.SHOW_ME == false {
			continue
		}
		// if this evnt name has spaces in it.. Usually means this is a reverse color type event
		// We need to do an indent
		if strings.Contains(evt.NAME, " ") {
			W.Println(" ")
		}

		// Search through the EVENTS list, and find the one that matches the name
		// This is a hack we need because of the behavior from LOAD from Disk and SAVE to DISK (via struct as JSON)
		// We have to do it this way because coolors arent not saved to the struct (color.Color is NOT exportable)
		// nor does it have a JSON representation
		// So we have to use the EVENT_LIST or the EVT_COLOR_MATRIX....to find and print the right color/style
		find_EVENT_ColorStyle_and_PRINT(evt).Print(evt.NAME)
		BOLD_WHITE.Print(" ")
	}

	W.Println("")
	RESET_TERM()
}

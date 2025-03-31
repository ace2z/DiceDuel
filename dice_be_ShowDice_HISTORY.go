package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	//. "local/INGEST_ENGINE"

	// "bufio"
	// "os"
	// "strings"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

var MAX_HIST_BACK = 10

// func update_with_WINNER(red_win bool, blue_win bool, tie bool, IND int) {
// 	/// Show and save the winner to history
// 	if red_win {

// 		HISTORY[IND].WINNER = " R "
// 		HISTORY[IND].WIN_COLOR = WR
// 	} else if blue_win {

// 		HISTORY[IND].WINNER = " B "
// 		HISTORY[IND].WIN_COLOR = WB

// 		// else its a TIE
// 	} else if tie {

// 		HISTORY[IND].WINNER = " TIE "
// 		HISTORY[IND].WIN_COLOR = BG
// 	}

// 	// Save the
// }

// Shows dice history with the Winner of each (based on the NEXT that posts up)
var DONT_SHOW_HIST = false

func ShowDice_HISTORY(show_all bool) {

	// a little override hack to prevent showing history twice (due to the wway the for loop works)
	if DONT_SHOW_HIST {
		DONT_SHOW_HIST = false
		return
	}
	// Show the last 10 items in the history
	hlen := len(HISTORY)

	if hlen == 0 {
		return
	}

	TITLE := "HISTORY"
	TITLE_COLOR := BOLD_CYAN
	ACCENT_COLOR := BOLD_YELLOW

	if show_all {

		TITLE = "FULL History (all hands/rolls)"
		TITLE_COLOR = MAGENTA_WHITE
		ACCENT_COLOR = BOLD_WHITE
	}
	start_at := len(HISTORY) - MAX_HIST_BACK
	if start_at < 0 || show_all {
		// If we are showing all, then start at 0
		// otherwise, start at the last 10 items
		start_at = 0
	}

	W.Println("")
	ACCENT_COLOR.Print("⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤")
	TITLE_COLOR.Print(" " + TITLE + " ")
	ACCENT_COLOR.Println("⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢")
	W.Println("")
	for x := start_at; x < hlen; x++ {

		// // We need the next item to determine "which side won this item"
		// // basically dont show the first item you save from the history
		// var have_next = false
		// n := x + 1
		// if n < hlen {
		// 	have_next = true
		// }

		// if have_next {
		// 	next := HISTORY[n]

		// 	var red_win = false
		// 	var blue_win = false
		// 	var tie = false

		// 	if next.RED_B > next.BLUE_B {
		// 		red_win = true
		// 	} else if next.BLUE_B > next.RED_B {
		// 		blue_win = true
		// 	} else {
		// 		tie = true
		// 	}

		// 	//3. Since we have a next to look at, we can now save the winner
		// 	update_with_WINNER(red_win, blue_win, tie, x)
		// }

		// Show the actual record (and winner if applicable)
		Show_HAND(x, ACCENT_COLOR)

	}

	W.Println("")
	if show_all {
		ACCENT_COLOR.Print("⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤")
		ACCENT_COLOR.Println("⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢")

	} else {
		ACCENT_COLOR.Print("⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤")
		ACCENT_COLOR.Println("⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢")
	}

}

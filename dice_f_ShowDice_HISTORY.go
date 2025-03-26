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

// func show_Winne_and_SAVE(red_win bool, blue_win bool, tie bool, IND int) {
// 	/// Show and save the winner to history
// 	if red_win {
// 		WR.Print(" R ")
// 		W.Print(" ")

// 		HISTORY[IND].WINNER = " R "
// 		HISTORY[IND].WIN_COLOR = WR
// 	} else if blue_win {
// 		WB.Print(" B ")
// 		W.Print(" ")

// 		HISTORY[IND].WINNER = " B "
// 		HISTORY[IND].WIN_COLOR = WB

// 		// else its a TIE
// 	} else if tie {
// 		BG.Print(" TIE ")
// 		W.Print(" ")

// 		HISTORY[IND].WINNER = " TIE "
// 		HISTORY[IND].WIN_COLOR = WG
// 	}

// 	// Save the
// }

func save_WINNER(red_win bool, blue_win bool, tie bool, IND int) {
	/// Show and save the winner to history
	if red_win {

		HISTORY[IND].WINNER = " R "
		HISTORY[IND].WIN_COLOR = WR
	} else if blue_win {

		HISTORY[IND].WINNER = " B "
		HISTORY[IND].WIN_COLOR = WB

		// else its a TIE
	} else if tie {

		HISTORY[IND].WINNER = " TIE "
		HISTORY[IND].WIN_COLOR = BG
	}

	// Save the
}

// Shows dice history with the Winner of each (based on the NEXT that posts up)
func ShowDice_HISTORY() {
	// Show the last 10 items in the history
	hlen := len(HISTORY)

	if hlen == 0 {
		return
	}

	start_at := len(HISTORY) - MAX_HIST_BACK
	if start_at < 0 {
		start_at = 0
	}

	W.Print("⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤")
	G.Print(" HISTORY ")
	W.Println("⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢")
	for x := start_at; x < hlen; x++ {

		// We need the next item to determine "which side won this item"
		// basically dont show the first item you save from the history
		var have_next = false
		n := x + 1
		if n < hlen {
			have_next = true
		}

		if have_next {
			next := HISTORY[n]

			var red_win = false
			var blue_win = false
			var tie = false

			if next.RED_B > next.BLUE_B {
				red_win = true
			} else if next.BLUE_B > next.RED_B {
				blue_win = true
			} else {
				tie = true
			}

			//3. Since we have a next to look at, we can now save the winner
			save_WINNER(red_win, blue_win, tie, x)
		}

		// Show the actual record (and winner if applicable)
		Show_Record(x)

	}

	W.Print("⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤⥤")
	W.Println("⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢⥢")

}

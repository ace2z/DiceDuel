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

func show_Winner(red_win bool, blue_win bool, tie bool) {
	if red_win {
		WR.Print(" R ")
		W.Print(" ")
	} else if blue_win {
		WB.Print(" B ")
		W.Print(" ")

		// else its a TIE
	} else if tie {
		BG.Print(" TIE ")
		W.Print(" ")
	}
}
func Show_Dice_History() {
	// Show the last 10 items in the history
	hlen := len(HISTORY)

	if hlen == 0 || hlen == 1 {
		return
	}

	start_at := len(HISTORY) - 10
	if start_at < 0 {
		start_at = 0
	}
	for x := start_at; x < hlen; x++ {

		// We need the next item to determine "which side won this item"
		// basically dont show the first item you save from the history
		var have_next = true
		n := x + 1
		if n >= hlen {
			have_next = false
		}

		var red_win = false
		var blue_win = false
		var tie = false

		if have_next {
			next := HISTORY[n]
			if next.RED_B > next.BLUE_B {
				red_win = true
			} else if next.BLUE_B > next.RED_B {
				blue_win = true
			} else {
				tie = true
			}
		}
		Y.Print("(", x, ") ")
		show_Winner(red_win, blue_win, tie)

		SHOW_DICE_Item(HISTORY[x], red_win, blue_win, tie)
	}
}

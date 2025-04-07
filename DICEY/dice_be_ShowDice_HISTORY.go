package DICEY

import (
	// = = = = = Native Libraries

	. "local/CORE"
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

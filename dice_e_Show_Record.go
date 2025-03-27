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

	if DL.ALL_EVEN {
		WHITE_BLUE.Print(" ALL_EVEN ")
		W.Print(" ")
	} else if DL.ALL_ODD {
		WHITE_RED.Print(" ALL_ODD ")
		W.Print(" ")
	}

	W.Print(" |")

	if DL.RED_INC {
		M.Print(" RED_INC")
		W.Print(" |")
	} else if DL.RED_DROP {
		M.Print(" RED_DROP")
		W.Print(" |")
	}

	if DL.HAVE_RED_6 {
		// W.Print(" ")
		// WM.Print(" RED")
		// BY.Print("_6 ")
		// W.Print(" |")
	} else if DL.HAVE_RED_1 {
		W.Print(" ")
		WM.Print(" RED")
		BW.Print("_1 ")
		W.Print(" |")
	}

	if DL.RED_by_1 {
		W.Print(" ")
		WM.Print(" RED*by*ONE ")
		W.Print(" |")
	}
	W.Print(" ███")

	// = = =
	//2. for BLU

	if DL.BLUE_INC {
		C.Print(" BLUE_INC")
		W.Print(" |")
	} else if DL.BLUE_DROP {
		C.Print(" BLUE_DROP")
		W.Print(" |")
	}

	if DL.HAVE_BLUE_6 {
		// W.Print(" ")
		// WB.Print(" BLUE")
		// BY.Print("_6 ")
		// W.Print(" |")
	} else if DL.HAVE_BLUE_1 {
		W.Print(" ")
		WB.Print(" BLUE")
		BW.Print("_1 ")
		W.Print(" |")
	}

	if DL.BLUE_by_1 {
		W.Print(" ")
		WB.Print(" BLU*by*ONE ")
		W.Print(" |")
	}

	W.Println("")
}

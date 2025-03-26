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
	BLUE_YELLOW.Print(" ", DL.B_DIFF, " ")
	W.Print(" ")
	BLUE_WHITE.Print(" ", DL.A_DIFF, " ")
	W.Print(" ")

	WHITE_BLUE.Print(" ", DL.EVEN, "x_EVEN ")
	W.Print(" ")
	WHITE_RED.Print(" ", DL.ODD, "x_ODD ")
	W.Print(" |")

	if DL.RED_INC {
		M.Print(" RED_INC")
		W.Print(" |")
	} else if DL.RED_DROP {
		M.Print(" RED_DROP")
		W.Print(" |")
	}

	// if DL.HAVE_RED_6 {
	// 	M.Print(" RED_6")
	// 	W.Print(" |")
	// } else if DL.HAVE_RED_1 {
	// 	M.Print(" RED_1")
	// 	W.Print(" |")
	// }

	if DL.RED_by_1 {
		W.Print(" ")
		WM.Print(" RED*by*ONE ")
		W.Print(" |")
	}
	W.Print(" ███")

	// = = =
	//2. for BLU
	//M.Print("   ")
	if DL.BLUE_INC {
		C.Print(" BLUE_INC")
		W.Print(" |")
	} else if DL.BLUE_DROP {
		C.Print(" BLUE_DROP")
		W.Print(" |")
	}

	// if DL.HAVE_BLUE_6 {
	// 	C.Print(" BLUE_6")
	// 	W.Print(" |")
	// } else if DL.HAVE_BLUE_1 {
	// 	C.Print(" BLUE_1")
	// 	W.Print(" |")
	// }

	if DL.BLUE_by_1 {
		W.Print(" ")
		WB.Print(" BLU*by*ONE ")
		W.Print(" |")
	}

	// Show the DICE Diff

	W.Println("")
}

package main

import (
	// = = = = = Native Libraries

	//. "local/CORE"

	//. "local/INGEST_ENGINE"

	//	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/rivo/tview" // https://github.com/rivo/tview/tree/master
	// "strings"
	// "time"
	//"github.com/Delta456/box-cli-maker/v2"
	//"github.com/fatih/color"
)

func INT_GetDiff(first int, sec int) int {

	var result = 0
	if first > sec {
		result = first - sec

	} else if sec > first {
		result = sec - first

	}
	return result
}

func main_Dice_Logic(diceval_b_int int, diceval_a_int int, dice_color string, DL *DLOGIC_OBJ) {
	PLACEHOLDER()

	var is_BLUE = false
	if dice_color == "BLUE" {
		is_BLUE = true
	}

	if diceval_b_int > diceval_a_int {

		if is_BLUE {
			DL.BLUE_INC = true
		} else {
			DL.RED_INC = true
		}

	} else if diceval_b_int < diceval_a_int {

		if is_BLUE {
			DL.BLUE_DROP = true
		} else {
			DL.RED_DROP = true
		}
	}

	if diceval_b_int == 6 {

		if is_BLUE {
			DL.HAVE_BLUE_6 = true
		} else {
			DL.HAVE_RED_6 = true
		}

	} else if diceval_b_int == 1 {

		if is_BLUE {
			DL.HAVE_BLUE_1 = true
		} else {
			DL.HAVE_RED_1 = true
		}
	}

	var dv_diff = INT_GetDiff(diceval_a_int, diceval_b_int)
	if dv_diff == 1 {
		if is_BLUE {
			//WB.Print(" BLU*by*ONE ")
			DL.BLUE_by_1 = true
		} else {
			//WM.Print(" RED*by*ONE ")
			DL.RED_by_1 = true
		}
	}
}

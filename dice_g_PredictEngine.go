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




type NUM_OBJ struct {
	NUM int
	COUNT int
}

func Find_Most_Common(item DLOGIC_OBJ) string {
	var NUMS = []NUM_OBJ
	
	NUMS = append(NUMS, NUM_OBJ{item.RED_A, 1})
	NUMS = append(NUMS, NUM_OBJ{item.RED_B, 1})
	NUMS = append(NUMS, NUM_OBJ{item.BLUE_A, 1})
	NUMS = append(NUMS, NUM_OBJ{item.BLUE_B, 1})


	for _, num := range NUMS {
		if num == 1 {
			return "1"
		}
		if num == 6 {
			return "6"
		}
	}
	
}


// Predicts ... by looking back at the history
func Predict_Engine() {

	// Get the LAST  item in the history
	var n = len(HISTORY) - 1
	last := HISTORY[n]

	most_common := Find_Most_Common(last)

	for i, DL := range HISTORY {
		if i == n {
			break
		}

		tmp := HISTORY[i]

		if 

	}

	
}

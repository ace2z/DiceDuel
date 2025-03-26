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
	NUM   int
	COUNT int
}

func Find_Most_Common(item DLOGIC_OBJ) []int {
	PLACEHOLDER()
	var NUMS = []NUM_OBJ{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
	}

	var TMPS = []int{
		item.RED_A,
		item.RED_B,
		item.BLUE_A,
		item.BLUE_B,
	}

	for _, num := range TMPS {
		for i, obj := range NUMS {
			if obj.NUM == num {
				NUMS[i].COUNT++
				break
			}
		}
	}
	// Find most common
	//highest := 0
	high_list := []int{}

	for _, obj := range NUMS {
		if obj.COUNT >= 2 {
			high_list = append(high_list, obj.NUM)
		}
	}

	// SHOW_STRUCT(high_list)
	// PressAny()
	return high_list
}

// Predicts ... by looking back at the history
func Predict_Engine() {

	// Get the LAST  item in the history
	var hlen = len(HISTORY)
	var n = hlen - 1

	if n < 0 {
		return
	}

	curr := HISTORY[n]

	//2. Get the most common occuring number in this dice game
	var common_nums = Find_Most_Common(curr)
	//3. Get the red Blue Diff
	//var rb_diff = curr.RB_DIFF

	//3b error handling
	if len(common_nums) == 0 {
		return
	}

	//4. Now, lets look at the last 10 games.. and see if we can predict the next one
	var start_at = n - MAX_HIST_BACK
	if start_at < 0 {
		start_at = 0
	}

	found_index := -69

	final_ranking := 0
	for x := start_at; x < hlen; x++ {
		tmp := HISTORY[x]

		ranking := 0
		// if tmp.RB_DIFF == rb_diff {
		// 	ranking++
		// }
		tmp_common := Find_Most_Common(tmp)

		com_rank := 0
		for _, num := range common_nums {
			for _, tmp_num := range tmp_common {
				if num == tmp_num {
					com_rank++
				}
			}
		}

		if com_rank >= 2 {
			ranking = ranking + com_rank
			final_ranking = ranking
			found_index = x
			break
		}
	}

	if found_index >= 0 {
		G.Println(" PREDICTION: ")
		// Show the next item

		Show_Record(found_index)
		W.Println(" vs ")
		Show_Record(n)

		M.Print(" Ranking: ")
		G.Println(final_ranking)
	}

	// for i, DL := range HISTORY {
	// 	if i == n {
	// 		break
	// 	}

	// }
}

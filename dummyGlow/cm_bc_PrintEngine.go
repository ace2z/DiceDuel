package main

import (
	// = = = = = Native Libraries
	"strings"

	. "github.com/ace2z/GOGO/REFGADS/BLOX"
	. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//"github.com/fatih/color"
	//"sort"
	//"strconv"
	//"fmt"
	//"strings"
)

// *************************** MAIN AREA ****************************

func Print_ENGINE(ALL_PARAMS ...interface{}) {
	PLACEHOLDER()
	COLOR_PLACEHOLDER()

	// var fgnum = -9
	// var bgnum = -9
	// var attrs = []string{}

	// //var use_passed_objs = false
	// var MSG_LINES = []string{} // holds all the lines of the users message to be printed

	var all_except []interface{}
	var newline = true
	for _, param := range ALL_PARAMS {
		// int_val, isINT := param.(int)
		str_val, isSTR := param.(string)

		// If they asked for a newline
		if is_FuncParam(param, NOEOL) {
			newline = false
			continue
		}

		if isSTR {
			if strings.HasPrefix(str_val, color_PREF) {
				continue
			}

			all_except = append(all_except, param)
		}

	}

	//2. Generate a new Color object from the params passed
	nc := NewColor(ALL_PARAMS...)

	if newline {
		nc.Println(ALL_PARAMS...)
	} else {
		nc.Print(ALL_PARAMS...)
	}
}

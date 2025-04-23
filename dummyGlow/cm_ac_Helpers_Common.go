package main

import (
	// = = = = = Native Libraries
	. "github.com/ace2z/GOGO/REFGADS/BLOX"
	. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//"github.com/fatih/color"
	//"sort"
	//"strconv"
	//"strings"
	//"fmt"
	"strings"
)

// *************************** MAIN AREA ****************************

func is_FuncParam(paramITEM any, extras ...string) bool {
	PLACEHOLDER()
	COLOR_PLACEHOLDER()
	// SHOW_STRUCT(paramITEM, CYAN)
	// SHOW_STRUCT(extras, YELLOW)
	// PressAny()
	// defaults we use
	var PARAM_PREFIX_LIST = []string{
		"--",
		"-",
	}

	// // Otehrwise if there is something in param_PREF we use this instead
	// // Seperate multiple values witha  PIPE
	// if param_PREF != "" {
	// 	msplit := strings.Split(param_PREF, "|")
	// 	if len(msplit) > 1 {
	// 		LOOK_LIST = msplit
	// 		// Else they have a single value (either a character or pattern)
	// 	} else {
	// 		LOOK_LIST = []string{} // clear out the default
	// 		LOOK_LIST = append(LOOK_LIST, param_PREF)
	// 	}
	// }
	// W.Println("param_PREF: ", param_PREF)
	// C.Println("LOOK_LIST: ")
	// SHOW_STRUCT(LOOK_LIST, CYAN)
	// PressAny()
	str_val, isSTR := paramITEM.(string)
	if isSTR == false {
		//M.Println(" No String param passed!!")
		return false
	}
	var param_2check = ""
	var lookfor = ""
	var is_valid_named_param = false

	for _, pref := range PARAM_PREFIX_LIST {
		// W.Println("")
		// W.Println(" pref: " + pref)
		// C.Println("    S: " + str_val)

		if strings.HasPrefix(str_val, pref) {
			// If there are any spaces in the string, we know we DONT have a valid param
			if strings.Contains(str_val, " ") {
				//	M.Println(" Invalid Param FOUND!!")
				return false
			}
			//param_name = strings.ReplaceAll(str_val, look, "")
			param_2check = str_val
			is_valid_named_param = true
			//G.Println(" Valid Param FOUND!!")
			break
		}
	}

	// Now lets go through the extras and see if they passed somethign we can set lookfor to
	//SHOW_STRUCT(extras, YELLOW)
	for _, look := range extras {
		lookfor = look
	} //end of params
	//end of params
	// G.Println(" param_2check: **" + param_2check + "** ")
	// G.Println("      Lookfor: **" + lookfor + "**")
	// C.Println("")
	//PressAny()
	//G.Println(" Now param_name: ", param_name+"** ")

	if param_2check != "" && lookfor != "" {
		//2. Now lets compare and see if lookfor matches the param_name
		if param_2check == lookfor {
			return true
		}

		// Otherwise if lookfor wasnt specified, we just want to check that this was a valid param PERIOD
		// Which is true if param_name has a value
	} else if is_valid_named_param {
		return true
	}

	return false
}

func not_FuncParam(paramITEM any, extras ...string) bool {

	result := is_FuncParam(paramITEM, extras...)
	if result == false {
		return true
	}

	return false
}

package main

import (
	// = = = = = Native Libraries
	//. "github.com/ace2z/GOGO/REFGADS/BLOX"
	//. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//"github.com/fatih/color"
	//"sort"
	//"strings"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	//"strconv"
	"strings"
)

// *************************** MAIN AREA ****************************

func PrintColorEngine(message string, ALL_PARAMS ...any) {
	var msg = message
	var tmplist = []string{}

	var add_newline = false
	for _, param := range ALL_PARAMS {
		str_val, isSTR := param.(string)
		if isSTR == false {
			continue
		}
		str_val = strings.TrimSpace(str_val)
		// For Println behavior
		if strings.HasPrefix(str_val, ADD_CRLF) {
			add_newline = true
			continue
		}
		if strings.HasPrefix(str_val, "-") {
			continue
		}

		//Otherwise, If they pass a string, assume it is a number (for color) or a word like bold, under etc
		tmplist = append(tmplist, str_val)
	}

	// = = = =
	//4. The default behavior assumes user passed items on a command line
	// Build the style based on ATTRs we find. The first numeric we find is assumed to be the FOREGROUND color
	var style = lipgloss.NewStyle()

	//5. See if the items in tmplist are items that match COLOR_MATRIX DESC or ID
	var matrix []string

	for _, tmp := range tmplist {
		for _, mat := range COLOR_MATRIX {
			if tmp == mat.DESC || tmp == mat.ID {
				matrix = append(matrix, mat.ATTRS...)
			}
		} //end of for
	} //end of for
	use_matrix := len(matrix) > 0

	//5b. If found, we generate our color based on what is in the matrix
	// If there are TWO items from the matrix passed, we assume they are the Fore and Back colors

	if use_matrix {
		set_Style_from_List(&style, &matrix)

		//6. Otherwise, we try to assemble the color absed on the items passed (manually)
	} else if use_matrix == false && len(tmplist) > 0 {
		set_Style_from_List(&style, &tmplist)

		// 8 RARE but if we get this far, we dont have anything valid in tmplist OR the matrix
	} else {
		return
	}

	// 10. Finally PRINT IT
	fmt.Print(style.Render(msg))

	// and for the Println behavior, we add a new line
	if add_newline == true {
		fmt.Println("")
	}
}

// alias to PrintColorEngine
func Print(message string, ALL_PARAMS ...any) {
	PrintColorEngine(message, ALL_PARAMS...)
}
func Println(message string, ALL_PARAMS ...any) {
	ALL_PARAMS = append(ALL_PARAMS, ADD_NEWLINE)
	PrintColorEngine(message, ALL_PARAMS...)
}

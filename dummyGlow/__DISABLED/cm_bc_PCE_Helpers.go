package main

import (
	// = = = = = Native Libraries
	// . "github.com/ace2z/GOGO/REFGADS/BLOX"
	// . "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//"github.com/fatih/color"
	//"sort"
	//"strings"
	//	"fmt"

	"github.com/charmbracelet/lipgloss"
	"strconv"
)

// *************************** MAIN AREA ****************************

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func set_Fore_and_Back(attr string, fg_col *string, bg_col *string) {
	if *fg_col == "" {
		if isNumber(attr) {
			*fg_col = attr
		}
		return
	}
	if *bg_col == "" {
		if isNumber(attr) {
			*bg_col = attr
		}
	}
}
func set_ATTRs(attr string, style *lipgloss.Style) {
	//3. This handles all the other attributes
	switch attr {
	case "bold", "hi":
		*style = style.Bold(true)
	case "under":
		*style = style.Underline(true)
	case "blink":
		*style = style.Blink(true)
	case "reverse":
		*style = style.Reverse(true)
		// case "italic":
		// 	*style = style.Italic(true)
		// case "strike":
		// 	*style = style.Strikethrough(true)
		// case "cross":
		// 	*style = style.Strikethrough(true)
	}
}

func set_Style_from_List(style *lipgloss.Style, tmplist *[]string) {
	var fg_col = ""
	var bg_col = ""
	for _, attr := range *tmplist {

		// Sets the fore and background based on the first two number attributes we find
		set_Fore_and_Back(attr, &fg_col, &bg_col)

		set_ATTRs(attr, style)
	} //end of for

	// Now if we have both of them, we can set the style
	if fg_col != "" {
		*style = style.Foreground(lipgloss.Color(fg_col))
	}
	if bg_col != "" {
		*style = style.Background(lipgloss.Color(bg_col))
	}

}

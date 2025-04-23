package main

import (
	// = = = = = Native Libraries
	. "github.com/ace2z/GOGO/REFGADS/BLOX"
	. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	// "github.com/fatih/color"
	"fmt"
	"io"
	"os"
	"strings"
	// "strconv"
	// "strings"
	// "fmt"
	// "strings"
)

// *************************** MAIN AREA ****************************
func dummy_holder() {
	PLACEHOLDER()
	COLOR_PLACEHOLDER()
}

// NewColorableStdout returns new instance of Writer which handles escape sequence for stdout.
func NewColorableStdout() io.Writer {
	return os.Stdout
}

// NewColorableStderr returns new instance of Writer which handles escape sequence for stderr.
func NewColorableStderr() io.Writer {
	return os.Stderr
}

var (
	Output = NewColorableStdout()
)

//	func (c *CLR_OBJ) format() string {
//		return fmt.Sprintf("%s[%sm", escape, c.sequence())
//	}
func Unset_and_Cleanup() {
	// If we are in black and white mode, we don't need to do anything
	if NO_COLOR {
		return
	}

	//fmt.Fprintf(Output, "%s[%dm", escape, Reset)
	fmt.Fprintf(Output, "%s", ec_RESET)

}
func (c *CLR_OBJ) unset() {
	Unset_and_Cleanup()
}

func SetColor_and_Attribs(c *CLR_OBJ) {

	var formatted_line = fmt.Sprintf("%s%d%s", fg_256mode, c.FG, escEND)
	formatted_line += fmt.Sprintf("%s%d%s", fg_256mode, c.BG, escEND)
	for _, tmpattr := range c.ATTRS {
		if is_FuncParam(tmpattr) {
			// Split the string into two parts
			parts := strings.Split(tmpattr, "__")
			if len(parts) > 1 {
				attr_val := parts[1]
				formatted_line += fmt.Sprintf("%s", attr_val)
			}
		}
	}
	//SHOW_STRUCT(formatted_line, WHITE)
	fmt.Fprint(Output, formatted_line)
}

// func (c *CLR_OBJ) SetColor_and_Attribs() *CLR_OBJ {

//		var formatted_line = fmt.Sprintf("%s%d%s", escFG, c.FG, escEND)
//		formatted_line += fmt.Sprintf("%s%d%s", escBG, c.BG, escEND)
//		for _, tmpattr := range c.ATTRS {
//			if is_FuncParam(tmpattr) {
//				// Split the string into two parts
//				parts := strings.Split(tmpattr, "__")
//				if len(parts) > 1 {
//					SHOW_STRUCT(parts, WHITE)
//					attr_val := parts[1]
//					formatted_line += fmt.Sprintf("%s", attr_val)
//				}
//			}
//		}
//		SHOW_STRUCT(formatted_line, WHITE)
//		fmt.Fprint(Output, formatted_line)
//		return c
//	}
func (c *CLR_OBJ) Print(user_data ...interface{}) (int, error) {

	SetColor_and_Attribs(c)
	defer c.unset()

	return fmt.Fprint(Output, user_data...)
}

func (c *CLR_OBJ) Println(user_data ...interface{}) (int, error) {

	SetColor_and_Attribs(c)
	defer c.unset()

	return fmt.Fprintln(Output, user_data...)

}

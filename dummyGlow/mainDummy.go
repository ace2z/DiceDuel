package main

import (
	// = = = = = Native Libraries
	. "github.com/ace2z/GOGO/REFGADS/BLOX"
	//. "github.com/ace2z/GOGO/REFGADS/ColorOPS"
	//. "github.com/ace2z/GOGO/REFGADS/BLOX"
	//"github.com/fatih/color"
	//. "github.com/ace2z/GOGO/REFGADS/PRETTY"
	//"github.com/fatih/color"
	//"reflect"
	//"strings"
	//"os/exec"
)

// *************************** MAIN AREA ****************************
var VERSION_NUM = ""

func main() {

	//1. Init the program (and the command line parameters)
	// lines := []string{
	// 	"The DUMMY ran: ",
	// 	"Swimmingly!!",
	// 	"graciously",
	// 	"POWERFUlly",
	// }

	INIT_Main(VERSION_NUM, "a Dummy Program for testing")

	//Determine ifw e are currently in a GIT repo
	//2. Make sure GIT is installed (and we are currently IN  a GIT repo)
	// cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	// err2 := cmd.Run()

	// // no error means we are GOOD
	// if err2 == nil {

	// 	G.Println(" Currently in a GIT repo!!!")
	// } else {
	// 	M.Println(" Not in a GIT repo!")
	// }

	//SILENT_EXIT()

	// PREFIX("The Dummy Ran: ", NOEOL, BOLD_CYAN, "-prefix|***")
	// BOLD_MAGENTA.Println("Amazingly!!")
	// C.Println("Next Post")
	// err := errors.New("UNABLE  to ENCODE to Json, invalid fields! dummy error")
	// DO_EXIT(" FAKE Error but also, nvalid Form of ID", err)

	COLOR_Matrix_Engine("Dummy")

}

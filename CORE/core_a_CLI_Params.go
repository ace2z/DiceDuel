package CORE

import (
	"flag"
	//. "local/_API_MBOUM"
	//. "github.com/ace2z/HP_COMMON"
)

var INPUT_RED_DICE, INPUT_BLUE_DICE string

var LOAD_FILE_NUM = 0

func CLI_PARAM_INIT() {

	// flag.BoolVar(&REINIT_ALL, "reinit", REINIT_ALL, "  Reinit The Core (DB etc)")
	// flag.BoolVar(&PURGE_DATA, "purge", PURGE_DATA, "  purge ALL data from the database")

	// flag.BoolVar(&USE_DB_DELTA, "dbdelta", USE_DB_DELTA, "  DELTA METHOD: When writing to DB, ONLY save items that dont ALREADY exist")
	// flag.BoolVar(&USE_DB_UPSERT, "dbupsert", USE_DB_UPSERT, "  When write to DB FORCE use UPSERT method")

	// flag.IntVar(&MAX_DAYS_BACK, "max", MAX_DAYS_BACK, "  Max days BACK for the")
	// flag.IntVar(&MAX_DAYS_BACK, "daysback", MAX_DAYS_BACK, "  Max days --alias")

	flag.StringVar(&INPUT_RED_DICE, "red", INPUT_RED_DICE, " Value for RED dice (2 numbers from RIGHT to LEFT) ")
	flag.StringVar(&INPUT_BLUE_DICE, "blue", INPUT_BLUE_DICE, " Value for BLUE dice (2 numbers from RIGHT to LEFT) ")

	flag.IntVar(&LOAD_FILE_NUM, "load", LOAD_FILE_NUM, "  Automatically Loads the FILE at this number")

	//flag.StringVar(&ONLY_TF, "tf", ONLY_TF, " Only process for THIS TimeFrame (ie 1min, HOURLY, DAILY)")
	//  flag.IntVar(&DEFAULT_MAX_THREADS,  "threads", DEFAULT_MAX_THREADS,      "  Max THREADS to use at a time")

}

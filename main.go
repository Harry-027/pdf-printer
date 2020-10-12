package main

import (
	"github.com/Harry-027/pdf-printer/cmd"
	"github.com/Harry-027/pdf-printer/db"
	"github.com/Harry-027/pdf-printer/utils"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

func main() {
	home, err := homedir.Dir()                   // Fetch the current user home dir.
	utils.PanicErr(err)                          // Panic in case user dir not available
	dbPath := filepath.Join(home, utils.DB_FILE) // set the db file path for data storage
	err = db.Init(dbPath)                        // Initialize & connect with db
	utils.PanicErr(err)                          // Panic in case any issue while db connection
	err = cmd.RootCmd.Execute()                  // Execute the root command
	utils.PanicErr(err)                          // Panic in case any issue
}

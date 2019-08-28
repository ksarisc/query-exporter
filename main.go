package main

import (
	"fmt"
	"os"

	lib "./lib"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	argc := len(os.Args)
	aset := "exampleSettings.json"
	if argc > 1 {
		aset = os.Args[argc-1]
	}
	fmt.Printf("Testing: %s\n", aset)

	conf, err := lib.ParseAndCheckConfig(aset)
	if err != nil {
		fmt.Printf("Config Error: %s\n", err)
		return
	}

	fmt.Print("App Configuration\nDatabase\n")
	dbas := conf.Database
	fmt.Printf("Connect:   %s\nBuild SQL: %s\nGet SQL:   %s\nSet SQL:   %s\n",
		dbas.Connection, dbas.BuildSQL, dbas.GetSQL, dbas.SetSQL)
	fmt.Printf("Output Path: %s\n", conf.OutputPath)

	fmt.Println("Beginning Query to File")
	err = lib.QueryToFile(conf.Database, conf.OutputPath)
	if err != nil {
		fmt.Printf("Query to File Error: %s\n", err)
	} else {
		//fmt.Printf("File Written: %s\n", conf.OutputPath)
		fmt.Println("Complete")
	}
} // END main

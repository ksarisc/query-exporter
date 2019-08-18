package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	lib "./lib"
)

func main() {
	argc := len(os.Args)
	aset := "exampleSettings.json"
	if argc > 1 {
		aset = os.Args[argc-1]
	}
	fmt.Printf("Testing: %s\n", aset)

	file, err := os.Open(aset)
	if err != nil {
		fmt.Printf("File Open Error: %s\n", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Settings Error: %s\n", err)
		return
	}
	var conf lib.AppConfig
	json.Unmarshal(data, &conf)
	fmt.Print("App Configuration\nDatabase\n")
	dbas := conf.Database
	fmt.Printf("Connect:   %s\nBuild SQL: %s\nGet SQL:   %s\nSet SQL:   %s\n",
		dbas.Connection, dbas.BuildSQL, dbas.GetSQL, dbas.SetSQL)
	fmt.Printf("Output Path: %s\n", conf.OutputPath)
}

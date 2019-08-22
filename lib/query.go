package lib

import (
	"database/sql"
	"fmt"
	"net"
)

//"io/ioutil"

// ServerToAddress for SQL server, translate hostname to ip address
func ServerToAddress(hostname string) string {
	if StringIsWhitespace(hostname) || hostname == "." {
		return "127.0.0.1"
	}
	ips, err := net.LookupIP(hostname)
	if err != nil {
		//fmt.Fprintf()
		return "127.0.0.1"
	}
	for _, addr := range ips {
		fmt.Print("## IN A #%s#", hostname, addr.String())
	}
	return ips[0].String()
} // END ServerToAddress

// QueryToFile builds the file defined in the AppConfig
func QueryToFile(config DataConfig, fileName string) error {
	// open connection
	conn, err := sql.Open(config.Provider, config.Connection)
	if err != nil {
		return err
	}
	defer conn.Close()
	// run build
	if !StringIsWhitespace(config.BuildSQL) {
		_, err := conn.Exec(config.BuildSQL)
		if err != nil {
			return fmt.Errorf("Build SQL Error: %s", err)
		}
		//rslt.RowsAffected()
	}
	// run get
	rows, err := conn.Query(config.GetSQL)
	if err != nil {
		return err
	}
	// columns of get
	var cols []string
	if len(config.GetFields) > 0 {
		cols = config.GetFields
	} else {
		cols, err = rows.Columns()
		if err != nil {
			return err
		}
	}
	// build the file
	if err := loadFile(rows, cols, fileName); err != nil {
		return err
	}
	// run set
	if !StringIsWhitespace(config.SetSQL) {
		_, err := conn.Exec(config.SetSQL)
		if err != nil {
			return fmt.Errorf("Set SQL Error: %s", err)
		}
		//rslt.RowsAffected()
	}
	return nil
} // END RunQuery

func loadFile(rows *sql.Rows, cols []string, path string) error {
	collen := len(cols)
	scans := make([]interface{}, collen)
	rvals := make([][]byte, collen)
	for i := 0; i < collen; i++ {
		scans[i] = &rvals[i]
	}
	// loop through data
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			return err
		}
		// copy performed here
		//for i, bytes := range rvals {
		for i := 0; i < collen; i++ {
			if i > 0 {
				fmt.Print(", ")
			}
			//if bytes != nil { fmt.Printf("%s", bytes)
			// lookup vs copy cost?
			if rvals[i] != nil {
				fmt.Printf("%s", rvals[i])
			}
		}
		fmt.Println()
	}
	return rows.Err()
} // END loadFile

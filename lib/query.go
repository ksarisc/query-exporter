package lib

import (
	"bufio"
	"database/sql"
	"fmt"
	"net"
	"os"
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
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if err := loadFile(rows, cols, writer); err != nil {
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

func loadFile(rows *sql.Rows, cols []string, writer *bufio.Writer) error {
	collen := len(cols)
	scans := make([]interface{}, collen)
	rvals := make([][]byte, collen)
	// what about outputting column names
	for i := 0; i < collen; i++ {
		//count, err := w.WriteString("some data\n")
		if i > 0 {
			writer.WriteString("|")
		}
		writer.WriteString(cols[i])
		scans[i] = &rvals[i]
	}
	writer.WriteString("\n")
	// loop through data
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			return err
		}
		// copy performed here
		//for i, bytes := range rvals {
		for i := 0; i < collen; i++ {
			if i > 0 {
				writer.WriteString("|")
			}
			//if bytes != nil { fmt.Printf("%s", bytes)
			// lookup vs copy cost?
			if rvals[i] != nil {
				//fmt.Fprintf(writer, "%s", rvals[i])
				writer.Write(rvals[i])
			}
		}
		writer.WriteString("\n")
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return writer.Flush()
} // END loadFile

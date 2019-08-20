package lib

//https://github.com/denisenkom/go-mssqldb
import (
	"database/sql"
	"fmt"
	"log"
	//_ "github.com/denisenkom/go-mssqldb"
)

// QueryToFile builds the file defined in the AppConfig
func QueryToFile(config DataConfig, fileName string) error {
	// validate parameters
	if StringIsWhitespace(config.Connection) {
		return fmt.Errorf("SQL Connection configuration value REQUIRED")
	}
	if StringIsWhitespace(config.GetSQL) {
		return fmt.Errorf("Get SQL configuration value REQUIRED")
	}
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
	var (
		sqlversion string
	)
	rows, err := conn.Query(config.GetSQL)
	if err != nil {
		return err
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			return err
		}
		log.Println(sqlversion)
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

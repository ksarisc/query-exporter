package lib

//https://github.com/denisenkom/go-mssqldb
import (
	"database/sql"
	"log"
)

// AppConfig stores entire configuration defined
// in exampleSettings.json
type AppConfig struct {
	Database   DataConfig `json:"database"`
	OutputPath string     `json:"outputPath"`
}

// DataConfig stores Database section of AppConfig
type DataConfig struct {
	Connection string `json:"connection"`
	BuildSQL   string `json:"buildSql"`
	GetSQL     string `json:"getSql"`
	SetSQL     string `json:"setSql"`
}

// QueryToFile builds the file defined in the AppConfig
func QueryToFile(config DataConfig, output string) error {
	conn, err := sql.Open("mssql",
		"server=localhost;user id=sa;password=SA_PASSWORD=yourStrong(!)Password;")
	if err != nil {
		return err
	}
	defer conn.Close()
	getSql := config.GetSQL
	if StringIsWhitespace(getSql) {
		getSql = config.BuildSQL
	}
	var (
		sqlversion string
	)
	rows, err := conn.Query(getSql)
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
	return nil
} // END RunQuery

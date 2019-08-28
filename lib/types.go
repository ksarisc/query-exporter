package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// AppConfig stores entire configuration defined
// in exampleSettings.json
type AppConfig struct {
	Database   DataConfig `json:"database"`
	OutputPath string     `json:"outputPath"`
}

// DataConfig stores Database section of AppConfig
type DataConfig struct {
	Connection string   `json:"connection"`
	Provider   string   `json:"provider"`
	BuildSQL   string   `json:"buildSql"`
	GetSQL     string   `json:"getSql"`
	GetFields  []string `json:"getFields"` //DataField `json:"getFields"`
	SetSQL     string   `json:"setSql"`
}

//type DataField struct { }

// ParseAndCheckConfig read json file into AppConfig struct
// and return any errors with configuration
func ParseAndCheckConfig(fileName string) (AppConfig, error) {
	data, err := readFile(fileName)
	if err != nil {
		fmt.Printf("Settings Error: %s\n", err)
		return AppConfig{}, err
	}
	var conf AppConfig
	json.Unmarshal(data, &conf)
	err = ConfigHasRequiredValues(&conf)
	// if err != nil {
	// 	fmt.Printf("Config Error: %s\n", err)
	// 	return AppConfig{}, err
	// }
	return conf, err
} // END ParseAndCheckConfig

// ConfigHasRequiredValues checks whether configuration
// has specific fields filled in
func ConfigHasRequiredValues(config *AppConfig) error {
	// validate app parameters
	if StringIsWhitespace(config.OutputPath) {
		return fmt.Errorf("Output Path configuration value REQUIRED")
	}
	// check the directory exists

	// validate data parameters
	if StringIsWhitespace(config.Database.Connection) {
		return fmt.Errorf("SQL Connection configuration value REQUIRED")
	}
	if StringIsWhitespace(config.Database.Provider) {
		//return fmt.Errorf("SQL Connection configuration value REQUIRED")
		config.Database.Provider = "mssql"
	}
	if StringIsWhitespace(config.Database.GetSQL) {
		return fmt.Errorf("Get SQL configuration value REQUIRED")
	}
	// load files for SQL statements as needed
	if isSQLFile(config.Database.BuildSQL) {
		data, err := readFile(config.Database.BuildSQL)
		if err != nil {
			return err
		}
		config.Database.BuildSQL = string(data)
	}
	if isSQLFile(config.Database.GetSQL) {
		data, err := readFile(config.Database.GetSQL)
		if err != nil {
			return err
		}
		config.Database.GetSQL = string(data)
	}
	if isSQLFile(config.Database.SetSQL) {
		data, err := readFile(config.Database.SetSQL)
		if err != nil {
			return err
		}
		config.Database.SetSQL = string(data)
	}
	return nil
} // END ConfigHasRequiredValues

func isSQLFile(value string) bool {
	vlen := len(value) - 1
	if vlen < 5 {
		return false
	}
	for i := vlen; i >= 0; i-- {
		if value[i] == '.' {
			// check file end
			if (i + 4) < vlen {
				v1 := value[i+1]
				v2 := value[i+2]
				v3 := value[i+3]
				return (v1 == 's' || v1 == 'S') &&
					(v2 == 'q' || v2 == 'Q') &&
					(v3 == 'l' || v3 == 'L')
			}
			return false
		}
	}
	return false
} // END isSQLFile

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
} // END readFile

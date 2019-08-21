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
	file, err := os.Open(fileName)
	if err != nil {
		//fmt.Printf("File Open Error: %s\n", err)
		return AppConfig{}, err
	}
	data, err := ioutil.ReadAll(file)
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
	return nil
} // END ConfigHasRequiredValues

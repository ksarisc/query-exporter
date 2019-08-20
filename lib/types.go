package lib

import "fmt"

// AppConfig stores entire configuration defined
// in exampleSettings.json
type AppConfig struct {
	Database   DataConfig `json:"database"`
	OutputPath string     `json:"outputPath"`
}

// DataConfig stores Database section of AppConfig
type DataConfig struct {
	Connection string `json:"connection"`
	Provider   string `json:"provider"`
	BuildSQL   string `json:"buildSql"`
	GetSQL     string `json:"getSql"`
	SetSQL     string `json:"setSql"`
}

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
		config.Provider = "mssql"
	}
	if StringIsWhitespace(config.Database.GetSQL) {
		return fmt.Errorf("Get SQL configuration value REQUIRED")
	}
	return nil
} // END ConfigHasRequiredValues

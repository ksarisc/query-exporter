package lib

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

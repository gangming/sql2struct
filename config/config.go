package config

var Cnf *Config

type Config struct {
	Tables      []string `yaml:"tables"`        // Need to generate tables, default is all tables
	DBType      string   `yaml:"db_type"`       // default: types.DriverTypeMysql
	DSN         string   `yaml:"dsn"`           // data source name (DSN) to use when connecting to the database
	TablePrefix string   `yaml:"table_prefix"`  // table prefixed with the table name
	DBTag       string   `yaml:"db_tag"`        // db tag. default: gorm
	WithJsonTag bool     `yaml:"with_json_tag"` // with json tag. default: true
	OutputDir   string   `yaml:"output_dir"`    // output dir. default: ./model
	Scheme      string   `yaml:"scheme"`        // scheme. default: "public"
}

func init() {
	Cnf = &Config{
		DBType:      "mysql",
		Tables:      []string{},
		DBTag:       "gorm",
		WithJsonTag: true,
		OutputDir:   "./model",
	}
}

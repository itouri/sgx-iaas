package mongo

type AppConfig struct {
	Name  string
	Port  uint
	Debug bool
}

type DbConfig struct {
	Host     string
	Port     uint
	Database string
}

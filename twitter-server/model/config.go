package model

var Configurationins AppConfigs

type AppConfigs struct {
	Server struct {
		Mode string `yaml:"mode"`
		Port string `yaml:"port"`
	}
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Sslmode  string `yaml:"sslmode"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	}
	Migration struct {
		Sourceurl   string `yaml:"sourceurl"`
		Databaseurl string `yaml:"databaseurl"`
	}
}
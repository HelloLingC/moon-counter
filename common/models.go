package common

type Config struct {
	Host      string   `yaml:"host"`
	Port      int      `yaml:"port"`
	ImgTheme  string   `yaml:"imgTheme"`
	Cors      bool     `yaml:"cors"`
	Hostnames []string `yaml:"hostnames"`
	DBCfg     DBConfig `yaml:"db"`
}

type DBConfig struct {
	Type     string `yaml:"type"`
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

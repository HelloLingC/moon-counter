package common

type Config struct {
	Host      string   `yaml:"host"`
	Port      int      `yaml:"listen"`
	ImgTheme  string   `yaml:"img_yheme"`
	Cors      bool     `yaml:"cors"`
	Hostnames []string `yaml:"hostnames"`
	ErrorLog  string   `yaml:"error_log"`
	DBCfg     DBConfig `yaml:"db"`
}

type DBConfig struct {
	Type     string `yaml:"type"`
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

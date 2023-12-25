package common

type Config struct {
	Host      string      `yaml:"host"`
	Port      int         `yaml:"listen"`
	ImgTheme  string      `yaml:"img_theme"`
	Cors      bool        `yaml:"cors"`
	Hostnames []string    `yaml:"hostnames"`
	ErrorLog  string      `yaml:"error_log"`
	DBCfg     DBConfig    `yaml:"db"`
	AdminCfg  AdminConfig `yaml:"admin"`
}

type DBConfig struct {
	Type     string `yaml:"type"`
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AdminConfig struct {
	Enabled    bool   `yaml:"enable"`
	Path       string `yaml:"path"`
	Passphrase string `yaml:"passphrase"`
}

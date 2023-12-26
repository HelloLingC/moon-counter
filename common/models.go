package common

import "time"

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
	GuestLogin string `yaml:"guest_login_password"`
}

type Counter struct {
	Id          int       `db:"id"`
	Identifier  string    `db:"identifier"`
	Count       int       `db:"count"`
	CreatedTime time.Time `db:"created_id"`
}

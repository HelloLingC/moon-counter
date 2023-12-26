package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/HelloLingC/moon-counter/common"
	"github.com/HelloLingC/moon-counter/database"
	"github.com/HelloLingC/moon-counter/server"
	"gopkg.in/yaml.v2"
)

func main() {
	configPath := flag.String("c", "config.yaml", "The path to config file")
	flag.Parse()

	cfgFile, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	var config common.Config
	err = yaml.Unmarshal(cfgFile, &config)
	if err != nil {
		log.Fatal("Error unmarshaling YAML:", err)
	}

	db, err := database.NewDBAdapter(config.DBCfg.Type, &config.DBCfg)
	if err != nil {
		log.Fatal("Err DB Init:", err)
	}
	db.InitDB()
	defer db.CloseDB()

	server.LoadAssets(fmt.Sprintf("assets/%s", config.ImgTheme))

	s := server.NewInstance(&config, db)
	s.Start()
}

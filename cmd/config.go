package cmd

import (
	"fmt"
	"github.com/chazari-x/ningyotsukai/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var configFile = "etc/"

type Config struct {
	Fonts config.Fonts `yaml:"fonts"`
	Log   config.Log   `yaml:"log"`
}

func getConfig(cmd *cobra.Command) *Config {
	var cfg Config

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		PadLevelText:              true,
		EnvironmentOverrideColors: true,
	})

	file, err := cmd.Flags().GetString("config")
	if err != nil {
		log.Fatalf("get flag err: %s", err)
	} else if file != "" {
		file += "."
	}

	configFile += fmt.Sprintf("config.%syaml", file)

	f, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("open config file \"%s\": %s", configFile, err)
	}

	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		log.Fatalf("decode config file: %s", err)
	}

	level, err := log.ParseLevel(cfg.Log.Level)
	if err != nil {
		log.Fatalf("parse level err: %s", err)
	}

	if cfg.Log.Level == "" {
		cfg.Log.Level = "trace"
	}
	log.SetLevel(level)

	return &cfg
}

func PersistentConfigFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("config", "", "dev")
}

package config

import (
	"fmt"
	"log"
	"os"

	configs "github.com/dinhcanh303/mail-server/pkg/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		configs.App     `yaml:"app"`
		configs.HTTP    `yaml:"http"`
		configs.Log     `yaml:"logger"`
		configs.Request `yaml:"request"`
		GRPC            `yaml:"grpc"`
	}
	GRPC struct {
		MailHost string `env-required:"true" yaml:"mail_host" env:"GRPC_MAIL_HOST"`
		MailPort int    `env-required:"true" yaml:"mail_port" env:"GRPC_MAIL_PORT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// debug
	fmt.Println("config path: " + dir)
	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

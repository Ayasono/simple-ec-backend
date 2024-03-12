package utils

import (
  "log"

  "gopkg.in/ini.v1"
)

type JwtConfig struct {
  Token string
}

func LoadEnvVariables() JwtConfig {

  cfg, err := ini.Load("conf/app.ini")

  if err != nil {
    log.Fatalf("Fail to read file: %v", err)
  }

  JwtConfig := JwtConfig{
    Token: cfg.Section("JWT").Key("Token").String(),
  }

  return JwtConfig
}

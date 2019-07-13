package config

import "os"

var (
	// Port : Porta da aplicação
	Port string
)

// LoadEnvVars : Carregar variaveis de ambiente
func LoadEnvVars() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}
}

package db

import (
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
	"log"
	"os"
)

var Supabase *supabase.Client

func InitSupabase() {
	// Pega as variáveis de ambiente
	err := godotenv.Load()

	if err != nil {
		log.Println("Arquivo .env não encontrado, carregando do ambiente")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL ou SUPABASE_KEY não estão definidos no .env")
	}

	// Inicializa o cliente do Supabase
	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		log.Fatalf("Erro ao conectar ao Supabase: %v", err)
	}

	Supabase = client
}

// Função para acessar o Supabase em outras partes do código
func GetSupabase() *supabase.Client {
	return Supabase
}

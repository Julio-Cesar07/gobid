package main

import (
	"fmt"
	"os"

	"github.com/Julio-Cesar07/gobid/cmd/api"
	"github.com/Julio-Cesar07/gobid/cmd/sqlc"
	terndotenv "github.com/Julio-Cesar07/gobid/cmd/tern-dotenv"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Inicia a api",
		Run: func(cmd *cobra.Command, args []string) {
			api.Main()
		},
	}

	ternCmd := &cobra.Command{
		Use:   "tern",
		Short: "Rodar as migrações do Tern",
		Run: func(cmd *cobra.Command, args []string) {
			terndotenv.Main()
		},
	}

	sqlcCmd := &cobra.Command{
		Use:   "sqlc",
		Short: "Gerar código Go a partir de queries SQL com sqlc",
		Run: func(cmd *cobra.Command, args []string) {
			sqlc.Main()
		},
	}

	rootCmd.AddCommand(ternCmd)
	rootCmd.AddCommand(sqlcCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

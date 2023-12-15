package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/ttacon/chalk"
)

func New(env string) {

	// Get env argument
	if env == "dev" {
		fmt.Println(chalk.White, "🚨", "Loading .env.local")
		err := godotenv.Load(".env.local")
		if err != nil {
			fmt.Println(chalk.Red, "❌", err)

		}
		fmt.Println(chalk.Green, "✅ .env.local loaded")
	}

	if env == "prod" {
		fmt.Println(chalk.White, "🚨", "Loading .env")
		err := godotenv.Load()
		if err != nil {
			fmt.Println(chalk.Red, "❌", err)
			fmt.Println(chalk.Red, "Please create .env file")

		}
		fmt.Println(chalk.Green, "✅ .env loaded")
	}

}

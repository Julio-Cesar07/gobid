package sqlc

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func Main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"sqlc",
		"generate",
		"-f",
		"./internal/store/pgstore/sqlc.yaml",
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Command execution failed: ", err)
		fmt.Println("Output: ", string(output))
		panic(err)
	}

	fmt.Println("Command execution successfully: ", string(output))
}

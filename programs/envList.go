package main

import (
	"os"
	"errors"
	"strings"
	"fmt"
)

func getEnvVars(names []string) (map[string]string, error) {
	result := map[string]string{}
	var missingVariables []string

	for _, name := range names {
		envValue, found := os.LookupEnv(name)

		if found == true {
			result[name] = envValue
		} else {
			missingVariables = append(missingVariables, envValue)
		}
	}

	if len(missingVariables) > 0 {
		message := "Missing environment variables: " + strings.Join(missingVariables, " ")
		err := errors.New(message)
		return result, err
	} else {
		return result, nil
	}

}

func main() {
	envNames := []string{"USER", "HOME", "SHELL"}

	envValues, err := getEnvVars(envNames)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for key, value := range envValues {
		fmt.Println(key, ":", value)
	}
}

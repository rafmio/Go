package cmd

import (
	"fmt"
	"os"
)

func setEnv() {
	os.Setenv("LTR_DB_NAMES", "cute_ganymede,black_oxygenium")
	os.Setenv("LTR_DB_HOST", "194.58.102.129")
	os.Setenv("LTR_DB_USER", "raf")
	os.Setenv("LTR_DB_PASSWORD", "qwq121")
	os.Setenv("LTR_DB_PORT", "5432")
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("LTR_LISTEN_PORT", "8080")
	os.Setenv("CACHE_DURATION", "5m")
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("CACHE_REDIS_ADDR", "localhost:6379")
	os.Setenv("CACHE_REDIS_PASSWORD", "")
	os.Setenv("CACHE_REDIS_DB", "0")
}

func getEnv() {
	fmt.Println("LTR_DB_NAMES:", os.Getenv("LTR_DB_NAMES"))
	fmt.Println("LTR_DB_HOST:", os.Getenv("LTR_DB_HOST"))
	fmt.Println("LTR_DB_USER:", os.Getenv("LTR_DB_USER"))
	fmt.Println("LTR_DB_PASSWORD:", os.Getenv("LTR_DB_PASSWORD"))
	fmt.Println("LTR_DB_PORT:", os.Getenv("LTR_DB_PORT"))
	fmt.Println("LOG_LEVEL:", os.Getenv("LOG_LEVEL"))
	fmt.Println("LTR_LISTEN_PORT:", os.Getenv("LTR_LISTEN_PORT"))
	fmt.Println("CACHE_DURATION:", os.Getenv("CACHE_DURATION"))
	fmt.Println("CACHE_TYPE:", os.Getenv("CACHE_TYPE"))
	fmt.Println("SHELL:", os.Getenv("SHELL"))
	fmt.Println("---------------------")
	fmt.Println("LTR_COOL_VARIABLE:", os.Getenv("LTR_COOL_VARIABLE"))

	fmt.Println("---------------------")

	homeEnv := os.Getenv("SHELL")
	arbrEnv := os.Getenv("ARBITRARY_ENV")

	if homeEnv == "" {
		fmt.Println("'HOME' environment is empty")
	} else {
		fmt.Println("homeEnv:", homeEnv)
	}

	if arbrEnv == "" {
		fmt.Println("'ARBITRARY_ENV' is empty")
	} else {
		fmt.Println("arbrEnv:", arbrEnv)
	}


		
}

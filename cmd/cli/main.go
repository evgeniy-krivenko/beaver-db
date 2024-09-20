package main

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func main() {
	logOpts := &tint.Options{Level: slog.LevelDebug}

	logger := slog.New(tint.NewHandler(os.Stdout, logOpts))

	db, err := initDatabase(logger)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("[beaver-cli] > ")
		query, err := reader.ReadString('\n')
		if err != nil {
			logger.Error("failed to read query", slog.Any("error", err))
		}
		result := db.ParseQuery(ctx, query)
		fmt.Println(result)
	}
}

package main

import (
	"fmt"
	"log/slog"

	"beaver/internal/database"
	"beaver/internal/database/compute"
	"beaver/internal/database/storage"
	"beaver/internal/database/storage/engine"
)

func initDatabase(logger *slog.Logger) (*database.Database, error) {
	inMemoryEnginge := engine.NewInMemory()

	computeParser := compute.New(logger)

	storage, err := storage.New(storage.NewOptions(inMemoryEnginge, logger))
	if err != nil {
		return nil, fmt.Errorf("create storage: %v", err)
	}

	db, err := database.New(database.NewOptions(computeParser, storage, logger))
	if err != nil {
		return nil, fmt.Errorf("create database: %v", err)
	}

	return db, nil
}

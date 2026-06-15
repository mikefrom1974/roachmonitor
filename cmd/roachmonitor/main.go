package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mikefrom1974/roachmonitor/internal/cockroach"
	"github.com/mikefrom1974/roachmonitor/internal/snapshot"
)

func main() {
	dbUser := os.Getenv("ROACH_DB_USER")
	if dbUser == "" {
		fmt.Println("ROACH_DB_USER env var not found")
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg := cockroach.Config{
		Host:     "localhost",
		Port:     26257,
		User:     dbUser,
		Database: "defaultdb",
	}

	db, err := cockroach.Connect(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	snap, err := snapshot.CollectSnapshot(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to DB. \n Version: %s\n Latency: %v\n", snap.ClusterVersion, snap.ProbeLatency)
}

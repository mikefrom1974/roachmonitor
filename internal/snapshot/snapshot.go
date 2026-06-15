package snapshot

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snapshot struct {
	ClusterVersion string
	timeStamp      time.Time
	ProbeLatency   time.Duration
}

func CollectSnapshot(ctx context.Context, db *pgxpool.Pool) (Snapshot, error) {
	var snap Snapshot

	// version
	if err := db.QueryRow(ctx, "SELECT version();").Scan(&snap.ClusterVersion); err != nil {
		return Snapshot{}, err
	}

	// latency
	lat, err := probeLatency(ctx, db)
	if err != nil {
		return Snapshot{}, err
	}
	snap.ProbeLatency = lat

	return snap, nil
}

func probeLatency(ctx context.Context, db *pgxpool.Pool) (time.Duration, error) {
	start := time.Now()

	var one int
	err := db.QueryRow(ctx, "SELECT 1;").Scan(&one)
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}

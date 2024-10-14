package distlock

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"fmt"

	_ "github.com/lib/pq"
)

// PGDistributedLock implements DistributedLock using PostgreSQL advisory locks
type PGDistributedLock struct {
	db *sql.DB
}

// NewPGDistributedLock creates a new instance of PGDistributedLock
func NewPGDistributedLock(db *sql.DB) *PGDistributedLock {
	return &PGDistributedLock{
		db: db,
	}
}

// hashKey hashes the key string into a 64-bit integer for advisory locking
func hashKey(key string) int64 {
	hash := sha256.Sum256([]byte(key))
	return int64(binary.BigEndian.Uint64(hash[:8]))
}

// Acquire tries to acquire the PostgreSQL advisory lock for the given key.
func (l *PGDistributedLock) Acquire(ctx context.Context, key string) error {
	lockKey := hashKey(key)
	query := "SELECT pg_advisory_lock($1)"
	_, err := l.db.ExecContext(ctx, query, lockKey)
	if err != nil {
		return fmt.Errorf("failed to acquire lock for key '%s': %w", key, err)
	}
	return nil
}

// TryAcquire tries to acquire the lock for the given key without blocking.
func (l *PGDistributedLock) TryAcquire(ctx context.Context, key string) (bool, error) {
	lockKey := hashKey(key)
	query := "SELECT pg_try_advisory_lock($1)"
	var success bool
	err := l.db.QueryRowContext(ctx, query, lockKey).Scan(&success)
	if err != nil {
		return false, fmt.Errorf("failed to try acquire lock for key '%s': %w", key, err)
	}
	return success, nil
}

// Release releases the PostgreSQL advisory lock for the given key.
func (l *PGDistributedLock) Release(ctx context.Context, key string) error {
	lockKey := hashKey(key)
	query := "SELECT pg_advisory_unlock($1)"
	_, err := l.db.ExecContext(ctx, query, lockKey)
	if err != nil {
		return fmt.Errorf("failed to release lock for key '%s': %w", key, err)
	}
	return nil
}

// IsHeld returns whether the lock for the given key is held by this instance.
// Note: PostgreSQL advisory locks do not provide a direct way to check if the lock is held,
// so this method may have limited utility without additional state management.
func (l *PGDistributedLock) IsHeld(key string) bool {
	// This function would require additional tracking within your app.
	return false
}

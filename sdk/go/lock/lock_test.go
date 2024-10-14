package lock

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	pglock "github.com/singulatron/singulatron/sdk/go/lock/pg"

	"github.com/stretchr/testify/require"
)

func TestLocks(t *testing.T) {
	lockStores := map[string]func(instance any) DistributedLock{
		"pgLock": func(instance any) DistributedLock {
			// Use the same PostgreSQL connection string as in your existing tests
			db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable")
			require.NoError(t, err)
			lockService := pglock.NewPGDistributedLock(db)
			return lockService
		},
	}

	tests := map[string]func(t *testing.T, lock DistributedLock){
		"AcquireRelease": LockAcquireRelease,
		"TryAcquire":     LockTryAcquire,
		"LockContention": LockContention,
	}

	for testName, test := range tests {
		for storeName, storeFunc := range lockStores {
			t.Run(fmt.Sprintf("%v %v", storeName, testName), func(t *testing.T) {
				lock := storeFunc(nil)
				test(t, lock)
			})
		}
	}
}

// LockAcquireRelease tests acquiring and releasing a lock
func LockAcquireRelease(t *testing.T, lock DistributedLock) {
	ctx := context.Background()
	key := "test_lock_acquire_release"

	// Acquire the lock
	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	// Release the lock
	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	// Ensure IsHeld returns false after release
	held := lock.IsHeld(key)
	require.False(t, held, "lock should not be held after release")
}

// LockTryAcquire tests non-blocking lock acquisition
func LockTryAcquire(t *testing.T, lock DistributedLock) {
	ctx := context.Background()
	key := "test_lock_try_acquire"

	// Try acquiring the lock, should succeed
	success, err := lock.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.True(t, success, "should acquire the lock successfully")

	// Try acquiring the lock again, should fail
	success, err = lock.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.False(t, success, "should not acquire the lock a second time")

	// Release the lock
	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")
}

// TestLockContention tests contention between two "clients" trying to acquire the same lock
func LockContention(t *testing.T, lock DistributedLock) {
	ctx := context.Background()
	key := "test_lock_contention"

	// First acquire the lock
	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	// Spin up a goroutine to try acquiring the same lock (this should block)
	acquireResult := make(chan error, 1)
	go func() {
		err := lock.Acquire(ctx, key)
		acquireResult <- err
	}()

	// Ensure that the goroutine is blocked (it shouldn't be able to acquire the lock)
	select {
	case <-acquireResult:
		t.Fatal("lock acquisition should be blocked by the first acquisition")
	case <-time.After(100 * time.Millisecond):
		// No response, as expected
	}

	// Release the lock and allow the goroutine to acquire it
	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	// The goroutine should now be able to acquire the lock
	err = <-acquireResult
	require.NoError(t, err, "second acquisition should succeed after release")
}

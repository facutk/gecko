package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/benbjohnson/litestream"
	lss3 "github.com/benbjohnson/litestream/s3"
	"github.com/facutk/golaburo/models"
)

func Litestream(ctx context.Context, dsn string, bucket string) (*litestream.DB, error) {
	// Create a Litestream DB and attached replica to manage background replication.
	log.Println("litestream: opening", dsn)
	lsdb, err := replicate(ctx, dsn, bucket)

	models.ConnectDB(dsn)

	return lsdb, err
}

func replicate(ctx context.Context, dsn, bucket string) (*litestream.DB, error) {
	// Create Litestream DB reference for managing replication.
	lsdb := litestream.NewDB(dsn)

	// Build S3 replica and attach to database.
	client := lss3.NewReplicaClient()
	client.Bucket = bucket

	replica := litestream.NewReplica(lsdb, "s3")
	replica.Client = client

	lsdb.Replicas = append(lsdb.Replicas, replica)

	if err := restore(ctx, replica); err != nil {
		return nil, err
	}

	// Initialize database.
	if err := lsdb.Open(); err != nil {
		return nil, err
	}

	return lsdb, nil
}

func restore(ctx context.Context, replica *litestream.Replica) (err error) {
	// Skip restore if local database already exists.
	if _, err := os.Stat(replica.DB().Path()); err == nil {
		log.Println("litestream: local database already exists, skipping restore")
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	// Configure restore to write out to DSN path.
	opt := litestream.NewRestoreOptions()
	opt.OutputPath = replica.DB().Path()
	opt.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)

	// Determine the latest generation to restore from.
	if opt.Generation, _, err = replica.CalcRestoreTarget(ctx, opt); err != nil {
		return err
	}

	// Only restore if there is a generation available on the replica.
	// Otherwise we'll let the application create a new database.
	if opt.Generation == "" {
		log.Println("litestream: no generation found, creating new database")
		return nil
	}

	log.Printf("litestream: restoring replica for generation %s\n", opt.Generation)
	if err := replica.Restore(ctx, opt); err != nil {
		return err
	}
	fmt.Println("litestream: restore complete")
	return nil
}

// Copyright (C) 2017 ScyllaDB

// +build all integration

package migrate

import (
	"context"
	"testing"

	"github.com/scylladb/go-log"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
	. "github.com/scylladb/scylla-manager/pkg/testutils"
	"github.com/scylladb/scylla-manager/schema"
)

func TestCreateDefaultHealthCheckRestTaskForClusterAfter018Integration(t *testing.T) {
	saveRegister()
	defer restoreRegister()
	session := CreateSessionWithoutMigration(t)

	cb := findCallback("018-empty.cql", migrate.AfterMigration)
	registerCallback("018-empty.cql", migrate.AfterMigration, func(ctx context.Context, session gocqlx.Session, logger log.Logger) error {
		Print("Given: clusters")
		const insertClusterCql = `INSERT INTO cluster (id) VALUES (uuid())`
		ExecStmt(t, session, insertClusterCql)
		ExecStmt(t, session, insertClusterCql)

		Print("When: migrate")
		if err := cb(ctx, session, logger); err != nil {
			t.Fatal(err)
		}

		Print("Then: tasks are created")
		const countSchedulerTask = `SELECT COUNT(*) FROM scheduler_task`
		q := session.Query(countSchedulerTask, nil)
		defer q.Release()

		var count int
		if err := q.Scan(&count); err != nil {
			t.Fatal(err)
		}
		if count != 2 {
			t.Fatal()
		}

		return nil
	})

	if err := migrate.FromFS(context.Background(), session, schema.Files); err != nil {
		t.Fatal("migrate:", err)
	}
}

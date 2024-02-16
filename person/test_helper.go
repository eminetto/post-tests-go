package person

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MysqlDBContainer struct {
	testcontainers.Container
	URI string
}

const (
	dbUser         = "workshop"
	dbPassword     = "workshop"
	database       = "workshop"
	dbRootPassword = "db-root-password"
)

func SetupMysqL(t testing.TB) *MysqlDBContainer {
	t.Helper()
	ctx := context.TODO()
	req := testcontainers.ContainerRequest{
		Image:        "mariadb:11.3.1-rc-jammy",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("Version: '11.3.1-MariaDB-1:11.3.1+maria~ubu2204'  socket: '/run/mysqld/mysqld.sock'  port: 3306  mariadb.org binary distribution"),
		Env: map[string]string{
			"MARIADB_USER":          dbUser,
			"MARIADB_PASSWORD":      dbPassword,
			"MARIADB_ROOT_PASSWORD": dbRootPassword,
			"MARIADB_DATABASE":      database,
		},
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Errorf("error creating container %s", err.Error())
	}
	mappedPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		t.Errorf("error getting container port %s", err.Error())
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		t.Errorf("error getting container host address %s", err.Error())
	}
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", dbRootPassword, hostIP, mappedPort.Port(), database)
	t.Cleanup(func() {
		container.Terminate(ctx)
	})

	return &MysqlDBContainer{Container: container, URI: uri}
}

func InitMySQL(t testing.TB, db *sql.DB) {
	t.Helper()
	ctx := context.TODO()
	query := []string{
		fmt.Sprintf("use %s;", database),
		"create table if not exists person (id int AUTO_INCREMENT,first_name varchar(100), last_name varchar(100), created_at datetime, updated_at datetime, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;",
	}
	for _, q := range query {
		_, err := db.ExecContext(ctx, q)
		if err != nil {
			t.Errorf("error executing create query %s", err.Error())
		}
	}
}

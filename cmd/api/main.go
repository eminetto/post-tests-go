package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/eminetto/post-tests-go/internal/api"
	"github.com/eminetto/post-tests-go/internal/http/echo"
	"github.com/eminetto/post-tests-go/person"
	"github.com/eminetto/post-tests-go/person/mysql"
	"github.com/eminetto/post-tests-go/weather"
	_ "github.com/go-sql-driver/mysql"
)

const ( //@todo pegar essa informação de variáveis de ambiente
	dbUser         = "workshop"
	dbPassword     = "workshop"
	database       = "workshop"
	dbRootPassword = "db-root-password"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, "localhost", "3306", database)
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		logger.Error("error opening database", err)
	}
	repo := mysql.NewMySQL(db)
	pService := person.NewService(repo)

	wService := weather.NewService(os.Getenv("API_KEY"))

	h := echo.Handlers(logger, pService, wService)
	err = api.Start(logger, "8000", h)
	if err != nil {
		logger.Error("error running api", err)
	}
}

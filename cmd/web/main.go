package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/nmdra/snipbox/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	Logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	var port = flag.String("port", "4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@tcp(db:3306)/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		Logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   formDecoder,
		sessionManager: sessionManager,
	}

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	envPort := os.Getenv("PORT") // for docker
	if envPort != "" {
		*port = envPort
		app.Logger.Info("Using environment variable PORT:" + *port)
	} else {
		app.Logger.Warn("Environment variable PORT is not set, using default or flag value")
	}

	app.Logger.Info("Starting server", slog.Any("PORT", *port))

	err = http.ListenAndServe(":"+*port, app.routes())
	app.Logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

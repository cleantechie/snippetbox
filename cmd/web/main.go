package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.harshasv.net/internal/datamodels"
)

type application struct {
	infoLogger   *log.Logger
	errLogger    *log.Logger
	snippetModel *datamodels.SnippetModel
}

func main() {
	host := flag.String("host", ":5000", "Http Port")
	dsn := flag.String("dsn", "web:1234@/snippetbox?parseTime=true", "MySQL ")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

	database, error := openDatabase(*dsn)

	if error != nil {
		errLog.Fatal(error)
	}
	infoLog.Print("Connected to Database Succcessfully")
	defer database.Close()

	// creating an combined object of both these loggers so then can be injected
	app := &application{
		// Info log instance
		infoLogger: infoLog,
		// Error log instance
		errLogger: errLog,
		// snippetModel points to database instance
		snippetModel: &datamodels.SnippetModel{DB: database},
	}
	// new server struct so even http error will use the err log instance
	server := &http.Server{
		Addr:     *host,
		Handler:  app.routes(),
		ErrorLog: app.errLogger,
	}
	app.infoLogger.Printf("Starting server on port %s", *host)
	serverStartErr := server.ListenAndServe()
	app.errLogger.Fatal(serverStartErr)
}

func openDatabase(connectionString string) (*sql.DB, error) {
	// the "mysql" refers to the driver that is blank identifier at the imports without directly referencing the driver's code
	database, error := sql.Open("mysql", connectionString)
	if error != nil {
		return nil, error
	}

	if error = database.Ping(); error != nil {
		return nil, error
	}
	return database, nil
}

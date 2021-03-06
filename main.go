package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/satit13/bcsync/app"
	"github.com/satit13/bcsync/mssql"
	"log"
	"net/http"
	"github.com/satit13/bcsync/auth"
	"database/sql"
)

const (
	sslMode = "disable"
	dbPort = "1433"
	dbUser = "sa"
	dbHost = "nebula"
	dbPass = "[ibdkifu"
	dbName = "bcnp"
)

func main() {
	log.Println("Sync Invoice Service")
	//conn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", dbHost, dbUser, dbPass, dbPort, dbName)
	//conn := "sqlserver://sa:[ibdkifu@nebula?database=bcnp"
	conn := "server=nebula;user id=sa;password=[ibdkifu;database=bcnp;app name=MyAppName;encrypt=disable"
	fmt.Println(conn)
	db, err := sql.Open("mssql", conn)
	//db := sqlx.MustConnect("mssql", conn)


	defer db.Close()

	//create repositories
	authRepo, err := mssql.NewAuthRepository(db)
	must(err)

	//create services
	authService, err := auth.NewService(authRepo)
	must(err)

	appRepo, err := mssql.NewAppeRepository(db)
	must(err)

	appService, err := app.NewService(appRepo)
	must(err)

	// mount handler
	fmt.Println("begin handling")
	mux := http.NewServeMux()
	mux.HandleFunc("/version", apiVersionHandler)

	mux.Handle("/v1/invoice/", http.StripPrefix("/v1/invoice", app.MakeHandler(appService)))
	// create main handler
	h := auth.MakeMiddleware(authService)(mux)
	http.ListenAndServe(":8888", h)

}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}

func apiVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//t := time.Now()
	json.NewEncoder(w).Encode(struct {
		Version     string `json:"version"`
		Description string `json:"description"`
		Creator     string `json:"creator"`
		LastUpdate  string `json:"lastupdate"`
	}{
		"0.1 BETA",
		"BC Sync client service",
		"Nopadol Dev team 2016",
		"2018-04-22",
	})
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const (
	server_port = ":8080"
	db_host     = "localhost"
	db_port     = 5432
	db_user     = "db_user"
	db_password = "db_password"
	db_name     = "postgress"
)

type UserID struct {
	user_id string
}

func main() {
	if err := serverRun(); err != nil {
		fmt.Println("ERROR:%w", err)
	}
}

func serverRun() error {
	r := chi.NewRouter()
	r.Post("/users/{id}", responseTest)
	fmt.Printf("Server starting on%s\n", server_port)
	http.ListenAndServe(server_port, r)
	return fmt.Errorf("error loading server")
}

func responseTest(w http.ResponseWriter, r *http.Request) {
	var u UserID
	u.user_id = chi.URLParam(r, "id")
	if _, err := fmt.Fprintf(w, "%s", u.user_id); err != nil {
		log.Println(err)
	}

	if err := u.createUserDB(); err != nil {
		fmt.Println("ERROR:%w", err)
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserID) createUserDB() error {
	db, err := connectDB()
	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("failure ping: %w", err)
	}

	sql := `insert into "dockertable"("user_id") values($1);`
	_, err = db.Exec(sql, u.user_id)
	if err != nil {
		return fmt.Errorf("error with create user")
	}

	return nil
}

func connectDB() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", db_host, db_port, db_user, db_password, db_name)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, fmt.Errorf("problem with connect to DB")
	}

	return db, err
}

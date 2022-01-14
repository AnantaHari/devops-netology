package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func pg_conn_main() {
	db_url := "postgres://postgres:Primanka90@192.168.4.82:5432/demo2"
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// col_names_str := `SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = 'PermissionPolicyUser'`
	sql := `select * from "PermissionPolicyUser"`

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		fmt.Println(values)
	}
}

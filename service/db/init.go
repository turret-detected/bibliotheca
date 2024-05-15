package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func InitDatabase(conn *pgx.Conn) error {
	ctx := context.Background()

	// Don't create tables if DB exists
	exists := false
	out := conn.QueryRow(ctx, QueryLibraryExists)
	err := out.Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// Create tables
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, strings.TrimSpace(CreateFileTable))
	fmt.Println(result.String())
	if err != nil {
		return err
	}

	result, err = tx.Exec(ctx, strings.TrimSpace(CreateTagTable))
	fmt.Println(result.String())
	if err != nil {
		return err
	}

	result, err = tx.Exec(ctx, strings.TrimSpace(CreateLibraryTable))
	fmt.Println(result.String())
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil

}

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var MIGRATIONS_DIR = "./database/migrate/migrations"

func main() {
	ctx := context.Background()
	godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer conn.Close(ctx)

	fresh := flag.Bool("fresh", false, "Drop all tables and rerun all migrations")
	flag.Parse()

	if *fresh {
		if err := WipeDB(ctx, conn); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Running migrations...")

	migrationsRun := GetRunMigrations(ctx, conn)
	migrationFileNames := GetMigrationFileNames()

	var numRan int
	var numSkipped int
	for _, migrationFileName := range migrationFileNames {
		if ran := migrationsRun[migrationFileName]; ran {
			fmt.Printf("%s: already ran...\n", migrationFileName)
			numSkipped++
		} else {
			if err := RunMigration(ctx, migrationFileName, conn); err != nil {
				fmt.Printf("error: %s\n", err)
			}
			numRan++
		}
	}
	fmt.Printf("\n%d migrations ran\n", numRan)
	fmt.Printf("%d migrations skipped\n", numSkipped)
}

func GetRunMigrations(ctx context.Context, conn *pgx.Conn) map[string]bool {
	rows, err := conn.Query(ctx, "select name from migrations")
	if err != nil {
		if err.Error() != "ERROR: relation \"migrations\" does not exist (SQLSTATE 42P01)" {
			fmt.Println("error fetching migrations")
		}
	}

	files_run := map[string]bool{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println("error reading migration table")
			break
		}
		files_run[name] = true
	}
	return files_run
}

func GetMigrationFileNames() []string {
	entries, err := os.ReadDir(MIGRATIONS_DIR)
	if err != nil {
		fmt.Printf("error reading entries: %s\n", err)
	}

	fileNames := []string{}
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasPrefix(name, "_") {
			fileNames = append(fileNames, entry.Name())
		}
	}

	return fileNames
}

func RunMigration(ctx context.Context, fileName string, conn *pgx.Conn) error {
	path := fmt.Sprintf("%s/%s", MIGRATIONS_DIR, fileName)
	migration, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, string(migration))
	if err != nil {
		return err
	}

	addSql := `
		insert into migrations(name, created_on) values(@name, @created_on)
	`
	args := pgx.NamedArgs{
		"name":       fileName,
		"created_on": time.Now(),
	}
	_, err = conn.Exec(ctx, addSql, args)
	if err != nil {
		return err
	}
	fmt.Printf("%s: successfully executed\n", fileName)
	return nil
}

func WipeDB(ctx context.Context, conn *pgx.Conn) error {
	path := fmt.Sprintf("%s/_wipe.sql", MIGRATIONS_DIR)
	wipe, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, string(wipe))
	if err != nil {
		return err
	}

	fmt.Println("Successfully wiped database...")
	return nil
}

package migrate

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type Migrator interface {
	WipeDB() error
}

var MigrationsDir = "./migrate/migrations"

func MigrateDB(ctx context.Context, getenv func(string) string, fresh bool) {
	dbUrl := getenv("DATABASE_URL")

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer conn.Close(ctx)

	if fresh {
		if err := wipeDB(ctx, conn); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Running migrations...")

	migrationsRun := getRunMigrations(ctx, conn)
	migrationFileNames := getMigrationFileNames()

	var numRan int
	var numSkipped int
	for _, migrationFileName := range migrationFileNames {
		if ran := migrationsRun[migrationFileName]; ran {
			fmt.Printf("%s: already ran...\n", migrationFileName)
			numSkipped++
		} else {
			if err := runMigration(ctx, migrationFileName, conn); err != nil {
				fmt.Printf("error: %s\n", err)
			}
			numRan++
		}
	}
	fmt.Printf("\n%d migrations ran\n", numRan)
	fmt.Printf("%d migrations skipped\n", numSkipped)
}

func wipeDB(ctx context.Context, conn *pgx.Conn) error {
	path := fmt.Sprintf("%s/_wipe.sql", MigrationsDir)
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
func getRunMigrations(ctx context.Context, conn *pgx.Conn) map[string]bool {
	rows, err := conn.Query(ctx, "select name from migrations")
	if err != nil {
		if err.Error() != "ERROR: relation \"migrations\" does not exist (SQLSTATE 42P01)" {
			fmt.Println("error fetching migrations")
		}
	}

	filesRun := map[string]bool{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println("error reading migration table")
			break
		}
		filesRun[name] = true
	}
	return filesRun
}

func getMigrationFileNames() []string {
	entries, err := os.ReadDir(MigrationsDir)
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

func runMigration(ctx context.Context, fileName string, conn *pgx.Conn) error {
	path := fmt.Sprintf("%s/%s", MigrationsDir, fileName)
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

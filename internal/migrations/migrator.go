package migrations

import (
	"database/sql"
	"io/ioutil"
	"strings"
)

func Migrate(db *sql.DB) error {
	sqlFile := "../internal/migrations/0001_create_author_table.sql" // Укажите путь к вашему файлу
	sqlBytes, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	// Выполнение SQL-скрипта
	if err := executeSQLScript(db, string(sqlBytes)); err != nil {
		return err
	}

	sqlFile = "../internal/migrations/0002_create_quote_table.sql" // Укажите путь к вашему файлу
	sqlBytes, err = ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	// Выполнение SQL-скрипта
	if err := executeSQLScript(db, string(sqlBytes)); err != nil {
		return err
	}

	return nil
}

func executeSQLScript(db *sql.DB, script string) error {
	// Разделяем скрипт на отдельные команды по точке с запятой
	commands := strings.Split(script, ";")

	for _, cmd := range commands {
		// Удаляем лишние пробелы и пропускаем пустые команды
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		// Выполняем каждую команду
		if _, err := db.Exec(cmd); err != nil {
			return err
		}
	}

	return nil
}

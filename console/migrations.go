package console

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	// Postgres driver import from migrate
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	"github.com/spf13/viper"

	"github.com/contacts_api_go/logger"
	"github.com/contacts_api_go/appcontext"
)

const (
	MIGRATION_SOURCE_PATH = "file://migrations"
	MIGRATION_FILE_PATH   = "./migrations"
)

func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return errors.New("Migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", MIGRATION_FILE_PATH, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", MIGRATION_FILE_PATH, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}
	logger.Log.Info("created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	logger.Log.Info("created %s\n", downMigrationFilePath)

	return nil
}

func RunDatabaseMigrations() error {
	m, error := getMigrate()

	if error != nil {
		return joinErrors(error)
	}
	error = m.Up()
	if error != nil {
		logger.Log.Error("error", error)
		return joinErrors(error)
	}

	logger.Log.Info("Migrations successful")
	return nil
}

func RollbackLatestMigration() error {
	m, error := getMigrate()
	if error != nil {
		return joinErrors(error)
	}
	stepError := m.Steps(-1)
	if stepError != nil {
		logger.Log.Info("Rollback error", stepError)
		return joinErrors(stepError)
	}
	logger.Log.Info("Migrations Rollback successful")
	return nil
}

func joinErrors(inputErrors error) error {
	var errorMsgs []string
	errorMsgs = append(errorMsgs, inputErrors.Error())
	errMsgJoined := strings.Join(errorMsgs, ",")
	return fmt.Errorf(errMsgJoined)
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func getMigrate() (*migrate.Migrate, error) {
	driver, _ := appcontext.GetDriver()
	return migrate.NewWithDatabaseInstance(
		MIGRATION_SOURCE_PATH,
		viper.GetString("DB_NAME"), driver)
}

package database

import (
	"reflect"

	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
)

// Migration interface.
type Migration interface {
	Migrate(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

// Migrator engine to work with migrations.
type Migrator struct {
	migrations []Migration
}

// SetMigrations to run.
func (m *Migrator) SetMigrations(migrations ...Migration) {
	m.migrations = migrations
}

// Migrate database.
func (m *Migrator) Migrate(db *gorm.DB) error {
	return m.makeGormigrate(db).Migrate()
}

// Rollback last migration.
func (m *Migrator) Rollback(db *gorm.DB) error {
	return m.makeGormigrate(db).RollbackLast()
}

// Reset all migrations.
func (m *Migrator) Reset(db *gorm.DB) error {
	return nil
}

// Make new gormigrate instance.
func (m *Migrator) makeGormigrate(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		m.transformMigrations(),
	)
}

// Transform migrations to gormigrate
func (m *Migrator) transformMigrations() []*gormigrate.Migration {
	var gormigrations []*gormigrate.Migration

	for _, migration := range m.migrations {
		gormigrations = append(gormigrations, &gormigrate.Migration{
			ID:       m.getMigrationName(migration),
			Migrate:  migration.Migrate,
			Rollback: migration.Rollback,
		})
	}

	return gormigrations
}

// Get unigue migration ID from struct name.
func (m *Migrator) getMigrationName(migration Migration) string {
	t := reflect.TypeOf(migration)

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}

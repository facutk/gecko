package models

func GetMigrationVersion() int {
	row := db.QueryRow("select version from schema_migrations;")

	var schema_migration_version int
	row.Scan(&schema_migration_version)

	return schema_migration_version
}

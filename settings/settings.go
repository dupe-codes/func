// Global application settings go here

package settings

import (
// nothing right now
)

var (
	App      = getAppConfig()
	Security = getSecurityConfig()
	Database = getDbConfig()
)

// The App struct encapsulates overall application
// settings
type AppConfig struct {
	Name string
	Port string
}

// The SecurityConfig struct encapsulates settings
// specific to maintaining and establishing applications security
type SecurityConfig struct {
	SessionKeyLen int
}

// The DbConfig struct holds all relevant settings for the
// application database
type DbConfig struct {
	Url  string
	Name string
}

// getConfig sets all needed application settings
func getAppConfig() *AppConfig {
	config := new(AppConfig)

	config.Name = "Serve"
	config.Port = ":8080"

	return config
}

// getSecuritySettings returns a SecurityConfig struct
// set with the desired settings
func getSecurityConfig() *SecurityConfig {
	config := new(SecurityConfig)
	config.SessionKeyLen = 16

	return config
}

// Returns a DbConfig struct with the correct desired settings
func getDbConfig() *DbConfig {
	config := new(DbConfig)
	config.Url = "localhost"
	config.Name = "serve-dev"

	return config
}

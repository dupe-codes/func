// Global application settings go here
// TODO: Look in to doing this with the Flags package

package settings

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

	config.Name = "Func"
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
	config.Name = "func-dev"

	return config
}

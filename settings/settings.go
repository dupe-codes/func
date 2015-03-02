// Global application settings go here

package settings

import (
    // nothing right now
)

var (
    App = getAppConfig()
    Security = getSecurityConfig()
)

// The App struct encapsulates overall application
// settings
type App struct {
    Name     string
    Port     string
}

// The SecurityConfig struct encapsulates settings
// specific to maintaining and establishing applications security
type SecurityConfig struct {
    SessionKeyLen   int
}

// getConfig sets all needed application settings
func getAppConfig() *App {
    config := new(App)

    config.Name = "Serve"
    config.Port = ":8080"

    return config
}

// getSecuritySettings returns a SecurityConfig struct
// set with the desired settings
func getSecuritySettings *SecurityConfig {
    config := new(SecurityConfig)
    config.SessionKeyLen = 16

    return config
}

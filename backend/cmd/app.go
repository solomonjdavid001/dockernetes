package cmd

import (
	"context"
	"embed"
	"fmt"
	"os"
)

// App struct
type App struct {
	ctx        context.Context
	configFile embed.FS
}

// NewApp creates a new App application struct, accepting configFile as an argument
func NewApp(configFile embed.FS) *App {
	return &App{
		configFile: configFile,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	StartServer(a.configFile)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Shutdown(ctx context.Context) {
	os.Exit(0)
}

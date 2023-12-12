package application

import (
	"app/internal/handler"
	"app/internal/storage"

	"github.com/gin-gonic/gin"
)

// ConfigDefault represents the default application configuration
type ConfigDefault struct {
	// ServerHost is the server host
	ServerHost string
	// ServerPort is the server port
	ServerPort string
	// FilePathMovies is the file path of the movies
	FilePathMovies string
}

// NewDefault creates a new default application
func NewDefault(cfg *ConfigDefault) *Default {
	// default values
	defaultCfg := &ConfigDefault{
		ServerHost:     "localhost",
		ServerPort:     "8080",
		FilePathMovies: "movies.json",
	}
	if cfg != nil {
		if cfg.ServerHost != "" {
			defaultCfg.ServerHost = cfg.ServerHost
		}
		if cfg.ServerPort != "" {
			defaultCfg.ServerPort = cfg.ServerPort
		}
		if cfg.FilePathMovies != "" {
			defaultCfg.FilePathMovies = cfg.FilePathMovies
		}
	}

	return &Default{
		serverHost:     defaultCfg.ServerHost,
		serverPort:     defaultCfg.ServerPort,
		filePathMovies: defaultCfg.FilePathMovies,
	}
}

// Default represents the default application
type Default struct {
	// serverHost is the server host
	serverHost string
	// serverPort is the server port
	serverPort string
	// filePathMovies is the file path of the movies
	filePathMovies string
}

// Run runs the application
func (a *Default) Run() (err error) {
	// dependencies
	// - storage: json
	st := storage.NewMovieReaderJSON(a.filePathMovies)
	// - handler: movie
	hd := handler.NewMovieDefault(st)
	// - server
	sv := gin.Default()
	sv.GET("/movies", hd.Get())

	// run the server
	err = sv.Run(a.serverHost + ":" + a.serverPort)
	return
}
package world

import (
	"net/http"
	"sync"
)

// Config représente la configuration pour World.
type API struct {
	HTTPRegister func(method, path string, handler http.HandlerFunc)
}

// World représente la structure principale de la catégorie World.
type World struct {
	conf   *API
	confMu sync.RWMutex
}

type Config struct {
	// Filename is the name of the database file.
	Filename string

	// Enabled tells if the statistics are enabled.
	Enabled bool
}

// Start registers web handlers and starts world updates loop.
func (d *World) Init(){
	d.RegisterWorldHandlers()
}
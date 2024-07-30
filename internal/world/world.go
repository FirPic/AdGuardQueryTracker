package world

import (
	"net/http"
	"sync"
)

// Config représente la configuration pour World.
type Config struct {
	HTTPRegister func(method, path string, handler http.HandlerFunc)
}

// World représente la structure principale de la catégorie World.
type World struct {
	conf   *Config
	confMu sync.RWMutex
}

// NewWorld crée une nouvelle instance de World.
func NewWorld(conf *Config) *World {
	return &World{conf: conf}
}

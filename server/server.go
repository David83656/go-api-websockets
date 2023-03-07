package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port      string
	JWTSecret string
	DB_URL    string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {

	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	//Manejo de errores por falta de parte en campos requeridos
	if config.Port == "" {
		return nil, errors.New("Puerto requerido ")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("Token requerido ")
	}
	if config.DB_URL == "" {
		return nil, errors.New("Url de DB requerido ")
	}

	bk := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return bk, nil
}

func (b *Broker) StartSv(bind func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	bind(b, b.router)
	log.Println("Inicializado servidor en el puerto: ", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package products

import "github.com/jcsalgadomeza/desafio-goweb-camilo-salgado/desafio/internal/domains"

type Repository interface {
	GetAll() []domains.Ticket
}

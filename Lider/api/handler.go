package api

import (
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) ParticiparJuego(ctx context.Context, in *PeticionParticipar) (*ConfirmacionParticipacion, error) {

}

func (s *Server) Jugar(ctx context.Context, in *Jugada) (*Estado, error) {

}

func (s *Server) Monto(ctx context.Context, in *PedirMonto) (*MontoJugador, error) {

}

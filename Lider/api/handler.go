package api

import (
	"golang.org/x/net/context"
)

var jugadores int = 15
var etapa_actual int = 0
var ronda_atual int = 0
var lider_e1_r1 int = -1
var lider_e1_r2 int = -1
var lider_e1_r3 int = -1
var lider_e1_r4 int = -1
var lider_e2 int = -1
var lider_e3 int = -1

type Server struct {
}


func (s *Server) ParticiparJuego(ctx context.Context, in *PeticionParticipar) (*ConfirmacionParticipacion, error) {
	if (in.participar == "participar"){
		if (jugadores < 16){

			jugadores = jugadores+1
			return &ConfirmacionParticipacion {confirmacion: "ya esta participando en el juego del calamar"}
		}
		else{

			return &ConfirmacionParticipacion {confirmacion: "ya no puede participar, el juego esta lleno"}
		}
		
	}
	else{

		return &ConfirmacionParticipacion {confirmacion: "verifique que ha escrito bien participar"}
	}
}

func (s *Server) Estado(ctx context.Context, in *Check) (*State, error) {
	if (etapa_actual == 0){
		return &State {etapa: "Aun no comienza el juego, espere que el Lider de la señal"}
	}

}

func (s *Server) Jugar(ctx context.Context, in *Jugadas) (*Estado, error) {
	

	}
	else

}

func (s *Server) Monto(ctx context.Context, in *PedirMonto) (*MontoJugador, error) {

}

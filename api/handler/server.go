package handler

import "api/gen"

type Server struct{}

// Make sure we conform to StrictServerInterface

var _ gen.StrictServerInterface = (*Server)(nil)

func NewServer() *Server {
	return &Server{}
}

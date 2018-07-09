package messaging

type Server struct {
	Transport
	started bool
}

func NewServer(tp Transport) {
	return &Server{
		Transport: tp,
		started:   false,
	}
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Wait() {

}

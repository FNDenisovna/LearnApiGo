package grpc

import (
	"LearnApiGo/internal/grpc/handler"
	"LearnApiGo/internal/grpc/proto"
	"LearnApiGo/internal/services"
	"flag"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
	ServerSet
	handl         *handler.GrpcAlbumHandler
	albumServices services.IAlbums
}

type ServerSet struct {
	addr        *string
	logrusEntry *logrus.Entry
}

func New(service services.IAlbums) {
	logrusLogger := logrus.New()
	settings := &ServerSet{
		addr: flag.String("addr", ":8000", "HTTPS network address"),

		logrusEntry: logrus.NewEntry(logrusLogger),
	}
	server := &Server{
		srv: grpc.NewServer(
			grpc.StreamInterceptor(
				grpc_middleware.ChainStreamServer(
					grpc_logrus.StreamServerInterceptor(settings.logrusEntry),
					grpc_recovery.StreamServerInterceptor(),
				),
			),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(settings.logrusEntry),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		), //, authStream  , authUnary
		handl: handler.New(service),
	}

	server.ListenAndServe()
}

func (s *Server) ListenAndServe() error {
	listner, err := net.Listen("tcp", *s.addr)
	defer s.Close()

	if err != nil {
		return err
	}

	proto.RegisterGrpcAlbumServer(s.srv, s.handl)
	if err := s.srv.Serve(listner); err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() {
	s.srv.GracefulStop()
}

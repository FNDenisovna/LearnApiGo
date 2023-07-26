package grpc

import (
	handler "LearnApiGo/internal/grpc/handler"
	"LearnApiGo/internal/grpc/proto"
	"LearnApiGo/internal/services"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
	ServerSet
	handl         proto.GrpcAlbumServer
	albumServices services.IAlbums
}

type ServerSet struct {
	addr        string
	logrusEntry *logrus.Entry
}

var addr = ":8000"

func New(service services.IAlbums) {
	logrusLogger := &logrus.Logger{
		Out:   io.MultiWriter(os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:00:00",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
	settings := &ServerSet{
		addr:        addr,
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
		handl:     handler.New(service),
		ServerSet: *settings,
	}

	go func(s *Server) {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Grpc-service has errors on start: %v\n", err)
		}
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
		<-signalCh
		defer s.Close()
	}(server)
}

func (s *Server) ListenAndServe() error {
	fmt.Printf("Grpc-service started successfully on http port %s\n", s.addr)

	listner, err := net.Listen("tcp", s.addr)

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

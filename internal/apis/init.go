package apis

import (
	//"LearnApiGo/internal/docs"

	"LearnApiGo/internal/models"
	services "LearnApiGo/internal/services"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "LearnApiGo/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	//"github.com/go-chi/render"
)

type Api struct {
	service services.IAlbums
	//server  *http.Server
}

var addr *string

func New(service services.IAlbums) {
	addr = flag.String("addr", ":8080", "HTTPS network address")
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	flag.Parse()

	var api = &Api{
		service: service,
		//server:  srv,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: getHandler(api),
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	go func() {
		err := srv.ListenAndServeTLS(*certFile, *keyFile)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	fmt.Printf("Service started successfully on http port %s\n", *addr)

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

// @title Swagger Example API
// @version 1.0

// @host petstore.swagger.io
// @BasePath /v2
func getHandler(api *Api) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(middleware.URLFormat)
	//r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://localhost"+*addr+"/swagger/doc.json"),
	))

	// Public Routes
	r.Group(func(r chi.Router) {

		r.Route("/albums/{limit}", func(r chi.Router) {
			r.Use(LimitParamCtx)
			r.Get("/", api.GetAlbums)
		})
		r.Route("/album/{id}", func(r chi.Router) {
			r.Use(IdParamCtx)
			r.Get("/", api.GetAlbum)
		})

		r.With(api.BasicAuth).Post("/album/create", api.CreateAlbum)

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	// Private Routes
	// Require Authentication
	/*r.Group(func(r chi.Router) {

		r.Post("/secret/", func(w http.ResponseWriter, req *http.Request) {
			user, pass, ok := req.BasicAuth()
			if ok && verifyUserPass(user, pass) {
				fmt.Fprintf(w, "You get to see the secret\n")
			} else {
				w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})

		r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
			// Simulates some hard work.
			//
			// We want this handler to complete successfully during a shutdown signal,
			// so consider the work here as some background routine to fetch a long running
			// search query to find as many results as possible, but, instead we cut it short
			// and respond with what we have so far. How a shutdown is handled is entirely
			// up to the developer, as some code blocks are preemptible, and others are not.
			time.Sleep(5 * time.Second)

			w.Write([]byte(fmt.Sprintf("all done.\n")))
		})
	})*/

	return r
}

func IdParamCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var ctx context.Context
		if albumId := chi.URLParam(r, "id"); albumId != "" {
			id, err := uuid.Parse(albumId)
			if err != nil {
				render.Render(w, r, models.ErrInvalidRequest(err, models.ParseErr+"albumId", 400))
				return
			}
			//Todo Проверка наличия записи в БД с таким айди
			ctx = context.WithValue(r.Context(), "id", id)
		} else {
			render.Render(w, r, models.ErrInvalidRequest(err, models.ParseErr+"albumId", 400))
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LimitParamCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		if limit := chi.URLParam(r, "limit"); limit != "" {
			limitInt, err := strconv.Atoi(limit)
			if err != nil {
				render.Render(w, r, models.ErrInvalidRequest(err, models.ParseErr+"limit", 400))
				return
			}
			ctx = context.WithValue(r.Context(), "limit", limitInt)
		} else {
			ctx = context.WithValue(r.Context(), "limit", 0)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

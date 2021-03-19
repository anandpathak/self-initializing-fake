package server

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"self_initializing_fake/internal/server/admin"
	"self_initializing_fake/internal/server/mock"
	"self_initializing_fake/internal/service"
	"self_initializing_fake/pkg/memorydb"
	"time"
)

var (
	g errgroup.Group

)
func Start() {

	schema := memorydb.CreateSchema("mock_request","ID", "id" )
	db, err := memorydb.New(&schema)
	if err != nil {
		panic(err)
	}
	configurationService := service.Configure{DB: db}

	adminServer := &http.Server{
		Addr:              ":8112",
		Handler:           admin.AdminRoutes(configurationService),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	mockServer := &http.Server{
		Addr:              ":8113",
		Handler:           mock.MockRoutes(),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	g.Go(func() error {
		err := adminServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := mockServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}


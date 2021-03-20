package server

import (
	"log"
	"net/http"
	"self_initializing_fake/internal/server/setup"
	"self_initializing_fake/internal/server/fake"
	"self_initializing_fake/pkg/memorydb"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func Start() {

	schema := memorydb.CreateSchema("mock_request", "ID", "id")
	db, err := memorydb.New(&schema)
	if err != nil {
		panic(err)
	}

	configurationService := setup.Service{DB: db}
	mockService := fake.Mock{DB: db}

	setupServer := &http.Server{
		Addr:              ":8112",
		Handler:           setup.Routes(configurationService),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	mockServer := &http.Server{
		Addr:              ":8113",
		Handler:           fake.Routes(mockService),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	g.Go(func() error {
		err := setupServer.ListenAndServe()
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

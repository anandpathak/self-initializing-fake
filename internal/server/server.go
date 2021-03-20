package server

import (
	"fmt"
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
	TableName = "test_double"
)

func Start(setupServerPort, fakeServerPort string) {

	schema := memorydb.CreateSchema(TableName, "ID", "id")
	db, err := memorydb.New(&schema,TableName )
	if err != nil {
		panic(err)
	}

	configurationService := setup.Service{DB: db}
	mockService := fake.Mock{DB: db}

	port := fmt.Sprintf(":%s",setupServerPort)
	setupServer := &http.Server{
		Addr:              port,
		Handler:           setup.Routes(configurationService),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	port = fmt.Sprintf(":%s",fakeServerPort)
	mockServer := &http.Server{
		Addr:              port,
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

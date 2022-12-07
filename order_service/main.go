package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"mymachine707/config"
	blogpost "mymachine707/protogen/blogpost"
	"mymachine707/services/article"
	"mymachine707/services/author"
	"mymachine707/storage"
	"mymachine707/storage/postgres"

	_ "github.com/lib/pq"
	"github.com/swaggo/swag/example/basic/docs"
)

func initGRPC(cfg config.Config, stg storage.Interfaces) {

}

// @license.name	Apache 2.0
func main() {
	cfg := config.Load()

	psqlConfigString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	var err error
	var stg storage.Interfaces

	stg, err = postgres.InitDB(psqlConfigString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("gRPC server tutorial in Go in GRPCPort: %s", cfg.GRPCPort)

	listener, err := net.Listen("tcp", cfg.GRPCPort)

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	AuthorService := author.NewAuthorService(stg)
	blogpost.RegisterAuthorServiceServer(s, AuthorService)

	ArticleService := article.NewArticleService(stg)
	blogpost.RegisterArticleServiceServer(s, ArticleService)

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

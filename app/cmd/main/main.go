package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/koutaro0205/backend-template/app/internal/graph"
	"github.com/koutaro0205/backend-template/app/internal/graph/generated"
	"github.com/koutaro0205/backend-template/app/internal/middleware"
	db "github.com/koutaro0205/backend-template/app/internal/repository"
)

const defaultPort = "8000"
const endpoint = "graphql"

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/"+endpoint)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func initServer() *gin.Engine {
	server := gin.Default()
	middleware.Setup(server)

	server.POST("/"+endpoint, graphqlHandler())
	server.GET("/", playgroundHandler())

	return server
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()
	defer db.CloseDB()

	server := initServer()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(server.Run(":" + port))
}

package main

import (
	"net/http"
	"context"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/GibJob-ai/GObjob/utils"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/schema"
)

func main() {
	util.LoadEnv()

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	context.Background()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/query", handler.Authenticate(&handler.GraphQL{Schema: schema}))
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend/public", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Start and run the server
	router.Run(":3000")
}

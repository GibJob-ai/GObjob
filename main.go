package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	graphql "github.com/graph-gophers/graphql-go"
	// "github.com/GibJob-ai/GObjob/utils"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/handler"
	"github.com/GibJob-ai/GObjob/resolvers"
	"github.com/GibJob-ai/GObjob/schema"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	context.Background()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

	mux := http.NewServeMux()
	mux.Handle("/graphql", handler.GraphiQL{})
	mux.Handle("/query", handler.Authenticate(&handler.GraphQL{Schema: schema}))

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	//////////////////
	// serving single page application stuff
	//////////////////

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./frontend/public", true)))

	// if no route then serve the root file
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/public/index.html")
	})

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Start and run the server depending on mode
	if gin.Mode() == gin.DebugMode {
		log.Fatal(router.Run(":3000"))
	} else if gin.Mode() == gin.ReleaseMode {
		log.Fatal(autotls.Run(router, "gibjob.engineer"))
	}
}

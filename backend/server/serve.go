package server

import (
	"context"
	"log"
	"net/http"

	// "github.com/gin-gonic/autotls"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/GibJob-ai/GObjob/config"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/handler"
	"github.com/GibJob-ai/GObjob/resolvers"
	"github.com/GibJob-ai/GObjob/schema"
)

func Serve(db *db.DB) {
	context.Background()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(NoFlocMiddleware())

	//////////////////
	// serving single page application stuff
	//////////////////

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../frontend/public", true)))

	// if no route then serve the root file
	router.NoRoute(func(c *gin.Context) {
		c.File("../frontend/public/index.html")
	})

	// graphql routes and route handling
	// only enable graphql playground on dev mode
	if gin.Mode() == gin.DebugMode {
		router.GET("/graphql", gin.WrapH(handler.GraphiQL{}))
	}
	router.POST("/graphql", gin.WrapH(handler.Authenticate(&handler.GraphQL{Schema: schema})))

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
		log.Fatal(router.Run(":" + config.CONFIG.Port))
	} else if gin.Mode() == gin.ReleaseMode {
		log.Fatal(router.Run(":" + config.CONFIG.Port))
	}
			/*autotls.Run(router, "gibjob.com")*/
}

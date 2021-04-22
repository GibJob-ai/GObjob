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
	"github.com/GibJob-ai/GObjob/config"
	"github.com/GibJob-ai/GObjob/db"
	"github.com/GibJob-ai/GObjob/handler"
	"github.com/GibJob-ai/GObjob/resolvers"
	"github.com/GibJob-ai/GObjob/schema"
)

func main() {

	// load the config
	config.Load()

	// if the config is dev mode set server to dev mode
	// same with other modes
	if config.CONFIG.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if config.CONFIG.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.CONFIG.Mode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		log.Fatalf("Unkown server mode given in config: %s", config.CONFIG.Mode)
	}

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	context.Background()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

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
		log.Fatal(autotls.Run(router, "gibjob.engineer"))
	}
}

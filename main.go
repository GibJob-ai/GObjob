package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
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

// type Response struct {
// Message string `json:"message"`
// }

// type Jwks struct {
// Keys []JSONWebKeys `json:"keys"`
// }

// type JSONWebKeys struct {
// Kty string   `json:"kty"`
// Kid string   `json:"kid"`
// Use string   `json:"use"`
// N   string   `json:"n"`
// E   string   `json:"e"`
// X5c []string `json:"x5c"`
// }

// var jwtMiddleWare *jwtmiddleware.JWTMiddleware

// func main() {
// jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
// ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// aud := os.Getenv("AUTH0_API_AUDIENCE")
// checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
// if !checkAudience {
// return token, errors.New("Invalid audience.")
// }
// // verify iss claim
// iss := os.Getenv("AUTH0_DOMAIN")
// checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
// if !checkIss {
// return token, errors.New("Invalid issuer.")
// }

// cert, err := getPemCert(token)
// if err != nil {
// log.Fatalf("could not get cert: %+v", err)
// }

// result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
// return result, nil
// },
// SigningMethod: jwt.SigningMethodRS256,
// })

// jwtMiddleWare = jwtMiddleware
// // Set the router as the default one shipped with Gin
// router := gin.Default()

// // Serve the frontend
// router.Use(static.Serve("/", static.LocalFile("./views", true)))

// api := router.Group("/api")
// {
// api.GET("/", func(c *gin.Context) {
// c.JSON(http.StatusOK, gin.H{
// "message": "pong",
// })
// })
// api.GET("/jokes", authMiddleware(), JokeHandler)
// api.POST("/jokes/like/:jokeID", authMiddleware(), LikeJoke)
// }
// // Start the app
// router.Run(":3000")
// }

// func getPemCert(token *jwt.Token) (string, error) {
// cert := ""
// resp, err := http.Get(os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
// if err != nil {
// return cert, err
// }
// defer resp.Body.Close()

// var jwks = Jwks{}
// err = json.NewDecoder(resp.Body).Decode(&jwks)

// if err != nil {
// return cert, err
// }

// x5c := jwks.Keys[0].X5c
// for k, v := range x5c {
// if token.Header["kid"] == jwks.Keys[k].Kid {
// cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
// }
// }

// if cert == "" {
// return cert, errors.New("unable to find appropriate key")
// }

// return cert, nil
// }

// // authMiddleware intercepts the requests, and check for a valid jwt token
// func authMiddleware() gin.HandlerFunc {
// return func(c *gin.Context) {
// // Get the client secret key
// err := jwtMiddleWare.CheckJWT(c.Writer, c.Request)
// if err != nil {
// // Token not found
// fmt.Println(err)
// c.Abort()
// c.Writer.WriteHeader(http.StatusUnauthorized)
// c.Writer.Write([]byte("Unauthorized"))
// return
// }
// }
// }

// // JokeHandler returns a list of jokes available (in memory)
// func JokeHandler(c *gin.Context) {
// c.Header("Content-Type", "application/json")

// c.JSON(http.StatusOK, jokes)
// }

// func LikeJoke(c *gin.Context) {
// // Check joke ID is valid
// if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
// // find joke and increment likes
// for i := 0; i < len(jokes); i++ {
// if jokes[i].ID == jokeid {
// jokes[i].Likes = jokes[i].Likes + 1
// }
// }
// c.JSON(http.StatusOK, &jokes)
// } else {
// // the jokes ID is invalid
// c.AbortWithStatus(http.StatusNotFound)
// }
// }

// // getJokesByID returns a single joke
// func getJokesByID(id int) (*Joke, error) {
// for _, joke := range jokes {
// if joke.ID == id {
// return &joke, nil
// }
// }
// return nil, errors.New("Joke not found")
// }

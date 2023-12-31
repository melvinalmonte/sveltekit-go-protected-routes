package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {

	var body Auth
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	// Throws unauthorized error
	if body.Username != "test" || body.Password != "test" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "almontem"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Second * 60).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	log.Println(t)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success! The status is 200",
	})
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello " + username,
	})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	}))

	// Login route
	e.POST("/api/login", login)

	// Unauthenticated route
	e.GET("/api/unrestricted", accessible)

	// Restricted group
	r := e.Group("/api/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

package main
import (
	//Echo framework library
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `"${time_rfc3339}" with Response {"id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out}}` + "\n",
	  }))

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//Route => Handler
	e.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	//Server
	e.Logger.Fatal(e.Start(":1323"))
}
// // Route => handler  
// e.GET("/", func(c echo.Context) error {       
// 	return c.String(http.StatusOK, "Hello, World!\n")  
// })
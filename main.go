package main

import (
	"flag"
	"github.com/biezhi/gorm-paginator/pagination"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

const (
	VERSION = "v0.0.1"
)


type Project struct {
	gorm.Model
	Permalink	string `gorm:"size:64;unique"`
	Name	string `gorm:"size:64"`
	Type	string `gorm:"size:64"`
	Thumbnail	string `gorm:"size:512"`
	Media	string `gorm:"size:512"`
	Description	string	`gorm:"size:512"`
}

func main() {
	address := flag.String("port", ":3050", "specify the listening address of the application")
	dist := flag.String("dist", "ui/dist", "specify the dist directory of the application")
	sqlConfig := flag.String("sql", "portfolio:portfolio@tcp(10.6.6.66)/portfolio?charset=utf8&parseTime=True&loc=Local", "specify the sql uri of the application")
	flag.Parse()

	e := echo.New()
	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://example.com"},
		AllowMethods: []string{http.MethodGet},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))


	// Database
	db, err := gorm.Open("mysql", *sqlConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Project{})

	// Routes
	e.Static("/_nuxt", *dist + "/_nuxt/")

	api := e.Group("/api")

	api.GET("/projects", func(c echo.Context) error {
		queryPage, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		queryLimit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || queryLimit > 20 {
			return c.JSON(http.StatusBadRequest, nil)
		}

		var projects []Project
		paginator := pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    queryPage,
			Limit:   queryLimit,
			OrderBy: []string{"id desc"},
			ShowSQL: true,
		}, &projects)
		return c.JSON(http.StatusOK, paginator)
	})

	api.GET("/projects/:id", func(c echo.Context) error {
		id := c.Param("id")
		var project Project
		db.Where("permalink = ?", id).Find(&project)
		return c.JSON(http.StatusOK, project)
	})

	e.File("*", *dist + "/index.html")

	// Start server
	e.Logger.Fatal(e.Start(*address))
}

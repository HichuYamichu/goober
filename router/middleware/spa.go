package middleware

import "github.com/labstack/echo/v4/middleware"

var ServeSPA = middleware.StaticWithConfig(middleware.StaticConfig{
	Skipper: middleware.DefaultSkipper,
	Root:    "web/public/",
	Index:   "index.html",
	HTML5:   true,
	Browse:  false,
})

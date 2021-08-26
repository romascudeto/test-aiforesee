package fuel

import (
	"aiforesee/fuel/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	g := e.Group("/fuel")
	g.GET("", controllers.GetListFuel)
	g.GET("/:id", controllers.GetFuelByID)
	g.POST("", controllers.CreateFuel)
	g.PUT("", controllers.UpdateFuel)
	g.DELETE("", controllers.DeleteFuel)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func InvoiceRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/invoices", controllers.GetInvoices())
	incomingRoute.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingRoute.POST("/invoices", controllers.CreateInvoice())
	incomingRoute.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}

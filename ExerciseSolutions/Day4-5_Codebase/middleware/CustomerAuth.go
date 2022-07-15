package middleware

import (
	Controller "Day4-5_Codebase/controller"
	Utils "Day4-5_Codebase/db_utils"
	"Day4-5_Codebase/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuthCustomer(c *gin.Context) {
	// Get the Basic Authentication credentials
	userName, password, hasAuth := c.Request.BasicAuth()
	res, err := Controller.NewCustomer(service.NewCustomerService(Utils.GetDB())).IsCustomerAuthenticated(c)
	if hasAuth && err == nil && res.Name == userName && res.Password == password {
		c.Writer.Header().Set(userName+" with "+password, "is authenticated")
	} else {
		c.Writer.Header().Set("Authentication-Info", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "please provide correct credentials first"})
		c.Abort()
		return
	}
}

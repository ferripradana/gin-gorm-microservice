package errors

import (
	"gin-gorm-microservice/domain/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MessagesResponse is a struct that contains the response body for the message
type MessagesResponse struct {
	Message string `json:"message"`
}

// Handler is Gin Middleware to handle errors.
func Handler(c *gin.Context) {
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(errors.AppError)
		if ok {
			resp := MessagesResponse{Message: err.(*errors.AppErrorImpl).Err.Error()}
			switch err.(*errors.AppErrorImpl).Type {
			case errors.NotFound:
				c.JSON(http.StatusNotFound, resp)
			case errors.ValidationError:
				c.JSON(http.StatusBadRequest, resp)
				return
			case errors.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, resp)
				return
			case errors.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, resp)
			case errors.NotAuthorized:
				c.JSON(http.StatusForbidden, resp)
			case errors.RepositoryError:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
			default:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			}
		}
		return
	}
}

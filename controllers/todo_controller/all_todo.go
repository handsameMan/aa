package todo_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTodo(context *gin.Context) {
	parsedClaims, err := middlewares.GetClaims(context)
	authRes := helper.Res{Status: http.StatusBadRequest}
	res := helper.Res{}

	if err != nil {
		authRes.Err = err
		authRes.Send(context)
		return
	}

	todo := models.UserTodo{
		UserID: parsedClaims.UserId,
	}
	todos, err := todo.GetAllTodoItemsByID()

	if err != nil {
		authRes.Err = robust.DB_READ_FAILURE
		authRes.Send(context)
		return
	}

	res.Data = todos
	res.Send(context)
}

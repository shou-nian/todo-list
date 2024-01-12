package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/riny/demo-go-gin/app/config"
	"github.com/riny/demo-go-gin/app/model"
	"github.com/riny/demo-go-gin/app/util"
	"log/slog"
	"net/http"
)

type TodoListManagement interface {
	AddTodoList(context *gin.Context)
	QueryTodoList(context *gin.Context)
	UpdateTodoStatus(context *gin.Context)
	DeleteTodoList(context *gin.Context)
}

type TodoManager struct {
	TodoMap map[string]string
}

func New() *TodoManager {
	return &TodoManager{
		TodoMap: make(map[string]string),
	}
}

func (t *TodoManager) AddTodoList(context *gin.Context) {
	var todo model.TodoList
	// 反序列化json数据
	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			},
		)
		slog.Error(err.Error())
		return
	}

	// name长度检查
	if len(todo.TodoName) > 100 || len(todo.TodoName) == 0 {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.NameLengthError,
			},
		)
		return
	}

	// 是否已在todo列表中
	if t.TodoMap[todo.TodoName] != "" {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.NameExistsError,
			},
		)
		return
	}
	// 添加到todo列表中
	t.TodoMap[todo.TodoName] = "open"
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "add todo list success",
			"data":    map[string]string{"todoName": todo.TodoName, "status": "open"},
		},
	)
}

func (t *TodoManager) QueryTodoList(context *gin.Context) {
	if len(t.TodoMap) == 0 {
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": config.TodoListNullError,
			},
		)
		return
	}

	// 获取请求参数: name
	queryName, ok := context.GetQuery("name")
	if !ok {
		// 如果未指定name，返回所有数据
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "success",
				"data":    util.QueryAll(t.TodoMap),
			},
		)
		return
	}

	// 如果指定的name不存在
	if t.TodoMap[queryName] == "" {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.NameExistsError,
			},
		)
		return
	}

	// 返回满足查询条件的数据
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data":    map[string]string{"name": queryName, "status": t.TodoMap[queryName]},
		},
	)
}

func (t *TodoManager) UpdateTodoStatus(context *gin.Context) {
	var todo model.TodoList
	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			},
		)
		return
	}

	// 检查name是否存在
	if t.TodoMap[todo.TodoName] == "" {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.TodoNameNotExistsError,
			},
		)
		return
	}

	// name长度校验
	if len(todo.TodoName) > 100 || len(todo.TodoName) == 0 {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.NameLengthError,
			},
		)
		return
	}

	// 更新状态
	t.TodoMap[todo.TodoName] = util.ReverseStatus(t.TodoMap[todo.TodoName])
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "status updated",
			"data":    map[string]string{"name": todo.TodoName, "status": t.TodoMap[todo.TodoName]},
		},
	)
}

func (t *TodoManager) DeleteTodoList(context *gin.Context) {
	if context.GetHeader("user-id") != "1" {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": http.StatusText(http.StatusUnauthorized),
			},
		)
		return
	}

	var todo model.TodoList
	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			},
		)
		return
	}

	// 检查name是否存在
	if t.TodoMap[todo.TodoName] == "" {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.TodoNameNotExistsError,
			},
		)
		return
	}

	// name长度校验
	if len(todo.TodoName) > 100 || len(todo.TodoName) == 0 {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": config.NameLengthError,
			},
		)
		return
	}

	// 删除todo list
	delete(t.TodoMap, todo.TodoName)
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "delete todo list success",
		},
	)
}

package presenters

import (
	"github.com/IzumiSy/go-cleanarchitecture/domains/models"
	"github.com/labstack/echo"
)

type getTodosResponse struct {
	Todos []getTodosResponseItem `json:"todos"`
}

type getTodosResponseItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type getTodosJsonPresenter struct {
	ctx    echo.Context
	result error
}

func NewGetTodosJsonPresenter(ctx echo.Context) *getTodosJsonPresenter {
	return &getTodosJsonPresenter{ctx: ctx}
}

// エラーハンドリングはサボって全部500を返している
func (presenter *getTodosJsonPresenter) Write(todos []models.Todo, err error) {
	if err != nil {
		presenter.result = presenter.ctx.String(500, "Internal Server Error")
		return
	}

	response := getTodosResponse{
		Todos: []getTodosResponseItem{},
	}

	for _, todo := range todos {
		response.Todos = append(response.Todos, getTodosResponseItem{
			Id:   todo.Id,
			Name: todo.Name,
		})
	}

	presenter.result = presenter.ctx.JSON(200, response)
}

func (presenter getTodosJsonPresenter) Present() error {
	return presenter.result
}

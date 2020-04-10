package domains

import (
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

// repositoriesみたいなネームスペースを切ってもせいぜいインタフェースしか置かれないので
// domainsパッケージ直下にリポジトリを一つのファイルとして書いていくスタイルとした。

type TodoRepository interface {
	Get(id todo.Id) (error, models.Todo)
	GetByName(name todo.Name) (error, models.Todo)
	Store(todo models.Todo) error
}

type TodosRepository interface {
	Get() (error, models.Todos)
}

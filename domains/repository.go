package domains

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/todo"
)

// repositoriesみたいなネームスペースを切ってもせいぜいインタフェースしか置かれないので
// domainsパッケージ直下にリポジトリを一つのファイルとして書いていくスタイルとした。

type TodoRepository interface {
	Get(id todo.Id) (models.Todo, errors.Domain, bool)
	GetByName(name todo.Name) (models.Todo, errors.Domain, bool)
	Store(todo models.Todo) errors.Domain
}

type TodosRepository interface {
	Get() (models.Todos, errors.Domain, bool)
}

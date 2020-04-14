package domains

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/category"
	"go-cleanarchitecture/domains/models/todo"
)

// repositoriesみたいなネームスペースを切ってもせいぜいインタフェースしか置かれないので
// domainsパッケージ直下にリポジトリを一つのファイルとして書いていくスタイルとした。
// 三値リターンでboolを返しているのは、エラーで存在の有無を表すとエラーハンドリングを時点で
// 存在チェックを忘れてしまう可能性があるのを防ぐため。

type TodoRepository interface {
	Get(id todo.Id) (models.Todo, errors.Domain, bool)
	GetByName(name todo.Name) (models.Todo, errors.Domain, bool)
	Store(todo models.Todo) errors.Domain
}

// 以下のTodosRepositoryやCategoriesRepositoryなどは
// 集合に対するクエリなので戻り値に存在の有無を表すboolは持たない

type TodosRepository interface {
	GetByIds(ids []todo.Id) (models.Todos, errors.Domain)
	Get() (models.Todos, errors.Domain)
}

type CategoriesRepository interface {
	GetByIds(ids []category.Id) ([]models.Category, errors.Domain)
	Get() ([]models.Category, errors.Domain)
}

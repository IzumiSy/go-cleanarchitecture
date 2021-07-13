package domains

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/category"
	"go-cleanarchitecture/domains/models/session"
	"go-cleanarchitecture/domains/models/todo"
	"go-cleanarchitecture/domains/models/user"
)

// repositoriesみたいなネームスペースを切ってもせいぜいインタフェースしか置かれないので
// domainsパッケージ直下にリポジトリを一つのファイルとして書いていくスタイルとした。
// 三値リターンでboolを返しているのは、エラーで存在の有無を表すとエラーハンドリングを時点で
// 存在チェックを忘れてしまう可能性があるのを防ぐため。

// TodosRepositoryやCategoriesRepositoryなどは
// 集合に対するクエリなので戻り値に存在の有無を表すboolは持たない

type TodoRepository interface {
	Get(id todo.ID) (models.Todo, errors.Domain, bool)
	GetByName(name todo.Name) (models.Todo, errors.Domain, bool)
	Store(todo models.Todo) errors.Domain
}

type TodosRepository interface {
	GetByIDs(ids []todo.ID) (models.Todos, errors.Domain)
	GetByUserID(userID user.ID) (models.Todos, errors.Domain)
}

type CategoriesRepository interface {
	GetByIDs(ids []category.ID) ([]models.Category, errors.Domain)
	GetByUserID(userId user.ID) ([]models.Category, errors.Domain)
}

// 認証系

type AuthenticationRepository interface {
	Store(auth models.Authentication) errors.Domain
	GetByEmail(email authentication.Email) (models.Authentication, errors.Domain, bool)
}

type SessionRepository interface {
	Store(session models.Session) errors.Domain
	Get(id session.ID) (models.Session, errors.Domain, bool)
}

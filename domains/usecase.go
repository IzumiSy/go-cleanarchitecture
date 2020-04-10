package domains

import (
	"go-cleanarchitecture/domains/errors"
)

type OutputPort interface {
	Raise(err errors.Domain)
}

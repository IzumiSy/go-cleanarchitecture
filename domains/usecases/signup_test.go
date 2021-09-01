package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/testing/adapters"
	"go-cleanarchitecture/testing/adapters/dao"
	"testing"
	"time"

	"golang.org/x/xerrors"
)

type mockSignupOutputPort struct {
	Error  errors.Domain
	Result models.Authentication
}

func (m *mockSignupOutputPort) Raise(err errors.Domain) {
	m.Error = err
}

func (m *mockSignupOutputPort) Write(auth models.Authentication) {
	m.Result = auth
}

type mockSignupPublisher struct {
	event UserSignedUpEvent
}

func (p *mockSignupPublisher) Publish(event domains.Event) errors.Domain {
	p.event = event.(UserSignedUpEvent)
	return errors.None
}

func TestSignupUsecase(t *testing.T) {
	newUsecase := func(
		authDao domains.AuthenticationRepository,
		op *mockSignupOutputPort,
		p *mockSignupPublisher,
	) SignupUsecase {
		return SignupUsecase{
			OutputPort:        op,
			AuthenticationDao: authDao,
			Logger:            adapters.MockLogger{T: t},
			Publisher:         p,
		}
	}

	t.Run("Validation", func(t *testing.T) {
		// todo
	})

	t.Run("Store", func(t *testing.T) {
		op := &mockSignupOutputPort{}
		p := &mockSignupPublisher{}
		newUsecase(dao.NewMockAuthenticationDao(), op, p).
			Build(SignupParam{
				Email:    "test@example.com",
				Password: "password",
				UserName: "test",
			}).
			Run()

		if !xerrors.Is(op.Error, errors.None) {
			t.Errorf("Unexpected error raised: %s", op.Error)
		}

		if op.Result.CreatedAt().Value() == (time.Time{}) {
			t.Error("Error: invalid result")
		}

		if p.event == (UserSignedUpEvent{}) {
			t.Error("Error: event not published")
		}
	})
}

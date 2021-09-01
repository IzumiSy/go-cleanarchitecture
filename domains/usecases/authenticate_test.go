package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/user"
	"go-cleanarchitecture/testing/adapters"
	"go-cleanarchitecture/testing/adapters/dao"
	"testing"
	"time"

	"golang.org/x/xerrors"
)

type mockAuthenticateOutputPort struct {
	Error  errors.Domain
	Result models.Session
}

func (m *mockAuthenticateOutputPort) Raise(err errors.Domain) {
	m.Error = err
}

func (m *mockAuthenticateOutputPort) Write(session models.Session) {
	m.Result = session
}

type mockAuthenticatePublisher struct {
	event UserAuthenticatedEvent
}

func (p *mockAuthenticatePublisher) Publish(event domains.Event) errors.Domain {
	p.event = event.(UserAuthenticatedEvent)
	return errors.None
}

func TestAuthenticateUsecase(t *testing.T) {
	newUsecase := func(
		authDao domains.AuthenticationRepository,
		sessionDao domains.SessionRepository,
		op *mockAuthenticateOutputPort,
		p *mockAuthenticatePublisher,
	) AuthenticateUsecase {
		return AuthenticateUsecase{
			OutputPort:        op,
			AuthenticationDao: authDao,
			SessionDao:        sessionDao,
			Logger:            adapters.MockLogger{T: t},
			Publisher:         p,
		}
	}

	t.Run("Validation", func(t *testing.T) {
		// todo
	})

	t.Run("Store", func(t *testing.T) {
		aDao := dao.NewMockAuthenticationDao()
		aDao.GetByEmailResult = func() (models.Authentication, errors.Domain, bool) {
			email, _ := authentication.NewEmail("test@example.com")
			hash := authentication.NewHash("password")
			a := models.NewAuthentication(email, hash, user.Name{})
			return a, errors.None, true
		}

		op := &mockAuthenticateOutputPort{}
		p := &mockAuthenticatePublisher{}
		newUsecase(aDao, dao.NewMockSessionDao(), op, p).
			Build(AuthenticateParam{
				Email:    "test@example.com",
				Password: "password",
			}).
			Run()

		if !xerrors.Is(op.Error, errors.None) {
			t.Errorf("Unexpected error raised: %s", op.Error)
		}

		if op.Result.CreatedAt().Value() == (time.Time{}) {
			t.Error("Error: invalid result")
		}

		if p.event == (UserAuthenticatedEvent{}) {
			t.Error("Error: event not published")
		}
	})
}

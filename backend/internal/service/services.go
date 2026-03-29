package service

import(
	"github.com/sarbojitrana/go-boilerplate/internal/lib/job"
	"github.com/sarbojitrana/go-boilerplate/internal/server"
	"github.com/sarbojitrana/go-boilerplate/internal/repository"
)

type Services struct{
	Auth *AuthService
	Job *job.JobService
}

func NewServices( s *server.Server, repos *repository.Repositories) (*Services, error){
	authService := NewAuthService(s)

	return &Services{
		Job : s.Job,
		Auth: authService,
	}, nil
}
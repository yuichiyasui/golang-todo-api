package handler

import (
	"api/domain"
	domainRepository "api/domain/repository"
	"api/gen"
	"api/infrastructure"
	"api/repository"
	"database/sql"
)

type Server struct {
	db                               *sql.DB
	tasksRepository                  repository.TasksRepositoryInterface
	userRegistrationTokensRepository domainRepository.UserRegistrationTokensRepositoryInterface
	mailer                           domain.MailerInterface
}

// Make sure we conform to StrictServerInterface

var _ gen.StrictServerInterface = (*Server)(nil)

func NewServer(db *sql.DB) (*Server, error) {
	tasksRepo, err := repository.NewTasksRepository(db)
	if err != nil {
		return nil, err
	}

	userRegistrationTokensRepo := repository.NewUserRegistrationTokensRepository(db)
	mailer := infrastructure.NewMailer()

	return &Server{
		db:                               db,
		tasksRepository:                  tasksRepo,
		userRegistrationTokensRepository: userRegistrationTokensRepo,
		mailer:                           mailer,
	}, nil
}

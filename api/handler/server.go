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
	usersRepository                  domainRepository.UsersRepositoryInterface
	mailer                           domain.MailerInterface
}

// Make sure we conform to StrictServerInterface

var _ gen.StrictServerInterface = (*Server)(nil)

func NewServer(db *sql.DB) (*Server, error) {
	tasksRepo := repository.NewTasksRepository(db)
	userRegistrationTokensRepo := repository.NewUserRegistrationTokensRepository(db)
	usersRepo := repository.NewUsersRepository(db)
	mailer := infrastructure.NewMailer()

	return &Server{
		db:                               db,
		tasksRepository:                  tasksRepo,
		userRegistrationTokensRepository: userRegistrationTokensRepo,
		usersRepository:                  usersRepo,
		mailer:                           mailer,
	}, nil
}

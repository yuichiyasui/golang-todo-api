package handler

import (
	"api/domain"
	"api/gen"
	"context"
	"net/http"
)

func (s *Server) SignUp(ctx context.Context, request gen.SignUpRequestObject) (gen.SignUpResponseObject, error) {
	urt, err := s.userRegistrationTokensRepository.FindByToken(ctx, request.Body.Token)
	if err != nil {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "ERROR",
				Message: err.Error(),
			},
		}, err
	}

	isExpired := urt.IsExpired()
	if isExpired {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "ERROR",
				Message: "トークンの有効期限が切れています",
			},
		}, nil
	}

	email := urt.Email()
	usr, err := s.usersRepository.FindByEmail(ctx, email.Value())
	if err != nil {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "ERROR",
				Message: err.Error(),
			},
		}, err
	}
	exists := usr != nil
	if exists {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "ERROR",
				Message: "メールアドレスは既に登録されています",
			},
		}, nil
	}

	password, err := domain.NewPassword(request.Body.Password)
	if err != nil {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "ERROR",
				Message: err.Error(),
			},
		}, err
	}

	usr, err = domain.NewUser("", request.Body.Username, email.Value(), password.Value())
	if err != nil {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "ERROR",
				Message: err.Error(),
			},
		}, nil
	}

	err = s.usersRepository.Save(ctx, *usr)
	if err != nil {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "ERROR",
				Message: err.Error(),
			},
		}, err
	}

	err = s.userRegistrationTokensRepository.DeleteByToken(ctx, request.Body.Token)
	if err != nil {
		return gen.SignUpdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "ERROR",
				Message: err.Error(),
			},
		}, err
	}

	return gen.SignUp201Response{}, nil
}

package handler

import (
	"api/domain"
	"api/gen"
	"context"
	"errors"
	"net/http"
)

func (s *Server) SendSignUpEmail(ctx context.Context, request gen.SendSignUpEmailRequestObject) (gen.SendSignUpEmailResponseObject, error) {
	email, err := domain.NewEmail(string(request.Body.Email))
	if err != nil {
		return gen.SendSignUpEmaildefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	return gen.SendSignUpEmaildefaultJSONResponse{
		StatusCode: http.StatusInternalServerError,
		Body: gen.Error{
			Code:    "",
			Message: "not implemented",
		},
	}, errors.New("not implemented")

	token, err := domain.NewUserRegistrationToken()
	if err != nil {
		return gen.SendSignUpEmaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	url := token.GenerateSignUpUrl()

	err = s.mailer.SendEmail(email.Value(), "ユーザー登録のご案内", url)
	if err != nil {
		return gen.SendSignUpEmaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	err = s.userRegistrationTokensRepository.Save(ctx, token.Value())
	if err != nil {
		return gen.SendSignUpEmaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	return gen.SendSignUpEmail200Response{}, nil
}

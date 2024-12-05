package handler

import (
	"api/domain"
	"api/gen"
	"context"
	"net/http"
)

func (s *Server) SendRegistrationEmail(ctx context.Context, request gen.SendRegistrationEmailRequestObject) (gen.SendRegistrationEmailResponseObject, error) {
	email, err := domain.NewEmail(string(request.Body.Email))
	if err != nil {
		return gen.SendRegistrationEmaildefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, nil
	}

	token, err := domain.NewUserRegistrationToken()
	if err != nil {
		return gen.SendRegistrationEmaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, nil
	}

	url := token.GenerateUserRegistrationUrl()

	err = s.mailer.SendEmail(email.Value(), "ユーザー登録のご案内", url)
	if err != nil {
		return gen.SendRegistrationEmaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, nil
	}

	err = s.userRegistrationTokensRepository.Save(ctx, token.Value())
	if err != nil {
		return gen.SendRegistrationEmaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, nil
	}

	return gen.SendRegistrationEmail200Response{}, nil
}

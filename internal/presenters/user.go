package presenters

import "edot/internal/domain"

type UserPresenter struct{}

func (p *UserPresenter) Response(user *domain.User) Response {
	return Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       user,
	}
}

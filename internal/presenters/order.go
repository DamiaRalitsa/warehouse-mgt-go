package presenters

import "edot/internal/domain"

type OrderPresenter struct{}

func (p *OrderPresenter) Response(order *domain.Order) Response {
	return Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       order,
	}
}

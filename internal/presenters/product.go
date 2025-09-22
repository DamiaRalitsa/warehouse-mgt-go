package presenters

import "edot/internal/domain"

type ProductPresenter struct{}

func (p *ProductPresenter) Response(product *domain.Product) Response {
	return Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       product,
	}
}

package presenters

import "edot/internal/domain"

type ShopPresenter struct{}

func (p *ShopPresenter) Response(shop *domain.Shop) Response {
	return Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       shop,
	}
}

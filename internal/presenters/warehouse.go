package presenters

import "edot/internal/domain"

type WarehousePresenter struct{}

func (p *WarehousePresenter) Response(warehouse *domain.Warehouse) Response {
	return Response{
		StatusCode: 200,
		Message:    "success",
		Success:    true,
		Data:       warehouse,
	}
}

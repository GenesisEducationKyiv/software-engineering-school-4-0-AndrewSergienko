package presentation

import (
	"go_service/internal/rateservice/customers/services/createcustomer"
	"go_service/internal/rateservice/customers/services/deletecustomer"
	"go_service/internal/rateservice/customers/services/getall"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	CreateCustomer() Interactor[createcustomer.InputData, createcustomer.OutputData]
	DeleteCustomer() Interactor[deletecustomer.InputData, deletecustomer.OutputData]
	GetAll() Interactor[getall.InputDTO, getall.OutputDTO]
}

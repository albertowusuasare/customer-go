package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

type CustomerRetrieveResponseDTO struct {
	CustomerID       string `json:"customerId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	NationalID       string `json:"nationalId"`
	PhoneNumber      string `json:"phoneNumber"`
	AccountID        string `json:"accountId"`
	LastModifiedTime string `json:"lastModifiedTime"`
	CreatedTime      string `json:"createdTime"`
	Version          int    `json:"version"`
}

type RetrieveOneHandler struct {
	Workflow workflow.RetrieveSingleFunc
}

type RetrieveMultiHandler struct {
	Workflow workflow.RetrieveMultiFunc
}

func (handler RetrieveOneHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customerId := RetrieveCustomerId(r)
	customer := handler.Workflow(customerId)
	response := customerRetrieveResponseDTOFromCustomer(customer)
	json.NewEncoder(w).Encode(response)
}

func (handler RetrieveMultiHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := retrieving.MultiRequest{}
	log.Printf("Retrieving customers for request=%+v", request)
	customers := handler.Workflow(request)
	response := customerRetrieveResponseDTOsFromCustomers(customers)
	json.NewEncoder(w).Encode(response)
}

func customerRetrieveResponseDTOsFromCustomers(customers []retrieving.Customer) []CustomerRetrieveResponseDTO {
	response := make([]CustomerRetrieveResponseDTO, len(customers))
	for i, c := range customers {
		response[i] = customerRetrieveResponseDTOFromCustomer(c)
	}
	return response
}

func customerRetrieveResponseDTOFromCustomer(customer retrieving.Customer) CustomerRetrieveResponseDTO {
	return CustomerRetrieveResponseDTO{
		CustomerID:       customer.CustomerId,
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		NationalID:       customer.NationalId,
		PhoneNumber:      customer.PhoneNumber,
		AccountID:        customer.AccountId,
		LastModifiedTime: customer.LastModifiedTime,
		CreatedTime:      customer.CreatedTime,
		Version:          customer.Version,
	}
}

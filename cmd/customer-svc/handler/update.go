package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// UpdateHandler represents the http handler for a customer update http call
type UpdateHandler struct {
	Workflow workflow.UpdateFunc
}

// UpdateRequestDTO represents the json structure for a customer update request
type UpdateRequestDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalId  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumer"`
}

// UpdateResponseDTO represents the json structure for a customer update response
type UpdateResponseDTO struct {
	CustomerId       string `json:"customerId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	NationalId       string `json:"nationalId"`
	PhoneNumber      string `json:"phoneNumber"`
	AccountId        string `json:"accountId"`
	LastModifiedTime string `json:"lastModifiedTime"`
	CreatedTime      string `json:"createdTime"`
	Version          int    `json:"version"`
}

// Handle allows the UpdateHandler to act as an http call handler
func (handler UpdateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestDTO UpdateRequestDTO
	err := decoder.Decode(&requestDTO)
	if err != nil {
		panic(err)
	}
	customerId := RetrieveCustomerId(r)
	w.Header().Set("Content-Type", "application/json")
	request := updateRequestFromUpdateRequestDTO(customerId, requestDTO)
	log.Printf("Updating customer for request=%+v", request)
	updatedCustomer := handler.Workflow(request)
	response := updateResponseDTOFromUpdatedCustomer(updatedCustomer)
	json.NewEncoder(w).Encode(response)
}

func updateRequestFromUpdateRequestDTO(customerId string, dto UpdateRequestDTO) updating.Request {
	return updating.Request{
		CustomerId:  customerId,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		NationalId:  dto.NationalId,
		PhoneNumber: dto.PhoneNumber,
	}
}

func updateResponseDTOFromUpdatedCustomer(customer updating.UpdatedCustomer) UpdateResponseDTO {
	return UpdateResponseDTO{
		CustomerId:       customer.CustomerId,
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		NationalId:       customer.NationalId,
		PhoneNumber:      customer.PhoneNumber,
		AccountId:        customer.AccountId,
		LastModifiedTime: customer.LastModifiedTime,
		CreatedTime:      customer.CreatedTime,
		Version:          customer.Version,
	}
}

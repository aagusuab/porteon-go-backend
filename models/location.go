package models

// Location model
type Location struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	PickupAddress string `json:"pickup_address"`
	CreatedAt     string `json:"created_at"`
}

func NewLocation(name string, address string, phone string, pickupAddress string, createdAt string) Location {
	return Location{
		ID:            randomString(16),
		Name:          name,
		Address:       address,
		Phone:         phone,
		PickupAddress: pickupAddress,
		CreatedAt:     createdAt,
	}
}

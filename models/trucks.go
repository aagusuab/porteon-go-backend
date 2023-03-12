package models

type Truck struct {
	ID                 string      `json:"id"`
	VID                string      `json:"vid"`
	VehicleModel       string      `json:"vehicle_model"`
	VehicleBodyType    string      `json:"vehicle_body_type"`
	VehicleCreatedYear uint        `json:"vehicle_created_year"`
	CurrentDriver      string      `json:"current_driver"`
	Status             TruckStatus `json:"status" gorm:"type:truck_status"`
}

type TruckStatus string

const (
	TruckReady            TruckStatus = "Ready"
	TruckOnDuty           TruckStatus = "On Duty"
	TruckRepairInProgress TruckStatus = "Repair in Progress"
	TruckInAccident       TruckStatus = "InAccident"
	TruckDeactivated      TruckStatus = "Deactivated"
)

func NewTruck(vehicleModel string, vehicleBodyType string, vehicleCreatedYear uint, currentDriver string, status TruckStatus) Truck {
	return Truck{
		ID:                 randomString(16),
		VehicleModel:       vehicleModel,
		VehicleBodyType:    vehicleBodyType,
		VehicleCreatedYear: vehicleCreatedYear,
		CurrentDriver:      currentDriver,
		Status:             status,
	}
}

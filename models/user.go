package models

import (
	"time"
)

// User model
type User struct {
	ID        string     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       uint       `json:"age"`
	JoinDate  time.Time  `json:"join_date"`
	TruckID   string     `json:"truck_id"`
	Address   string     `json:"address"`
	SSN       string     `json:"ssn"`
	DL        string     `json:"dl"`
	IPN       string     `json:"ipn"`
	Phone     string     `json:"phone"`
	Status    UserStatus `json:"status" gorm:"type:user_status"`
}

type UserStatus string

const (
	UserReady              UserStatus = "Ready"
	UserOnDuty             UserStatus = "On Duty"
	UserRegistrationInProg UserStatus = "Registration In Progress"
	UserOnVacation         UserStatus = "On Vacation"
	UserQuit               UserStatus = "Quit"
)

func NewUser(firstName string, lastName string, age uint, joinDate time.Time, truckID string, address string, ssn string, dl string, ipn string, phone string) User {
	return User{
		ID:        randomString(16),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		JoinDate:  joinDate,
		TruckID:   truckID,
		Address:   address,
		SSN:       ssn,
		DL:        dl,
		IPN:       ipn,
		Phone:     phone,
		Status:    UserRegistrationInProg,
	}
}

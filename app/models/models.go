package models

import (
	"time"
)

type Administrator struct {
	AdminID     string    `json:"admin_id" bson:"admin_id"`
	Username    string    `json:"username" bson:"username"`
	Permissions string    `json:"permissions" bson:"permissions"`
	Password    string    `json:"password" bson:"password"`
	Email       string    `json:"email" bson:"email"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type Shipper struct {
	ShipperID        string    `json:"shipper_id" bson:"shipper_id"`
	Name             string    `json:"name" bson:"name"`
	Password         string    `json:"password" bson:"password"`
	Email            string    `json:"email" bson:"email"`
	ContactNumber    string    `json:"contact_number" bson:"contact_number"`
	Permissions      string    `json:"permissions" bson:"permissions"`
	ConfirmationCode string    `bson:"confirmation_code" json:"confirmation_code"`
	ConfirmationLink string    `bson:"confirmation_link" json:"confirmation_link"`
	EmailConfirmed   bool      `json:"email_confirmed" bson:"email_confirmed"`
	AccessToken      string    `json:"access_token" bson:"-"`
	RefreshToken     string    `json:"refresh_token" bson:"-"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

type FreightCompany struct {
	FreightCompanyID string    `json:"freight_company_id" bson:"freight_company_id"`
	Name             string    `json:"name" bson:"name" `
	Description      string    `json:"description" bson:"description"`
	Email            string    `json:"email" bson:"email"`
	Password         string    `json:"password" bson:"password"`
	Fleet            []string  `json:"fleet" bson:"fleet"`
	Contacts         []string  `json:"contacts" bson:"contacts"`
	Address          string    `json:"address" bson:"address"`
	Permissions      string    `json:"permissions" bson:"permissions"`
	IsRegistered     bool      `json:"is_registered" bson:"is_registered"`
	ConfirmationLink string    `bson:"confirmation_link" json:"confirmation_link"`
	AccessToken      string    `json:"access_token" bson:"-"`
	RefreshToken     string    `json:"refresh_token" bson:"-"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

type User struct {
	UserID           string    `json:"user_id" bson:"user_id"`
	FreightCompanyId string    `json:"freight_company_id" bson:"freight_company_id"`
	Name             string    `json:"name" bson:"name"`
	Email            string    `json:"email" bson:"email"`
	Password         string    `json:"password" bson:"password"`
	Permissions      string    `json:"permissions" bson:"permissions"`
	ContactNumber    string    `json:"contact_number" bson:"contact_number"`
	Role             string    `json:"role" bson:"role"`
	Experience       string    `json:"experience" bson:"experience"`
	AccessToken      string    `json:"access_token" bson:"-"`
	RefreshToken     string    `json:"refresh_token" bson:"-"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"`
}

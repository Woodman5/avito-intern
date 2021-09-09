package models

type UserMoney struct {
	UUID   string `json:"uuid" gorm:"primaryKey"`
	Amount uint64 `json:"amount" gorm:"not null;default:0"`
}

type Transaction struct {
	UUID            string `json:"-" gorm:"primaryKey"`
	CreatedAt       int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UserUUID        string `json:"useruuid" gorm:"not null"`
	TransactionType uint8  `json:"type" gorm:"not null"`
	Amount          uint64 `json:"amount" gorm:"not null;default:0"`
	Balance         uint64 `json:"balance" gorm:"not null;default:0"`
	Source          string `json:"source" gorm:"not null"`
	Reason          string `json:"reason" gorm:"not null"`
}

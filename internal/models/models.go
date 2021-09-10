package models

type UserMoney struct {
	UUID   string `gorm:"primaryKey"`
	Amount uint64 `gorm:"not null;default:0"`
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

type TransactionRequest struct {
	UserUUID        string  `json:"useruuid"`
	TransactionType uint8   `json:"type"`
	Amount          float64 `json:"amount"`
	Balance         float64 `json:"balance"`
	Source          string  `json:"source"`
	Reason          string  `json:"reason"`
}

type UserMoneyFloat struct {
	UUID   string  `json:"uuid"`
	Amount float64 `json:"amount"`
}

type TransferRequest struct {
	FromUUID string  `json:"from_uuid"`
	ToUUID   string  `json:"to_uuid"`
	Amount   float64 `json:"amount"`
}

type TransferAnswer struct {
	To   UserMoneyFloat `json:"to"`
	From UserMoneyFloat `json:"from"`
}

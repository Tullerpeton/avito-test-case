package models

type UserBalance struct {
	Id       uint64  `json:"-"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type ImproveBalance struct {
	Id    uint64  `json:"id"`
	Value float64 `json:"value"`
}

type WithdrawBalance struct {
	Id    uint64  `json:"id"`
	Value float64 `json:"value"`
}

type Transfer struct {
	SenderId   uint64  `json:"sender_id"`
	ReceiverId uint64  `json:"receiver_id"`
	Value      float64 `json:"value"`
}

type TransferResult struct {
	Receiver UserBalance `json:"receiver"`
	Sender   UserBalance `json:"sender"`
}

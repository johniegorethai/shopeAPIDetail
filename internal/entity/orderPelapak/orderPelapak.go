package orderPelapak

import "time"

// Skeleton ...
type OrderPelapak struct {
	partner_id        int       `db:"partner_id" json:"partner_id"`
	shopid            int       `db:"shopid" json:"shopid"`
	timestamp         time.Time `db:"timestamp" json:"timestamp"`
	release_time_from int       `db:"release_time_from" json:"release_time_from"`
	release_time_to   int       `db:"release_time_to" json:"release_time_to"`
}

type TransactionList struct {
	Amount			float64	`json:"amount" db:"amount"`
	BuyerName		string	`json:"buyer_name" db:"buyer_name"`
	CreateTime		int	`json:"create_time" db:"create_time"`
	CurrentBalance	float64	`json:"current_balance" db:"current_balance"`
	Description		string	`json:"description" db:"description"`
	OrderSN			string	`json:"ordersn" db:"ordersn"`
	Reason			string	`json:"reason" db:"reason"`
	PayOrderList	[]string	`json:"pay_order_list"`
	RefundSN		string	`json:"refund_sn" db:"refund_sn"`
	Status			string	`json:"status" db:"status"`
	TransactionFee	float64	`json:"transaction_fee" db:"transaction_fee"`
	TransactionID	int64	`json:"transaction_id" db:"transaction_id"`
	TransactionType	string	`json:"transaction_type" db:"transaction_type"`
	WalletType		string	`json:"wallet_type" db:"wallet_type"`
	WithdrawalType	string	`json:"withdrawal_type" db:"withdrawal_type"`
}

type TestShope struct {
	TotalAmount		float64		`json:"total_amount"`
	//SaldoAkhir		float64		`json:"saldo_akhir"`
	TransactionList	[]TransactionList	`json:"transaction_list"`
}

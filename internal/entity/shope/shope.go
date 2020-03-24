package shope


// Skeleton ...
type Shope struct {
	PartnerID       			int    `db:"partner_id" json:"partner_id"`
	ShopID          			int    `db:"shopid" json:"shopid"`
	Timestamp       			int64  `db:"timestamp" json:"timestamp"`
	ReleaseTimeFrom 			string `db:"release_time_from" json:"release_time_from"`
	ReleaseTimeTo   			string `db:"release_time_to" json:"release_time_to"`
	PaginationOffset			int64	`db:"pagination_offset" json:"pagination_offset"`
	PaginationEntriesPerPage	int 	`db:"pagination_entries_per_page" json:"pagination_entries_per_page"`
	CreateTimeFrom			int64	`db:"create_time_from" json:"create_time_from"`
	CreateTimeTo				int64	`db:"create_time_to" json:"create_time_to"`
}

type Shope1 struct {
	PartnerID       int    `db:"partner_id" json:"partner_id"`
	ShopID          int    `db:"shopid" json:"shopid"`
	Timestamp       int64  `db:"timestamp" json:"timestamp"`
}

type Response struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	Ordersn           string  `json:"ordersn"`
	PayoutAmount      float64 `json:"payout_amount"`
	EscrowReleaseTime int64   `json:"escrow_release_time"`
}

type MutasiShopee struct {
	HasMore		bool	`json:"has_more"`
	RequestID	string	`json:"request_id"`
	TransactionList	[]TransactionList	`json:"transaction_list"`
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

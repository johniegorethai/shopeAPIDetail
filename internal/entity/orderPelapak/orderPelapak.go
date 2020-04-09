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
	Amount          float64   `json:"amount" db:"amount"`
	BuyerName       string    `json:"buyer_name" db:"buyer_name"`
	CreateTime      int       `json:"create_time" db:"create_time"`
	CreateTimes     time.Time `json:"create_times" db:"create_times"`
	CurrentBalance  float64   `json:"current_balance" db:"current_balance"`
	Description     string    `json:"description" db:"description"`
	OrderSN         string    `json:"ordersn" db:"ordersn"`
	Reason          string    `json:"reason" db:"reason"`
	PayOrderList    []string  `json:"pay_order_list"`
	RefundSN        string    `json:"refund_sn" db:"refund_sn"`
	Status          string    `json:"status" db:"status"`
	TransactionFee  float64   `json:"transaction_fee" db:"transaction_fee"`
	TransactionID   int64     `json:"transaction_id" db:"transaction_id"`
	TransactionType string    `json:"transaction_type" db:"transaction_type"`
	WalletType      string    `json:"wallet_type" db:"wallet_type"`
	WithdrawalType  string    `json:"withdrawal_type" db:"withdrawal_type"`
}

type TransactionList2 struct {
	Amount          float64  `json:"amount" db:"amount"`
	BuyerName       string   `json:"buyer_name" db:"buyer_name"`
	CreateTime      int      `json:"create_time" db:"create_time"`
	CreateTimes     string   `json:"create_times" db:"create_times"`
	CurrentBalance  float64  `json:"current_balance" db:"current_balance"`
	Description     string   `json:"description" db:"description"`
	OrderSN         string   `json:"ordersn" db:"ordersn"`
	Reason          string   `json:"reason" db:"reason"`
	PayOrderList    []string `json:"pay_order_list"`
	RefundSN        string   `json:"refund_sn" db:"refund_sn"`
	Status          string   `json:"status" db:"status"`
	TransactionFee  float64  `json:"transaction_fee" db:"transaction_fee"`
	TransactionID   int64    `json:"transaction_id" db:"transaction_id"`
	TransactionType string   `json:"transaction_type" db:"transaction_type"`
	WalletType      string   `json:"wallet_type" db:"wallet_type"`
	WithdrawalType  string   `json:"withdrawal_type" db:"withdrawal_type"`
}

type TransactionListFirebase struct {
	Amount          float64   `json:"Amount" db:"Amount"`
	BuyerName       string    `json:"BuyerName" db:"BuyerName"`
	CreateTime      int       `json:"CreateTime" db:"CreateTime"`
	CreateTimes     time.Time `json:"CreateTimes" db:"CreateTimes"`
	CurrentBalance  float64   `json:"CurrentBalance" db:"CurrentBalance"`
	Description     string    `json:"Description" db:"Description"`
	OrderSN         string    `json:"OrderSN" db:"OrderSN"`
	Reason          string    `json:"Reason" db:"Reason"`
	PayOrderList    []string  `json:"PayOrderList"`
	RefundSN        string    `json:"RefundSN" db:"RefundSN"`
	Status          string    `json:"Status" db:"Status"`
	TransactionFee  float64   `json:"TransactionFee" db:"TransactionFee"`
	TransactionID   int64     `json:"TransactionID" db:"TransactionID"`
	TransactionType string    `json:"TransactionType" db:"TransactionType"`
	WalletType      string    `json:"WalletType" db:"WalletType"`
	WithdrawalType  string    `json:"WithdrawalType" db:"WithdrawalType"`
}

type TestShope struct {
	TotalAmount float64 `json:"total_amount"`
	//SaldoAkhir		float64		`json:"saldo_akhir"`
	TransactionList []TransactionList `json:"transaction_list"`
}

type Header struct {
	ID         string `json:"id""`
	LastUpdate string `json:"last_update"`
}

type StoreData struct {
	Shop_ID    int    `json:"Shop_ID"`
	Partner_ID int    `json:"Partner_ID"`
	Api        string `json:"API"`
	Nama_Toko  string `Nama_Toko`
}

type StoreData2 struct {
	//Shop_ID		int			`json:"Shop_ID"`
	Partner_ID int `json:"Partner_ID"`
	//Api			string		`json:"API"`
	Nama_Toko string `Nama_Toko`
}

type Pagination struct {
	Partner_ID string `json:"PartnerID"`
	Page       int    `json:"Page"`
	Limit      int    `json:"Limit"`
}

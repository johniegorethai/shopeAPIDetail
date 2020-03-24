package orderPelapak

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
	"shopeAPIDetail/pkg/errors"
	"shopeAPIDetail/pkg/firebaseclient"
	"strconv"

	"github.com/jmoiron/sqlx"

	orderPelapakEntity "shopeAPIDetail/internal/entity/orderPelapak"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		c *firestore.Client

		stmt map[string]*sqlx.Stmt
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	getOrderPelapak  = "GetBiayaEkspedisiMaxCount"
	qGetOrderPelapak = "SELECT COUNT(FeeExp_RunningId) FROM FEE_EXP WHERE FeeExp_ActiveYN='Y' AND FeeExp_DataAktifYN='Y'"

	insertMutasiShopee = "insertMutasiShopee"
	qInsertMutasiShopee = "INSERT INTO test_mutasi (amount, buyer_name, create_time, current_balance, description, ordersn, reason, " +
		"refund_sn, status, transaction_fee, transaction_id, transaction_type, wallet_type, withdrawal_type) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
)

var (
	readStmt = []statement{
		{getOrderPelapak, qGetOrderPelapak},
		{insertMutasiShopee, qInsertMutasiShopee},
	}
)

// New ...
func New(db *sqlx.DB, fc *firebaseclient.Client) Data {
	d := Data{
		db: db,
		c: fc.Client,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetAllOrderPelapak ...
func (d Data) GetAllOrderPelapak(ctx context.Context) ([]orderPelapakEntity.OrderPelapak, error) {
	var (
		orderPelapak  orderPelapakEntity.OrderPelapak
		orderPelapaks []orderPelapakEntity.OrderPelapak
		err           error
	)

	rows, err := d.stmt[getOrderPelapak].QueryxContext(ctx)

	for rows.Next() {
		if err := rows.StructScan(&orderPelapak); err != nil {
			return orderPelapaks, errors.Wrap(err, "[DATA][GetAllOrderPelapak] ")
		}
		orderPelapaks = append(orderPelapaks, orderPelapak)
	}
	return orderPelapaks, err
}

func (d Data) InsertMutasiShopee(ctx context.Context, mutasi orderPelapakEntity.TransactionList ) (orderPelapakEntity.TransactionList , error) {
	var (
		err         error
	)
	if _, err := d.stmt[insertMutasiShopee].ExecContext(ctx,
		mutasi.Amount,
		mutasi.BuyerName,
		mutasi.CreateTime,
		mutasi.CurrentBalance,
		mutasi.Description,
		mutasi.OrderSN,
		mutasi.Reason,
		mutasi.RefundSN,
		mutasi.Status,
		mutasi.TransactionFee,
		mutasi.TransactionID,
		mutasi.TransactionType,
		mutasi.WalletType,
		mutasi.WithdrawalType,
	); err != nil {
		fmt.Println(err)
		return mutasi, errors.Wrap(err, "[DATA][InsertMutasiShopee] ")
	}
	return mutasi, err
}

func (d Data) InsertMutasiFirebase(ctx context.Context, transList orderPelapakEntity.TransactionList ) (orderPelapakEntity.TransactionList , error) {

	var err error//tCount, err := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Get(ctx)
	//_, err = d.c.Collection("mutasi_shopee").Doc("data_mutasi").Update(ctx, []firestore.Update{
	//	{Path: "count", Value: firestore.Increment(1)},
	//})
	//if err != nil {
	//	return transList, errors.Wrap(err, "[DATA][SimpanRepeatOrderBaru] Failed updating count of repeat order!")
	//}
	//
	//data := tCount.Data()
	//incre := strconv.Itoa(int(data["count"].(int64) + 1))
	//
	_, err = d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc("1").
		Collection("mutasi_detail").Doc(strconv.Itoa(int(transList.TransactionID))).Set(ctx,transList)

	if err != nil {
		return transList, errors.Wrap(err, "[DATA][SimpanRepeatOrderBaru] Failed setting purchase order!")
	}

	return transList, err
}

//func (d Data) InsertMutasiShopee2(ctx context.Context, mutasi orderPelapakEntity.TransactionList ) ([]orderPelapakEntity.TransactionList , error) {
//	var (
//		err         error
//
//	)
//	if _, err := d.stmt[insertMutasiShopee].ExecContext(ctx,
//		mutasi.Amount,
//		mutasi.BuyerName,
//		mutasi.CreateTime,
//		mutasi.CurrentBalance,
//		mutasi.Description,
//		mutasi.OrderSN,
//		mutasi.Reason,
//		mutasi.RefundSN,
//		mutasi.Status,
//		mutasi.TransactionFee,
//		mutasi.TransactionID,
//		mutasi.TransactionType,
//		mutasi.WalletType,
//		mutasi.WithdrawalType,
//	); err != nil {
//		fmt.Println(err)
//		return mutasi, errors.Wrap(err, "[DATA][InsertMutasiShopee] ")
//	}
//	return mutasi, err
//}

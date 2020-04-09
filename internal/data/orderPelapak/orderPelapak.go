package orderPelapak

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/iterator"
	"log"
	"shopeAPIDetail/pkg/errors"
	"shopeAPIDetail/pkg/firebaseclient"
	"strconv"
	"time"

	orderPelapakEntity "shopeAPIDetail/internal/entity/orderPelapak"
)

type (
	// Data ...
	Data struct {
		db *sqlx.DB
		c  *firestore.Client

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

	insertMutasiShopee  = "insertMutasiShopee"
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
		c:  fc.Client,
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

func (d Data) GetLastMutasiFirebase(ctx context.Context, partnerid int) (orderPelapakEntity.TransactionListFirebase, error) {
	var (
		tran orderPelapakEntity.TransactionListFirebase
		//header orderPelapakEntity.Header
		//trans []orderPelapakEntity.TransactionList
		err error
	)

	partID := strconv.Itoa(partnerid)

	tCount, err := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).Get(ctx)

	data := tCount.Data()
	ID := strconv.Itoa(int(data["Count"].(int64)))
	id := partID + ID

	fmt.Println(partID)
	fmt.Println(id)
	//	iter := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).OrderBy("LastUpdate", firestore.Desc).Limit(1).Documents(ctx)
	iter := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).Collection(id).
		OrderBy("CreateTimes", firestore.Desc).Limit(1).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return tran, errors.Wrap(err, "[DATA][SelectLastInput] Failed to iterate documents!")
		}
		err = doc.DataTo(&tran)
		fmt.Println(tran)
	}

	return tran, err
}

func (d Data) GetShopeeStoreData(ctx context.Context) ([]orderPelapakEntity.StoreData, error) {
	var (
		store  orderPelapakEntity.StoreData
		stores []orderPelapakEntity.StoreData
		err    error
	)

	iter := d.c.Collection("mutasi_shopee/data_shopee/shop_id").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return stores, errors.Wrap(err, "[DATA][GetShopeeStoreData] Failed to iterate documents!")
		}
		err = doc.DataTo(&store)
		if err != nil {
			return stores, errors.Wrap(err, "[DATA][GetShopeeStoreData] Failed to populate struct!")
		}
		stores = append(stores, store)
	}
	return stores, err
}

func (d Data) GetShopeeStoreData2(ctx context.Context) ([]orderPelapakEntity.StoreData2, error) {
	var (
		store  orderPelapakEntity.StoreData2
		stores []orderPelapakEntity.StoreData2
		err    error
	)

	iter := d.c.Collection("mutasi_shopee/data_shopee/shop_id").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return stores, errors.Wrap(err, "[DATA][GetShopeeStoreData] Failed to iterate documents!")
		}
		err = doc.DataTo(&store)
		if err != nil {
			return stores, errors.Wrap(err, "[DATA][GetShopeeStoreData] Failed to populate struct!")
		}
		stores = append(stores, store)
	}
	return stores, err
}

func (d Data) GetMutasiFirebase(ctx context.Context, pagination orderPelapakEntity.Pagination) ([]orderPelapakEntity.TransactionList, error) {
	var (
		tran    orderPelapakEntity.TransactionList
		trans   []orderPelapakEntity.TransactionList
		iter    *firestore.DocumentIterator
		lastDoc *firestore.DocumentSnapshot
		err     error
	)

	fmt.Println(pagination.Partner_ID)

	tCount, err := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(pagination.Partner_ID).Get(ctx)

	data := tCount.Data()
	ID := strconv.Itoa(int(data["Count"].(int64)))
	id := pagination.Partner_ID + ID

	if pagination.Page == 1 {
		iter = d.c.Collection("mutasi_shopee/data_mutasi/ID").Doc(pagination.Partner_ID).Collection(id).
			OrderBy("CreateTimes", firestore.Desc).Limit(pagination.Limit).Documents(ctx)
	} else {

		iter2 := d.c.Collection("mutasi_shopee/data_mutasi/ID").Doc(pagination.Partner_ID).Collection(id).
			OrderBy("CreateTimes", firestore.Desc).Limit((pagination.Page - 1) * pagination.Limit).Documents(ctx)

		doc, _ := iter2.GetAll()
		// Get the last document.
		lastDoc = doc[len(doc)-1]

		iter = d.c.Collection("mutasi_shopee/data_mutasi/ID").Doc(pagination.Partner_ID).Collection(id).
			StartAfter(lastDoc).
			OrderBy("CreateTimes", firestore.Desc).Limit(pagination.Limit).Documents(ctx)
		//
	}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return trans, errors.Wrap(err, "[DATA][GetShopeeStoreData] Failed to iterate documents!")
		}
		err = doc.DataTo(&tran)
		if err != nil {
			return trans, errors.Wrap(err, "[DATA][GetShopeeStoreData] Failed to populate struct!")
		}
		trans = append(trans, tran)
	}
	return trans, err
}

func (d Data) InsertMutasiShopee(ctx context.Context, mutasi orderPelapakEntity.TransactionList) (orderPelapakEntity.TransactionList, error) {
	var (
		err error
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

func (d Data) InsertMutasiFirebaseHeader(ctx context.Context, partnerid int) error {

	var err error

	partID := strconv.Itoa(partnerid)

	tCount, err := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).Get(ctx)
	_, err = d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).Update(ctx, []firestore.Update{
		{Path: "Count", Value: firestore.Increment(1)},
	})
	if err != nil {
		return errors.Wrap(err, "[DATA][SimpanRepeatOrderBaru] Failed updating count of repeat order!")
	}

	data := tCount.Data()
	incre := int(data["Count"].(int64) + 1)

	_, err = d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).Set(ctx, map[string]interface{}{
		"ID":         partID + strconv.Itoa(incre),
		"LastUpdate": time.Now().Format("2006-01-02 15:04:05"),
		"Count":      incre,
	})

	if err != nil {
		return errors.Wrap(err, "[DATA][SimpanRepeatOrderBaru] Failed setting purchase order!")
	}

	return err
}

func (d Data) InsertMutasiFirebase(ctx context.Context, transList orderPelapakEntity.TransactionList, partnerid int) (orderPelapakEntity.TransactionList, error) {

	var err error

	partID := strconv.Itoa(partnerid)

	tCount, err := d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).Get(ctx)

	if err != nil {
		return transList, errors.Wrap(err, "[DATA][SimpanRepeatOrderBaru] Failed updating count of repeat order!")
	}

	data := tCount.Data()
	ID := strconv.Itoa(int(data["Count"].(int64)))
	id := partID + ID

	_, err = d.c.Collection("mutasi_shopee").Doc("data_mutasi").Collection("ID").Doc(partID).
		Collection(id).Doc(strconv.Itoa(int(transList.TransactionID))).Set(ctx, transList)

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

package orderPelapak

import (
	"context"
	"fmt"
	_ "github.com/robfig/cron/v3"
	"shopeAPIDetail/pkg/errors"
	"time"

	orderPelapakEntity "shopeAPIDetail/internal/entity/orderPelapak"
	userEntity "shopeAPIDetail/internal/entity/shope"
)

// Data ...
type Data interface {
	GetAllOrderPelapak(ctx context.Context) ([]orderPelapakEntity.OrderPelapak, error)
	InsertMutasiShopee(ctx context.Context, mutasi orderPelapakEntity.TransactionList ) (orderPelapakEntity.TransactionList , error)
	InsertMutasiFirebase(ctx context.Context, transList orderPelapakEntity.TransactionList ) (orderPelapakEntity.TransactionList , error)
}

// GudangData ...
type ShopeData interface {
	GetDataShope(ctx context.Context) (interface{}, error)
	GetDataShope2(ctx context.Context) (interface{}, error)
	//GetDataShopeTransaction(ctx context.Context, shop userEntity.Shope) (interface{}, error)
	GetDataShopeTransaction(ctx context.Context, shop userEntity.Shope) (userEntity.MutasiShopee, error)
}

// Service ...
type Service struct {
	data  Data
	shope ShopeData
}

// New ...
func New(orderPelapakData Data, shopeData ShopeData) Service {
	return Service{
		data:  orderPelapakData,
		shope: shopeData,
	}
}

// GetAllSkeletons ...
func (s Service) GetAllOrderPelapak(ctx context.Context) (interface{}, error) {
	var (
		//result   []skeletonEntity.OrderPelapak
		err error
		// shope   shopeEntity.Shope
		result1 interface{}
	)

	result1, err = s.shope.GetDataShope(context.Background())

	return result1, err
}

func (s Service) GetAllOrderPelapak2(ctx context.Context) (interface{}, error) {
	var (
		//result   []skeletonEntity.OrderPelapak
		err error
		// shope   shopeEntity.Shope
		result1 interface{}
	)

	result1, err = s.shope.GetDataShope2(context.Background())

	return result1, err
}

func (s Service) GetShopeeTransaction2(ctx context.Context, shop userEntity.Shope) (interface{}, error) {
	var (
		//result   []skeletonEntity.OrderPelapak
		err error
		// shope   shopeEntity.Shope
		result1 interface{}
	)
	shop.Timestamp = time.Now().Unix()

	result1, err = s.shope.GetDataShopeTransaction(context.Background(), shop)

	fmt.Println()

	return result1, err
}

func (s Service) GetShopeeTransaction(ctx context.Context, shop userEntity.Shope) ([]userEntity.MutasiShopee, error) {
	var (
		//result   []skeletonEntity.OrderPelapak
		err error
		// shope   shopeEntity.Shope
		result []userEntity.MutasiShopee
		i int64
		indexTrans int
		indeksY []int
		indeksZ []int
//		indeks int
	)


	shop.Timestamp = time.Now().Unix()
	fmt.Println(shop)
	result1, err := s.shope.GetDataShopeTransaction(context.Background(), shop)

	if result1.HasMore == true {
		for i = 1 ; i <= 1000 ; i+=100 {
			shop.PaginationOffset = i
			shop.Timestamp = time.Now().Unix()
			result2, _ := s.shope.GetDataShopeTransaction(context.Background(), shop)

			result = append(result, result2)

			if result2.HasMore == false {
				if len(result1.TransactionList) < 101 {
					break
				}
			}
		}
	}else {
		if len(result1.TransactionList) > 0{
			result = append(result, result1)
		}
	}

	indexTrans = len(result)
	fmt.Println(len(result))

	for z := 0 ; z< indexTrans ; z++ {
		fmt.Println(z)
		for y := 0; y < len(result[z].TransactionList); y++ {

			if result[z].TransactionList[y].CurrentBalance == 0 {
				indeksY = append(indeksY, y)
				indeksZ = append(indeksZ, z)
				//indeks = y
				//fmt.Println("index 0 terakhir ",indeks)
				//fmt.Println(result[z].TransactionList[indeks+1].BuyerName)
			}
		}
	}

	fmt.Println(indeksY)
	fmt.Println(indeksZ)



	return result, err

	//{
	//	"partner_id":842234,
	//	"shopid":121791179,
	//	"pagination_offset": 1,
	//	"pagination_entries_per_page":100,
	//	"create_time_from":1584428400,
	//	"create_time_to":1584457200
	//}
}

//func (s Service) InsertMutasiShopee (ctx context.Context, shop userEntity.Shope)([]orderPelapakEntity.TransactionList, error){
//	var (
//		err error
//		//result1 userEntity.MutasiShopee
//		result []orderPelapakEntity.TransactionList
//		totalAmount	float64
//	)
//
//	shop.Timestamp = time.Now().Unix()
//	result1, err := s.shope.GetDataShopeTransaction(context.Background(), shop)
//
//	fmt.Println(result1)
//
//	fmt.Println(result1.TransactionList)
//
//	for i:= 0 ; i < len(result1.TransactionList)  ; i++ {
//		res, _ := s.data.InsertMutasiShopee(ctx, orderPelapakEntity.TransactionList(result1.TransactionList[i]))
//		if err != nil{
//			return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
//		}
//		totalAmount += res.Amount
//		result = append(result, res)
//	}
//
//	//result1.TotalAmount = totalAmount
//	return result, err
//}

func (s Service) InsertMutasiShopee (ctx context.Context, shop userEntity.Shope)(orderPelapakEntity.TestShope, error){
	var (
		err error
		//result1 userEntity.MutasiShopee
		result orderPelapakEntity.TestShope
		results []userEntity.MutasiShopee
		//results1 []userEntity.MutasiShopee
		totalAmount	float64
		indexTrans int
		i int64
		//indexPertamas int
		//indexTerakhir int

	)

	shop.Timestamp = time.Now().Unix()
	fmt.Println(shop)
	result1, err := s.shope.GetDataShopeTransaction(context.Background(), shop)

	if result1.HasMore == true {
		for i = 1 ; i <= 1000 ; i+=100 {
			shop.PaginationOffset = i
			shop.Timestamp = time.Now().Unix()
			result2, _ := s.shope.GetDataShopeTransaction(context.Background(), shop)

			results = append(results, result2)

			if result2.HasMore == false {
				if len(result1.TransactionList) < 101 {
					break
				}
			}
		}
	}else {
		if len(result1.TransactionList) > 0{
			results = append(results, result1)
		}
	}
	//results hasil array dari mutasiShopee

	fmt.Println(len(results))

	for x := 0 ; x < len(results) ; x++ {
		indexTrans += x
	}

	for z := 0 ; z<= indexTrans ; z++ {
		fmt.Println(z)
		for y := 0; y < len(results[z].TransactionList); y++ {

			res, _ := s.data.InsertMutasiFirebase(ctx, orderPelapakEntity.TransactionList(results[z].TransactionList[y]))

			if err != nil {
				return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
			}
			totalAmount += res.Amount
			result.TransactionList = append(result.TransactionList, res)
		}
	}

	//BUAT INSERT
	//for z := 0 ; z<= indexTrans ; z++ {
	//	fmt.Println(z)
	//	for y := 0; y < len(results[z].TransactionList); y++ {
	//
	//		if results[z].TransactionList[y].CurrentBalance == 0 {
	//			fmt.Println("kena")
	//		}
	//
	//		//res, _ := s.data.InsertMutasiShopee(ctx, orderPelapakEntity.TransactionList(results[z].TransactionList[y]))
	//		//
	//		//if err != nil {
	//		//	return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
	//		//}
	//		//totalAmount += res.Amount
	//		//result.TransactionList = append(result.TransactionList, res)
	//	}
	//}

	//fmt.Println("totalnya ", totalAmount)

	result.TotalAmount = totalAmount
	return result, err
}
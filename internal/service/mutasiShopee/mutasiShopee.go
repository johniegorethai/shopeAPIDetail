package mutasiShopee

import (
	"context"
	"fmt"
	_ "github.com/robfig/cron/v3"
	"log"
	"shopeAPIDetail/pkg/errors"
	"time"

	orderPelapakEntity "shopeAPIDetail/internal/entity/mutasiShopee"
	userEntity "shopeAPIDetail/internal/entity/shope"
)

// Data ...
type Data interface {
	GetAllOrderPelapak(ctx context.Context) ([]orderPelapakEntity.OrderPelapak, error)
	InsertMutasiShopee(ctx context.Context, mutasi orderPelapakEntity.TransactionList) (orderPelapakEntity.TransactionList, error)
	InsertMutasiFirebase(ctx context.Context, transList orderPelapakEntity.TransactionList, partnerid int) (orderPelapakEntity.TransactionList, error)
	InsertMutasiFirebaseHeader(ctx context.Context, partnerid int) error
	GetLastMutasiFirebase(ctx context.Context, partnerid int) (orderPelapakEntity.TransactionListFirebase, error)
	GetShopeeStoreData2(ctx context.Context) ([]orderPelapakEntity.StoreData2, error)
	GetShopeeStoreData(ctx context.Context) ([]orderPelapakEntity.StoreData, error)
	GetMutasiFirebase(ctx context.Context, pagination orderPelapakEntity.Pagination) ([]orderPelapakEntity.TransactionList, error)
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

func (s Service) GetShopeeStoreData2(ctx context.Context) ([]orderPelapakEntity.StoreData2, error) {
	var (
		err error
	)
	result1, err := s.data.GetShopeeStoreData2(context.Background())

	//
	//for _, data := range result1{
	//	fmt.Println("shopid nya", data.Shop_ID)
	//}
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
		result     []userEntity.MutasiShopee
		i          int64
		indexTrans int
		indeksY    []int
		indeksZ    []int
		//		indeks int
	)

	shop.Timestamp = time.Now().Unix()
	fmt.Println(shop)
	result1, err := s.shope.GetDataShopeTransaction(context.Background(), shop)

	if result1.HasMore == true {
		for i = 0; i <= 2000; i += 99 {
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
	} else {
		if len(result1.TransactionList) > 0 {
			result = append(result, result1)
		}
	}

	indexTrans = len(result)
	fmt.Println(len(result))

	for z := 0; z < indexTrans; z++ {
		fmt.Println(z)
		for y := 0; y < len(result[z].TransactionList); y++ {

			if result[z].TransactionList[y].CurrentBalance == 0 {
				indeksY = append(indeksY, y)
				indeksZ = append(indeksZ, z)
			}
		}
	}

	//z itu halaman
	//y itu index ke-?

	// tm := 1584428400
	//times := time.Unix(int64(tm),0)
	//times.UTC()
	//fmt.Println("times ",times)
	fmt.Println("Y ", indeksY)
	fmt.Println("X ", indeksZ)

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

func (s Service) InsertMutasiShopee(ctx context.Context, shop userEntity.Shope) (orderPelapakEntity.TestShope, error) {
	var (
		err error
		//result1 userEntity.MutasiShopee
		result  orderPelapakEntity.TestShope
		results []userEntity.MutasiShopee
		//results1 []userEntity.MutasiShopee
		totalAmount float64
		//indexTrans int
		i int64
		//index1 int
		//index2 int
		indeksY []int
		indeksZ []int
	)

	shop.Timestamp = time.Now().Unix()
	fmt.Println(shop)
	result1, err := s.shope.GetDataShopeTransaction(context.Background(), shop)

	if result1.HasMore == true {
		for i = 0; i <= 2000; i += 99 {
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
	} else {
		if len(result1.TransactionList) > 0 {
			results = append(results, result1)
		}
	}
	//results hasil array dari mutasiShopee

	fmt.Println(len(results))

	for z := 0; z < len(results); z++ {
		fmt.Println(z)
		for y := 0; y < len(results[z].TransactionList); y++ {

			if results[z].TransactionList[y].CurrentBalance == 0 {
				indeksY = append(indeksY, y)
				indeksZ = append(indeksZ, z)
			}
		}
	}
	s.data.InsertMutasiFirebaseHeader(ctx, shop.PartnerID)

	for z := indeksZ[0]; z < len(results); z++ {
		if len(indeksZ) == 1 {
			break
		}

		if z <= 0 {
			for y := indeksY[0]; y < len(results[z].TransactionList); y++ {
				tm := results[z].TransactionList[y].CreateTime - 25200
				times := time.Unix(int64(tm), 0)
				times.Format("2006-01-02 15:04:05")
				fmt.Println(y)
				results[z].TransactionList[y].CreateTimes = times

				res, _ := s.data.InsertMutasiFirebase(ctx, orderPelapakEntity.TransactionList(results[z].TransactionList[y]), shop.PartnerID)

				if err != nil {
					return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
				}
				totalAmount += res.Amount
				result.TransactionList = append(result.TransactionList, res)
			}
		} else if z > 0 {
			for y := 0; y < len(results[z].TransactionList); y++ {
				if z == indeksZ[1] && y == indeksY[1]+1 {
					break
				} else {
					tm := results[z].TransactionList[y].CreateTime - 25200
					times := time.Unix(int64(tm), 0)
					times.Format("2006-01-02 15:04:05")
					fmt.Println(y)
					results[z].TransactionList[y].CreateTimes = times

					res, _ := s.data.InsertMutasiFirebase(ctx, orderPelapakEntity.TransactionList(results[z].TransactionList[y]), shop.PartnerID)

					if err != nil {
						return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
					}
					totalAmount += res.Amount
					result.TransactionList = append(result.TransactionList, res)
				}
			}
		}

	}

	result.TotalAmount = totalAmount
	return result, err
}

func (s Service) InsertMutasiShopeeScheduled(ctx context.Context) (orderPelapakEntity.TestShope, error) {
	var (
		err     error
		result  orderPelapakEntity.TestShope
		results []userEntity.MutasiShopee
		i       int64
	)

	store, err := s.data.GetShopeeStoreData(context.Background())

	for _, data := range store {

		resultx, err := s.data.GetLastMutasiFirebase(ctx, data.Partner_ID)
		var indeksY []int
		var indeksZ []int

		indeksY = nil
		indeksZ = nil

		results = nil

		//fmt.Println(resultx.CreateTimes.UTC())
		//fmt.Println(resultx.CreateTimes.Unix())
		//fmt.Println(resultx.CreateTimes.Unix() + 25200)

		shop := userEntity.Shope{
			PartnerID:                data.Partner_ID, // ==> Di Sini untuk Partner ID
			ShopID:                   data.Shop_ID,    // ==> Di Sini Untuk SHOP ID
			Api:                      data.Api,
			PaginationOffset:         0,
			PaginationEntriesPerPage: 100,
			CreateTimeFrom:           resultx.CreateTimes.Unix(),
			//CreateTimeFrom: 1586577600,
			CreateTimeTo: time.Now().Unix(),
		}

		shop.Timestamp = time.Now().Unix()
		//shop.CreateTimeTo = time.Now().Unix()
		fmt.Println(shop)
		result1, _ := s.shope.GetDataShopeTransaction(context.Background(), shop)

		if result1.HasMore == true {
			for i = 0; i <= 5000; i += 99 {
				shop.PaginationOffset = i
				shop.Timestamp = time.Now().Unix()
				//shop.CreateTimeTo = time.Now().Unix()
				result2, _ := s.shope.GetDataShopeTransaction(context.Background(), shop)

				results = append(results, result2)

				if result2.HasMore == false {
					if len(result1.TransactionList) < 101 {
						break
					}
				}
			}
		} else {
			if len(result1.TransactionList) > 0 {
				results = append(results, result1)
			}
		}
		//results hasil array dari mutasiShopee

		fmt.Println("aa", len(results))

		for z := 0; z < len(results); z++ {
			fmt.Println(z)
			for y := 0; y < len(results[z].TransactionList); y++ {

				if results[z].TransactionList[y].CurrentBalance == 0 {

					indeksY = append(indeksY, y)
					indeksZ = append(indeksZ, z)
				}
			}
		}

		if len(indeksZ) > 1 {
			s.data.InsertMutasiFirebaseHeader(ctx, shop.PartnerID)

			fmt.Println("X nya ", indeksZ)
			fmt.Println("Y nya ", indeksY)

			//fmt.Println("len ", len(results))

			for z := indeksZ[0]; z <= indeksZ[1]; z++ {
				if len(indeksZ) == 1 {
					break
				}
				//masih bermsalah di sini
				fmt.Println("bb", len(results[z].TransactionList))

				if z <= 0 {
					for y := indeksY[0]; y < len(results[z].TransactionList); y++ {
						tm := results[z].TransactionList[y].CreateTime
						times := time.Unix(int64(tm), 0)
						times.Format("2006-01-02 15:04:05")
						fmt.Print(y)
						results[z].TransactionList[y].CreateTimes = times

						res, _ := s.data.InsertMutasiFirebase(ctx, orderPelapakEntity.TransactionList(results[z].TransactionList[y]), shop.PartnerID)

						if err != nil {
							return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
						}
						result.TransactionList = append(result.TransactionList, res)
					}
				} else if z > 0 {
					for y := 0; y < len(results[z].TransactionList); y++ {
						if z == indeksZ[1] && y == indeksY[1]+1 {
							fmt.Println("break")
							break
						} else {
							tm := results[z].TransactionList[y].CreateTime
							times := time.Unix(int64(tm), 0)
							times.Format("2006-01-02 15:04:05")
							fmt.Print(y)
							results[z].TransactionList[y].CreateTimes = times

							res, _ := s.data.InsertMutasiFirebase(ctx, orderPelapakEntity.TransactionList(results[z].TransactionList[y]), shop.PartnerID)

							if err != nil {
								return result, errors.Wrap(err, "[SERVICE][InsertMutasiShopee]")
							}
							result.TransactionList = append(result.TransactionList, res)
						}
					}
				}
			}
		}
	}
	return result, err
}

func (s Service) GetLastMutasiFirebase(ctx context.Context, partnerid int) (orderPelapakEntity.TransactionListFirebase, error) {
	result, err := s.data.GetLastMutasiFirebase(ctx, partnerid)

	if err != nil {
		log.Fatal(err)
	}
	result.CreateTimes.UTC()

	//fmt.Println(result.CreateTimes.UTC())
	return result, err
}

func (s Service) GetMutasiFirebase(ctx context.Context, pagination orderPelapakEntity.Pagination) ([]orderPelapakEntity.TransactionList, error) {
	result, err := s.data.GetMutasiFirebase(ctx, pagination)

	for i, _ := range result {

		tm := result[i].CreateTime
		times := time.Unix(int64(tm), 0)
		times.Format("2006-01-02 15:04:05")
		fmt.Println(times)
		result[i].CreateTimes = times
	}
	//result[0].CreateTimes.Format("2006-01-02 15:04:05")
	//fmt.Println(result[0].CreateTimes)

	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

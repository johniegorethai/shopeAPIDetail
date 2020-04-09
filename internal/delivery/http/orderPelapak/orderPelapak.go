package orderPelapak

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	orderEntity "shopeAPIDetail/internal/entity/orderPelapak"
	userEntity "shopeAPIDetail/internal/entity/shope"
	"shopeAPIDetail/pkg/response"
)

// ISkeletonSvc is an interface to Skeleton Service
type IOrderPelapakSvc interface {
	GetAllOrderPelapak(ctx context.Context) (interface{}, error)
	GetAllOrderPelapak2(ctx context.Context) (interface{}, error)
	GetShopeeTransaction2(ctx context.Context, shop userEntity.Shope) (interface{}, error)
	GetShopeeTransaction(ctx context.Context, shop userEntity.Shope) ([]userEntity.MutasiShopee, error)
	InsertMutasiShopee(ctx context.Context, shop userEntity.Shope) (orderEntity.TestShope, error)
	GetLastMutasiFirebase(ctx context.Context, partnerid int) (orderEntity.TransactionListFirebase, error)
	GetShopeeStoreData2(ctx context.Context) ([]orderEntity.StoreData2, error)
	InsertMutasiShopeeScheduled(ctx context.Context) (orderEntity.TestShope, error)
	GetMutasiFirebase(ctx context.Context, pagination orderEntity.Pagination) ([]orderEntity.TransactionList, error)
}

type (
	// Handler ...
	Handler struct {
		orderPelapakSvc IOrderPelapakSvc
	}
)

// New for bridging product handler initialization
func New(is IOrderPelapakSvc) *Handler {
	return &Handler{
		orderPelapakSvc: is,
	}
}

// SkeletonHandler return user data
func (h *Handler) OrderPelapakHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp       *response.Response
		metadata   interface{}
		err        error
		shop       userEntity.Shope
		pagination orderEntity.Pagination
		//result 	 []orderPelapakEntity.OrderPelapak
		errRes  response.Error
		result1 interface{}
		_type   string
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	rBody, _ := ioutil.ReadAll(r.Body)

	// Check if request method is GET
	if r.Method == http.MethodGet {

		if _, x := r.URL.Query()["type"]; x {
			_type = r.FormValue("type")
			if _type == "getLastMutasiFirebase" {
				var partnerid int
				partnerid, err = strconv.Atoi(r.FormValue("partnerid"))
				result1, err = h.orderPelapakSvc.GetLastMutasiFirebase(context.Background(), partnerid)
			} else if _type == "getStoreData" {
				result1, err = h.orderPelapakSvc.GetShopeeStoreData2(context.Background())
			}
		}
	} else if r.Method == http.MethodPost {
		json.Unmarshal(rBody, &shop)
		json.Unmarshal(rBody, &pagination)
		if _, x := r.URL.Query()["type"]; x {
			_type = r.FormValue("type")
			if _type == "orderPelapak" {
				fmt.Println("Tes")
				result1, err = h.orderPelapakSvc.GetAllOrderPelapak(context.Background())
				fmt.Println("resutlnya", result1)
			} else if _type == "getShop" {
				result1, err = h.orderPelapakSvc.GetAllOrderPelapak2(context.Background())
				fmt.Println(result1)
			} else if _type == "getShopTrans" {
				//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
				result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
			} else if _type == "getShopTrans2" {
				//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
				result1, err = h.orderPelapakSvc.GetShopeeTransaction2(context.Background(), shop)
			} else if _type == "insertShop" {
				//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
				result1, err = h.orderPelapakSvc.InsertMutasiShopee(context.Background(), shop)
			} else if _type == "insertShopScheduled" {
				//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
				result1, err = h.orderPelapakSvc.InsertMutasiShopeeScheduled(context.Background())
			} else if _type == "getMutasiShopee" {
				result1, err = h.orderPelapakSvc.GetMutasiFirebase(context.Background(), pagination)
			}
		}
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	resp.Data = result1
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}

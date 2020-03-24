	package orderPelapak

	import (
		"context"
		"encoding/json"
		"fmt"
		"io/ioutil"
		"log"
		"net/http"
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
	InsertMutasiShopee (ctx context.Context, shop userEntity.Shope)(orderEntity.TestShope, error)}

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
		resp     *response.Response
		metadata interface{}
		err      error
		shop	userEntity.Shope
		//result 	 []orderPelapakEntity.OrderPelapak
		errRes  response.Error
		result1 interface{}
		_type	string
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	rBody, _ := ioutil.ReadAll(r.Body)

	// Check if request method is GET
	if r.Method == http.MethodGet {
		// Check if page and length exists in URL parameters
		result1, err = h.orderPelapakSvc.GetAllOrderPelapak(context.Background())
	} else if r.Method == http.MethodPost {
		json.Unmarshal(rBody, &shop)
		if _, x := r.URL.Query()["type"]; x {
			_type = r.FormValue("type")
			if _type == "orderPelapak" {
				fmt.Println("Tes")
				result1, err = h.orderPelapakSvc.GetAllOrderPelapak(context.Background())
				fmt.Println("resutlnya",result1)
			} else if _type == "getShop" {
				result1, err = h.orderPelapakSvc.GetAllOrderPelapak2(context.Background())
				fmt.Println(result1)
			}else if _type == "getShopTrans" {
				//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
				result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
			}else if _type == "getShopTrans2" {
				//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
				result1, err = h.orderPelapakSvc.GetShopeeTransaction2(context.Background(), shop)
			}else if _type == "insertShop" {
			//result1, err = h.orderPelapakSvc.GetShopeeTransaction(context.Background(), shop)
			result1, err = h.orderPelapakSvc.InsertMutasiShopee(context.Background(), shop)
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

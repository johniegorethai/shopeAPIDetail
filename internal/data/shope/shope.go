package shope

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	shopeEntity "shopeAPIDetail/internal/entity/shope"
	"shopeAPIDetail/pkg/errors"
	"shopeAPIDetail/pkg/httpclient"
	"time"
)

// Data object for data layer needs
type Data struct {
	client    *httpclient.Client
	shopeURL  string
	shopeURL1 string
	shopeURL2 string
	authURL   string
}

// New return an ads resource
func New(client *httpclient.Client, shopeURL string, shopeURL1 string, shopeURL2 string) Data {
	d := Data{
		client:    client,
		shopeURL:  shopeURL,
		shopeURL1: shopeURL1,
		shopeURL2: shopeURL2,
	}

	return d
}

func (d Data) GenerateSignature(partnerKey string, json1 shopeEntity.Shope, shopeURL string) (result string) {

	json2, err := json.Marshal(json1)
	if err != nil {
		return ""
	}
	testt := string(json2)
	baseStr := shopeURL + "|" + testt
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write([]byte(baseStr))
	sha := hex.EncodeToString(mac.Sum(nil))
	return sha
}

// GetStockGudang will return Stock list
func (d Data) GetDataShope(ctx context.Context) (interface{}, error) {
	var (
		resp     shopeEntity.Response
		result   interface{}
		endpoint = d.shopeURL
		err      error
		headers  = make(http.Header)
		//autho string
	)

	json1 := shopeEntity.Shope{

		PartnerID:       844464,    // ==> Di Sini untuk Partner ID
		ShopID:          114789399, // ==> Di Sini Untuk SHOP ID
		Timestamp:       time.Now().Unix(),
		ReleaseTimeFrom: "2020-01-16",
		ReleaseTimeTo:   "2020-01-16",
	}
	json2, err := json.Marshal(json1)

	// ===> Di bawah ini untuk Key <===//
	autho := d.GenerateSignature("2a2437155def6d2885e7ff7a14069d94d368199576bf4cedb5f42a0950ae7bb7", json1, endpoint)
	fmt.Println(autho)

	headers.Set("Authorization", autho)
	headers.Set("Content-Type", "application/json")
	fmt.Println(headers)
	fmt.Println(json1)

	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetDataShope]")
	}

	fmt.Println(endpoint)
	fmt.Println(&resp)

	_, err = d.client.PostJSON(ctx, endpoint, headers, json2, &resp)
	fmt.Printf("%+v\n", resp)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetStockGudang]")
	}

	return result, nil
}

func (d Data) GetDataShope2(ctx context.Context) (interface{}, error) {
	var (
		resp     interface{}
		result   interface{}
		endpoint = d.shopeURL1
		err      error
		headers  = make(http.Header)
		//autho string
	)

	json1 := shopeEntity.Shope{

		PartnerID: 844464,    // ==> Di Sini untuk Partner ID
		ShopID:    114789399, // ==> Di Sini Untuk SHOP ID
		Timestamp: time.Now().Unix(),
	}
	json2, err := json.Marshal(json1)

	// ===> Di bawah ini untuk Key <===//
	autho := d.GenerateSignature("2a2437155def6d2885e7ff7a14069d94d368199576bf4cedb5f42a0950ae7bb7", json1, endpoint)
	fmt.Println(autho)
	fmt.Println(json1)

	headers.Set("Authorization", autho)
	headers.Set("Content-Type", "application/json")
	fmt.Println(headers)

	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetDataShope]")
	}

	fmt.Println(endpoint)
	fmt.Println(&resp)

	_, err = d.client.PostJSON(ctx, endpoint, headers, json2, &resp)
	fmt.Printf("%+v\n", resp)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetStockGudang]")
	}

	return result, nil
}

//func (d Data) GetDataShopeTransaction(ctx context.Context, shop shopeEntity.Shope) (interface{}, error) {
//	var (
//		resp     interface{}
//		result   interface{}
//		endpoint = d.shopeURL2
//		err      error
//		headers  = make(http.Header)
//		//autho string
//	)
//
//	//json1 := shopeEntity.Shope{
//	//	PartnerID:       844098, // ==> Di Sini untuk Partner ID
//	//	ShopID:          175375997, // ==> Di Sini Untuk SHOP ID
//	//	Timestamp:       time.Now().Unix(),
//	//	PaginationOffset: 1,
//	//	PaginationEntriesPerPage: 10,
//	//	CreateTimeFrom:	1583924400,
//	//	CreateTimeTo:	1583971200,
//	//}
//	json2, err := json.Marshal(shop)
//
//	//fmt.Println(shop)
//
//	// ===> Di bawah ini untuk Key <===//
//	autho := d.GenerateSignature("f4786d228ff53ae6090f3f735b32402a12d11988c8d8fdcc75aca0a5522ef750", shop, endpoint)
//
//	headers.Set("Authorization", autho)
//	headers.Set("Content-Type", "application/json")
//	//fmt.Println(headers)
//
//	if err != nil {
//		return result, errors.Wrap(err, "[DATA][GetDataShope]")
//	}
//
//	//fmt.Println(endpoint)
//	//fmt.Println(&resp)
//
//	result, err = d.client.PostJSON(ctx, endpoint, headers, json2, &resp)
//	//fmt.Printf("resultnyaaa %+v\n", resp)
//	if err != nil {
//		return result, errors.Wrap(err, "[DATA][GetStockGudang]")
//	}
//
//	return resp, nil
//}

func (d Data) GetDataShopeTransaction(ctx context.Context, shop shopeEntity.Shope) (shopeEntity.MutasiShopee, error) {
	var (
		resp shopeEntity.MutasiShopee
		//result   shopeEntity.TestShopee
		endpoint = d.shopeURL2
		err      error
		headers  = make(http.Header)
		//autho string
	)

	//json1 := shopeEntity.Shope{
	//	PartnerID:       844098, // ==> Di Sini untuk Partner ID
	//	ShopID:          175375997, // ==> Di Sini Untuk SHOP ID
	//	Timestamp:       time.Now().Unix(),
	//	PaginationOffset: 1,
	//	PaginationEntriesPerPage: 10,
	//	CreateTimeFrom:	1583924400,
	//	CreateTimeTo:	1583971200,
	//}
	json2, err := json.Marshal(shop)

	//fmt.Println(shop)

	// ===> Di bawah ini untuk Key <===//
	autho := d.GenerateSignature(shop.Api, shop, endpoint)
	//bd656ab344c897c9a37e7230a9835ab7b826e608ec84f9ea63009d40bde7f9f9
	//"partner_id":842234,
	//	"shopid":121791179,

	//2a2437155def6d2885e7ff7a14069d94d368199576bf4cedb5f42a0950ae7bb7
	//"partner_id":844464,
	//	"shopid":114789399,

	//3e38516c57b3e46f06f9b5072428ff839bb6b41c69c1bd0933406a00c71a879f
	//"partner_id":844783,
	//	"shopid":181521989,

	headers.Set("Authorization", autho)
	headers.Set("Content-Type", "application/json")
	//fmt.Println(headers)

	//if err != nil {
	//	return result, errors.Wrap(err, "[DATA][GetDataShope]")
	//}

	//fmt.Println(endpoint)
	//fmt.Println(&resp)

	_, err = d.client.PostJSON(ctx, endpoint, headers, json2, &resp)
	//fmt.Printf("resultnyaaa %+v\n", resp)
	if err != nil {
		return resp, errors.Wrap(err, "[DATA][GetStockGudang]")
	}

	return resp, nil
}

package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

type UrlHolder struct {
	urlList       map[string]string
	configuration Config
}

func (u *UrlHolder) handleGetRequest(ctx *fasthttp.RequestCtx) {
	fullurl, errstr := u.get(string(ctx.Path()))
	if nil == errstr {
		strLocation := []byte("Location")
		ctx.Response.Header.SetCanonical(strLocation, []byte(fullurl))
		ctx.Response.SetStatusCode(fasthttp.StatusMovedPermanently)
		return
	}
	fmt.Fprint(ctx, "Invalid Get Request")
	ctx.Response.SetStatusCode(400)
}

func (u *UrlHolder) handlePostRequest(ctx *fasthttp.RequestCtx) {

	if string(ctx.Path()) == "/Posturl" {
		res, er := u.store(string(ctx.Request.Body()))

		if nil != er {
			fmt.Fprintf(ctx, "Failed to generate short url"+res)
			return
		}

		ctx.Response.SetBody([]byte(res))
		ctx.Response.SetStatusCode(200)
		return

	} else {
		fmt.Fprint(ctx, "Invalid Post Request url")
		ctx.Response.SetStatusCode(400)
	}
}

func (u *UrlHolder) HandleFastHTTP(ctx *fasthttp.RequestCtx) {

	if ctx.IsGet() == true {
		u.handleGetRequest(ctx)

	} else if ctx.IsPost() == true {
		u.handlePostRequest(ctx)
	} else {
		fmt.Fprint(ctx, "Invalid type of Http Request")
		ctx.Response.SetStatusCode(404)
	}
}

func (objurl *UrlHolder) store(url string) (string, error) {

	baseurl := objurl.configuration.Server.Baseurl + ":" + strconv.FormatUint(uint64(objurl.configuration.Server.Port), 10)
	hexastring, err := randomHex(5)

	if nil != err {
		return "", errors.New("Failed to geenrate short url")
	}

	shorturl := baseurl + "/" + hexastring
	objurl.urlList["/"+hexastring] = url //TODO: use nosql instead of keeping data in memory map

	return shorturl, nil
}

func (obj *UrlHolder) get(hexa string) (string, error) {
	url, bFound := obj.urlList[hexa]

	if bFound == false {
		return "", errors.New("url not found")
	}

	return url, nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (u *UrlHolder) Initialize() error {

	er := u.configuration.Init()

	if er != nil {
		panic("Failed to initialize configuration.. Error: " + er.Error())
	}
	return nil
}

func (handler *UrlHolder) Start() {
	sport := handler.configuration.Server.Baseurl + ":" + strconv.FormatUint(uint64(handler.configuration.Server.Port), 10)
	fasthttp.ListenAndServe(sport, handler.HandleFastHTTP)
}

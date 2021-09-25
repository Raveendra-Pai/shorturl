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

	strPath := string(ctx.Path())

	Applog.Info("Received HTTP GET request for url: " + strPath)
	fullurl, errstr := u.get(strPath)

	if nil == errstr {
		strLocation := []byte("Location")
		ctx.Response.Header.SetCanonical(strLocation, []byte(fullurl))
		ctx.Response.SetStatusCode(fasthttp.StatusMovedPermanently)
		Applog.Info("Found short url" + strPath + "  mapped to " + fullurl)
		return
	}

	fmt.Fprint(ctx, "Invalid Get Request")
	Applog.Error("Invalid Get Request")
	ctx.Response.SetStatusCode(400)
}

func (u *UrlHolder) handlePostRequest(ctx *fasthttp.RequestCtx) {

	if string(ctx.Path()) == "/Posturl" {
		payload := string(ctx.Request.Body())
		Applog.Info("Received url shorten request for url: " + payload)
		res, er := u.store(payload)

		if nil != er {
			fmt.Fprintf(ctx, "Failed to generate short url"+res)
			Applog.Error("Failed to generate short url" + res)
			return
		}

		ctx.Response.SetBody([]byte(res))
		ctx.Response.SetStatusCode(200)
		return

	} else {
		fmt.Fprint(ctx, "Invalid Post Request url")
		ctx.Response.SetStatusCode(400)
		Applog.Info("Invalid Post Request url" + string(ctx.Path()))
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
		Applog.Error("Invalid type of Http Request")
	}
}

func (objurl *UrlHolder) store(url string) (string, error) {

	baseurl := objurl.configuration.Server.Baseurl + ":" + strconv.FormatUint(uint64(objurl.configuration.Server.Port), 10)
	hexastring, err := randomHex(5)

	if nil != err {
		Applog.Error("Failed to generate short url")
		return "", errors.New("Failed to geenrate short url")
	}

	shorturl := baseurl + "/" + hexastring
	objurl.urlList["/"+hexastring] = url //TODO: use nosql instead of keeping data in memory map
	Applog.Info("Stroing " + hexastring + " mapped to " + url)
	return shorturl, nil
}

func (obj *UrlHolder) get(hexa string) (string, error) {
	url, bFound := obj.urlList[hexa]

	if bFound == false {
		Applog.Error("url not found for hexa : " + hexa)
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

func (u *UrlHolder) Init() error {

	er := u.configuration.Init()

	if er != nil {
		panic("Failed to initialize configuration.. Error: " + er.Error())
	}

	if nil != Applog.Init(u.configuration.Server.Logfile) {
		panic("Failed to initialize log")
	}

	return nil
}

func (handler *UrlHolder) Start() {
	urltolisten := handler.configuration.Server.Baseurl + ":" + strconv.FormatUint(uint64(handler.configuration.Server.Port), 10)
	Applog.Info("starting http listener for url:" + urltolisten)
	fasthttp.ListenAndServe(urltolisten, handler.HandleFastHTTP)
}

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/adiatma85/golang-alter-url-shortener/storage"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Response Struct
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"shortUrl"`
}

// Handler Struct
type Handler struct {
	Schema  string
	Host    string
	Storage storage.Service
}

// Func to initiate handler
func New(schema, host string, storage storage.Service) *router.Router {
	router := router.New()

	handler := Handler{
		Schema:  schema,
		Host:    host,
		Storage: storage,
	}

	// Defining the API endpoints
	router.POST("/encode/", respondeHandler(handler.encode))
	router.GET("/{shortLink}", handler.redirect)
	router.GET("/{shortLink}/info", respondeHandler(handler.decode))

	return router
}

// Encode
func (h Handler) encode(ctx *fasthttp.RequestCtx) (interface{}, int, error) {
	var input struct {
		URL     string `json:"url"`
		Expires string `json:"expires"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body: %v", err)
	}

	uri, err := url.ParseRequestURI(input.URL)

	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("invalid url")
	}

	layoutISO := "2006-01-02 15:04:05"
	expires, err := time.Parse(layoutISO, input.Expires)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("invalid expiration date")
	}

	c, err := h.Storage.Save(uri.String(), expires)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("could not store in database: %v", err)
	}

	u := url.URL{
		Scheme: h.Schema,
		Host:   h.Host,
		Path:   c,
	}

	fmt.Printf("Generated link: %v \n", u.String())

	return u.String(), http.StatusCreated, nil
}

// Decode
func (h Handler) decode(ctx *fasthttp.RequestCtx) (interface{}, int, error) {
	code := ctx.UserValue("shortLink").(string)
	model, err := h.Storage.LoadInfo(code)
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("url not found")
	}

	return model, http.StatusOK, nil
}

// Redirect
func (h Handler) redirect(ctx *fasthttp.RequestCtx) {
	code := ctx.UserValue("shortLink").(string)
	uri, err := h.Storage.Load(code)

	if err != nil {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(http.StatusNotFound)
		return
	}
	ctx.Redirect(uri, http.StatusMovedPermanently)
}

// Helper function of responehandler
func respondeHandler(h func(ctx *fasthttp.RequestCtx) (interface{}, int, error)) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		data, status, err := h(ctx)
		if err != nil {
			data = err.Error()
		}

		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(status)

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(Response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

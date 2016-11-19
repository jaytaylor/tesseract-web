package interfaces

import (
	"fmt"
	"net/http"

	"github.com/gigawattio/web"
	"github.com/gigawattio/web/generics"
	"github.com/gigawattio/web/route"
	"github.com/jaytaylor/tesseract-web/pkg/tesseract"
	"github.com/nbio/hitch"
)

type WebService struct {
	*web.WebServer
}

func NewWebService(addr string) *WebService {
	service := &WebService{}
	options := web.WebServerOptions{
		Addr:    addr,
		Handler: service.activateRoutes().Handler(),
	}
	service.WebServer = web.NewWebServer(options)
	return service
}

func (service *WebService) activateRoutes() *hitch.Hitch {
	index := func(w http.ResponseWriter, req *http.Request) {
		web.RespondWithHtml(w, 200, `<html><head><title>hello world</title></head><body>hello world</body></html>`)
	}
	routes := []route.RouteMiddlewareBundle{
		route.RouteMiddlewareBundle{
			Middlewares: []func(http.Handler) http.Handler{service.LoggerMiddleware},
			RouteData: []route.RouteDatum{
				{"get", "/", index},
				{"post", "/v1/tesseract", service.tesseract},
				{"post", "/v1/tesseract/*url", service.tesseractUrl},
			},
		},
	}
	h := route.Activate(routes)
	return h
}

func (service *WebService) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("method=%s url=%s remoteAddr=%s referer=%s\n", req.Method, req.URL.String(), req.RemoteAddr, req.Referer())
		next.ServeHTTP(w, req)
	})
}

func (service *WebService) tesseract(w http.ResponseWriter, req *http.Request) {
	generics.GenericObjectEndpoint(w, req, func() (object interface{}, err error) {
		ocr, err := tesseract.FromReader(req.Body)
		if err != nil {
			return nil, err
		}
		return ocr, nil
	})
}

func (service *WebService) tesseractUrl(w http.ResponseWriter, req *http.Request) {
	url := hitch.Params(req).ByName("url")
	if url == "" {
		web.RespondWithJson(w, 400, web.Json{"error": "URL musn't be empty"})
		return
	}
	if url[0] == '/' {
		url = url[1:]
	}
	// response, err := http.Get(url)
	// if err != nil {
	// 	web.RespondWithJson(w, 500, web.Json{"error": err})
	// 	return
	// }
	// if response.StatusCode / 100 != 2 {
	// 	web.RespondWithJson(w, 500, web.Json{"error": fmt.Sprintf("received non-2xx response status-code=%v from url=%v", response.StatusCode, url)})
	// 	return
	// }
	generics.GenericObjectEndpoint(w, req, func() (object interface{}, err error) {
		ocr, err := tesseract.FromUrl(url)
		if err != nil {
			return nil, err
		}
		return ocr, nil
	})
}

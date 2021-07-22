package restcontroller

import (
	"github.com/emicklei/go-restful"
	"kubebuilderht/api/controllers"
	"net/http"
)

func HTService() (ws *restful.WebService) {
	ws = new(restful.WebService)

	ws.Path("/v1/ht")
	ws.Route(ws.GET("/").To(func(request *restful.Request, response *restful.Response) {
		CreateResp(response, http.StatusOK, "hello world!")
	}))
	ws.Route(ws.GET("/get").To(getHello))
	ws.Route(ws.GET("/list").To(listHT))

	return ws
}

func getHello(req *restful.Request, resp *restful.Response) {
	waitForHTReconsilerAvailable()
	controllers.HTReconsiler.GetHelloType("default", "hello")
	CreateResp(resp, http.StatusOK, "done")
}

func listHT(req *restful.Request, resp *restful.Response) {
	waitForHTReconsilerAvailable()
	controllers.HTReconsiler.ListHelloType()
	CreateResp(resp, http.StatusOK, "done")
}

func CreateResp(resp *restful.Response, statusCode int, msg string) {
	resp.WriteHeader(statusCode)
	resp.Write([]byte(msg))
}

func waitForHTReconsilerAvailable() {
	for {
		if controllers.HTReconsiler != nil {
			break
		}
	}
}

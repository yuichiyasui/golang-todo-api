// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// タスク一覧を取得する
	// (GET /tasks)
	ListTasks(c *gin.Context)
	// タスクを作成する
	// (POST /tasks)
	CreateTask(c *gin.Context)
	// タスク詳細を取得する
	// (GET /tasks/{taskId})
	GetTaskDetail(c *gin.Context, taskId string)
	// タスクを更新する
	// (PUT /tasks/{taskId})
	UpdateTask(c *gin.Context, taskId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ListTasks operation middleware
func (siw *ServerInterfaceWrapper) ListTasks(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListTasks(c)
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateTask(c)
}

// GetTaskDetail operation middleware
func (siw *ServerInterfaceWrapper) GetTaskDetail(c *gin.Context) {

	var err error

	// ------------- Path parameter "taskId" -------------
	var taskId string

	err = runtime.BindStyledParameter("simple", false, "taskId", c.Param("taskId"), &taskId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter taskId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTaskDetail(c, taskId)
}

// UpdateTask operation middleware
func (siw *ServerInterfaceWrapper) UpdateTask(c *gin.Context) {

	var err error

	// ------------- Path parameter "taskId" -------------
	var taskId string

	err = runtime.BindStyledParameter("simple", false, "taskId", c.Param("taskId"), &taskId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter taskId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateTask(c, taskId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/tasks", wrapper.ListTasks)
	router.POST(options.BaseURL+"/tasks", wrapper.CreateTask)
	router.GET(options.BaseURL+"/tasks/:taskId", wrapper.GetTaskDetail)
	router.PUT(options.BaseURL+"/tasks/:taskId", wrapper.UpdateTask)
}

type ListTasksRequestObject struct {
}

type ListTasksResponseObject interface {
	VisitListTasksResponse(w http.ResponseWriter) error
}

type ListTasks200JSONResponse []Task

func (response ListTasks200JSONResponse) VisitListTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListTasksdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response ListTasksdefaultJSONResponse) VisitListTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type CreateTaskRequestObject struct {
	Body *CreateTaskJSONRequestBody
}

type CreateTaskResponseObject interface {
	VisitCreateTaskResponse(w http.ResponseWriter) error
}

type CreateTask200JSONResponse struct {
	Id string `json:"id"`
}

func (response CreateTask200JSONResponse) VisitCreateTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateTaskdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response CreateTaskdefaultJSONResponse) VisitCreateTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetTaskDetailRequestObject struct {
	TaskId string `json:"taskId"`
}

type GetTaskDetailResponseObject interface {
	VisitGetTaskDetailResponse(w http.ResponseWriter) error
}

type GetTaskDetail200JSONResponse Task

func (response GetTaskDetail200JSONResponse) VisitGetTaskDetailResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTaskDetaildefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetTaskDetaildefaultJSONResponse) VisitGetTaskDetailResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type UpdateTaskRequestObject struct {
	TaskId string `json:"taskId"`
	Body   *UpdateTaskJSONRequestBody
}

type UpdateTaskResponseObject interface {
	VisitUpdateTaskResponse(w http.ResponseWriter) error
}

type UpdateTask200JSONResponse Task

func (response UpdateTask200JSONResponse) VisitUpdateTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateTaskdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response UpdateTaskdefaultJSONResponse) VisitUpdateTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// タスク一覧を取得する
	// (GET /tasks)
	ListTasks(ctx context.Context, request ListTasksRequestObject) (ListTasksResponseObject, error)
	// タスクを作成する
	// (POST /tasks)
	CreateTask(ctx context.Context, request CreateTaskRequestObject) (CreateTaskResponseObject, error)
	// タスク詳細を取得する
	// (GET /tasks/{taskId})
	GetTaskDetail(ctx context.Context, request GetTaskDetailRequestObject) (GetTaskDetailResponseObject, error)
	// タスクを更新する
	// (PUT /tasks/{taskId})
	UpdateTask(ctx context.Context, request UpdateTaskRequestObject) (UpdateTaskResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// ListTasks operation middleware
func (sh *strictHandler) ListTasks(ctx *gin.Context) {
	var request ListTasksRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListTasks(ctx, request.(ListTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(ListTasksResponseObject); ok {
		if err := validResponse.VisitListTasksResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateTask operation middleware
func (sh *strictHandler) CreateTask(ctx *gin.Context) {
	var request CreateTaskRequestObject

	var body CreateTaskJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateTask(ctx, request.(CreateTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreateTaskResponseObject); ok {
		if err := validResponse.VisitCreateTaskResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetTaskDetail operation middleware
func (sh *strictHandler) GetTaskDetail(ctx *gin.Context, taskId string) {
	var request GetTaskDetailRequestObject

	request.TaskId = taskId

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskDetail(ctx, request.(GetTaskDetailRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskDetail")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetTaskDetailResponseObject); ok {
		if err := validResponse.VisitGetTaskDetailResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateTask operation middleware
func (sh *strictHandler) UpdateTask(ctx *gin.Context, taskId string) {
	var request UpdateTaskRequestObject

	request.TaskId = taskId

	var body UpdateTaskJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateTask(ctx, request.(UpdateTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(UpdateTaskResponseObject); ok {
		if err := validResponse.VisitUpdateTaskResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xWwWsTTxT+V8r7/Y5LE/VS9latSMBDofUkOYzZ13Ta3ZlxZrYQyoLZXgRRS5HWoiB4",
	"UKuiBXtQUPvHTBvjfyEz06RpsklbczCnHXbfvPd9733vY9ehxhPBGTKtIFwHVVvGhLjjTSm5tAchuUCp",
	"KbrXNR6hfeqGQAhBaUlZHbIAElSK1Iu+ZQFIvJ9SiRGEd32G0/hq0Inn91awpm2uRaJWB0tHqGqSCk05",
	"K0RAo8LXShOdugT/S1yCEP4rnZIunTAu2ZILPjILQFMdX4AKjaATG5yB1606jN1CFxSyNLG5NI84BEDZ",
	"vOR1iUrZlJz1NqiLwpJlS9wB9EhhkUd8alaIqdn5CgSwhlK5PsGV6fJ02ZblAhkRFEK45l4FIIhedhhK",
	"mqhVd6qjtg/bdmKZVCII4TZVetFFWPpKcKb8RK6Wy14TTCNzF4kQMa25q6UV5QflW2xPVGNyoUm4GXjS",
	"RErS8JzPCABMfmjybyb/fPT1QfvNW5NvtQ+fmeYuuMglksb6UuBGYfLLUAjindnYMxvfnTZUmiRENgrB",
	"HT/dPv65Y5q7Jn9kVUPqyo3dsq1mAQiuClp/QyLR6FripYdKX+dR41LMRq7RkKaa5qf2+4+t508gGFyp",
	"7noMuXy8+XjwWt/u+ByD+5H5wDFUdpZuoSsM7nERkP5hH/142Xq4aZo7pvmqt1GVucnSnMm3OlCL1ZYF",
	"JytfWrePSpQN3f1b6FZ/DjWhsfMMSRLUKG2+oQKozDkrg9B5DATASGJ766tBb/e1TDHoaUn/pKpjquF8",
	"qxlhLe29L78O9ifUWrrgzrWWtGCwd0R06iz/ZqoT4mZj/SCM7YAjfhXGt8K/EX/rxUFre7/f5SZzBUy+",
	"1YE7xOnsPZRrHWWnMoYQYl4j8TJXOpwpz5TBqvHk3ggFffi9/dqX8b9YPeKHrJr9CQAA//8+epytSwsA",
	"AA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

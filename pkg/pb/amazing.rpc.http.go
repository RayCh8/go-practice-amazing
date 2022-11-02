// Code generated by protoc-gen-svc. DO NOT EDIT.
// source: pkg/pb/rpc.proto
package pb

import (
	"bytes"
	"io"
	"net/http"

	errCodes "github.com/AmazingTalker/at-error-code"
	"github.com/AmazingTalker/go-rpc-kit/contextkit"
	"github.com/AmazingTalker/go-rpc-kit/errorkit"
	"github.com/AmazingTalker/go-rpc-kit/jsonpbkit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CollectorType string

const (
	QueryCollector CollectorType = "query"
	ParamCollector CollectorType = "params"
)

type Collector struct {
	FromKey string
	ToKey   string
}

func RegisterGoAmazingHttpService(engine *gin.Engine, srv GoAmazingServer) {
	adapter := NewAmazingGinHttpAdapter(srv)
	EnrichGinRouter(engine, adapter)
}

type AmazingGinHttpAdapter struct {
	server GoAmazingServer
}

func NewAmazingGinHttpAdapter(srv GoAmazingServer) *AmazingGinHttpAdapter {
	return &AmazingGinHttpAdapter{
		server: srv,
	}
}

func corsMiddleware() gin.HandlerFunc {
	// https://github.com/gin-contrib/cors#default-allows-all-origins
	return cors.Default()
}

func EnrichGinRouter(e *gin.Engine, adapter *AmazingGinHttpAdapter) {
	e.Use(corsMiddleware())
	e.Use(logkit.Middleware())

	e.Handle(http.MethodGet, "/health", adapter.HealthHandler)

	e.Handle(http.MethodGet, "/config", adapter.ConfigHandler)

	e.Handle(http.MethodPost, "/api/record", adapter.CreateRecordHandler)

	e.Handle(http.MethodGet, "/api/records/:id", adapter.GetRecordHandler)

	e.Handle(http.MethodGet, "/api/records", adapter.ListRecordHandler)

}

func (a *AmazingGinHttpAdapter) HealthHandler(ctx *gin.Context) {

	req := &HealthReq{}

	err := jsonpbkit.Unmarshal(ctx.Request.Body, req)

	if err != nil && err != io.EOF {
		logkit.Errorf(ctx, "unmarshal body failed", logkit.Payload{"err": err})
		e := errorkit.NewFromError(errCodes.ErrUnmarshalBodyFailed, err, errorkit.WithHttpStatusCode(http.StatusBadRequest))
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx = logkit.EnrichRequestPayload(ctx, req)

	resp, err := a.server.Health(contextkit.ParseGinContext(ctx), req)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.Header("content-type", "application/json")

	if resp == nil {
		ctx.String(204, "")
		return
	}

	output, err := jsonpbkit.MarshalToString(resp)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.String(0, output)
}

func (a *AmazingGinHttpAdapter) ConfigHandler(ctx *gin.Context) {

	req := &ConfigReq{}

	err := jsonpbkit.Unmarshal(ctx.Request.Body, req)

	if err != nil && err != io.EOF {
		logkit.Errorf(ctx, "unmarshal body failed", logkit.Payload{"err": err})
		e := errorkit.NewFromError(errCodes.ErrUnmarshalBodyFailed, err, errorkit.WithHttpStatusCode(http.StatusBadRequest))
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx = logkit.EnrichRequestPayload(ctx, req)

	resp, err := a.server.Config(contextkit.ParseGinContext(ctx), req)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.Header("content-type", "application/json")

	if resp == nil {
		ctx.String(204, "")
		return
	}

	output, err := jsonpbkit.MarshalToString(resp)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.String(0, output)
}

func (a *AmazingGinHttpAdapter) CreateRecordHandler(ctx *gin.Context) {

	req := &CreateRecordReq{}

	err := jsonpbkit.Unmarshal(ctx.Request.Body, req)

	if err != nil && err != io.EOF {
		logkit.Errorf(ctx, "unmarshal body failed", logkit.Payload{"err": err})
		e := errorkit.NewFromError(errCodes.ErrUnmarshalBodyFailed, err, errorkit.WithHttpStatusCode(http.StatusBadRequest))
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx = logkit.EnrichRequestPayload(ctx, req)

	resp, err := a.server.CreateRecord(contextkit.ParseGinContext(ctx), req)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.Header("content-type", "application/json")

	if resp == nil {
		ctx.String(204, "")
		return
	}

	output, err := jsonpbkit.MarshalToString(resp.Record)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.String(201, output)
}

func (a *AmazingGinHttpAdapter) GetRecordHandler(ctx *gin.Context) {

	req := &GetRecordReq{}

	err := jsonpbkit.Unmarshal(ctx.Request.Body, req)

	if err != nil && err != io.EOF {
		logkit.Errorf(ctx, "unmarshal body failed", logkit.Payload{"err": err})
		e := errorkit.NewFromError(errCodes.ErrUnmarshalBodyFailed, err, errorkit.WithHttpStatusCode(http.StatusBadRequest))
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	v_ID := ctx.Param("id")
	req.ID = v_ID

	ctx = logkit.EnrichRequestPayload(ctx, req)

	resp, err := a.server.GetRecord(contextkit.ParseGinContext(ctx), req)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.Header("content-type", "application/json")

	if resp == nil {
		ctx.String(204, "")
		return
	}

	output, err := jsonpbkit.MarshalToString(resp.Record)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.String(200, output)
}

func (a *AmazingGinHttpAdapter) ListRecordHandler(ctx *gin.Context) {

	req := &ListRecordReq{}

	err := jsonpbkit.Unmarshal(ctx.Request.Body, req)

	if err != nil && err != io.EOF {
		logkit.Errorf(ctx, "unmarshal body failed", logkit.Payload{"err": err})
		e := errorkit.NewFromError(errCodes.ErrUnmarshalBodyFailed, err, errorkit.WithHttpStatusCode(http.StatusBadRequest))
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	v_PageSize, _ := ctx.GetQuery("size")
	req.PageSize = v_PageSize

	v_Page, _ := ctx.GetQuery("page")
	req.Page = v_Page

	ctx = logkit.EnrichRequestPayload(ctx, req)

	resp, err := a.server.ListRecord(contextkit.ParseGinContext(ctx), req)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.Header("content-type", "application/json")

	if resp == nil {
		ctx.String(204, "")
		return
	}

	buf := make([]bytes.Buffer, len(resp.Records))
	for i, m := range resp.Records {
		m := m
		var out bytes.Buffer
		if err := jsonpbkit.Marshal(&out, m); err != nil {
			logkit.Errorf(ctx, "marshal response failed", logkit.Payload{"err": err})
			e := errorkit.NewFromError(errCodes.ErrMarshalResponseFailed, err, errorkit.WithHttpStatusCode(http.StatusInternalServerError))
			ctx.JSON(e.HttpStatus(), e.GinHashMap())
			return
		}
		buf[i] = out
	}

	output, err := jsonpbkit.MarshalJsonBuffersToString(buf)

	if err != nil {
		e := errorkit.FormatError(err)
		ctx.JSON(e.HttpStatus(), e.GinHashMap())
		return
	}

	ctx.String(200, output)
}

package api

import (
	"encoding/json"
	"github.com/Smart-Machine/simplas-test-task/httpProxy/pkg/models/advertisement"
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/wrappers"
	"net/http"
	"regexp"
	"strconv"
)

func Create(ctx *gin.Context) {
	var ad advertisement.Advertisement

	err := ctx.ShouldBindJSON(&ad)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ServiceClient.Create(ctx, &proto.APICreateRequest{
		Id:         ad.ID,
		Categories: ad.Categories,
		Title:      ad.Title,
		Type:       ad.Type,
		Posted:     ad.Posted,
	})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get a response from the service"})
		return
	}

	regex, err := regexp.Compile("\\{(.*)")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	content := regex.FindString(res.Content)

	var body map[string]any
	err = json.Unmarshal([]byte(content), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status_code": res.StatusCode,
		"body":        body,
	})
}

func GetList(ctx *gin.Context) {
	title := ctx.Query("title")

	res, err := ServiceClient.GetList(ctx, &wrappers.StringValue{Value: title})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get a response from the service"})
		return
	}

	regex, err := regexp.Compile("\\{(.*)")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	content := regex.FindString(res.Content)

	var body map[string]any
	err = json.Unmarshal([]byte(content), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"body":        body,
	})
}

func GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := ServiceClient.GetOne(ctx, &wrappers.StringValue{Value: id})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get a response from the service"})
		return
	}

	regex, err := regexp.Compile("\\{(.*)")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	content := regex.FindString(res.Content)

	var body map[string]any
	err = json.Unmarshal([]byte(content), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"body":        body,
	})
}

func Update(ctx *gin.Context) {
	var ad advertisement.Advertisement

	err := ctx.ShouldBindJSON(&ad)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ad.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	res, err := ServiceClient.Update(ctx, &proto.APIUpdateRequest{
		Id: int64(id),
		Data: &proto.APICreateRequest{
			Id:         ad.ID,
			Categories: ad.Categories,
			Title:      ad.Title,
			Type:       ad.Type,
			Posted:     ad.Posted,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get a response from the service"})
		return
	}

	regex, err := regexp.Compile("\\{(.*)")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	content := regex.FindString(res.Content)

	var body map[string]any
	err = json.Unmarshal([]byte(content), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"body":        body,
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := ServiceClient.Delete(ctx, &wrappers.StringValue{Value: id})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get a response from the service"})
		return
	}

	regex, err := regexp.Compile("\\{(.*)")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	content := regex.FindString(res.Content)

	var body map[string]any
	err = json.Unmarshal([]byte(content), &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": res.StatusCode,
		"body":        body,
	})
}

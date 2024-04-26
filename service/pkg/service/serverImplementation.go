package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"strconv"
	"strings"
)

type ServiceServer struct {
	elasticClient *elasticsearch.Client
	proto.UnimplementedServiceServer
}

func (ss *ServiceServer) Create(_ context.Context, req *proto.APICreateRequest) (*proto.APIResponse, error) {
	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	reqStr := string(reqJson)

	res, err := ss.elasticClient.Index(
		"advertisement",
		strings.NewReader(reqStr),
		ss.elasticClient.Index.WithDocumentID(string(req.Id)),
		ss.elasticClient.Index.WithRefresh("true"),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &proto.APIResponse{
		StatusCode: int32(res.StatusCode),
		Content:    res.String(),
	}, nil
}

func (ss *ServiceServer) GetList(ctx context.Context, search *wrapperspb.StringValue) (*proto.APIResponse, error) {
	// TODO: Refactor query to multimatch for both `ro` and `ru` matching.
	query := `{"query": {"match": {"title.ro": "%s"} } }`

	req := esapi.SearchRequest{
		Index: []string{"advertisement"},
		Body:  strings.NewReader(fmt.Sprintf(query, search.GetValue())),
	}

	res, err := req.Do(ctx, ss.elasticClient)
	if err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
	defer res.Body.Close()

	return &proto.APIResponse{
		StatusCode: int32(res.StatusCode),
		Content:    res.String(),
	}, nil
}

func (ss *ServiceServer) GetOne(_ context.Context, id *wrapperspb.StringValue) (*proto.APIResponse, error) {
	res, err := ss.elasticClient.Get(
		"advertisement",
		id.GetValue(),
		ss.elasticClient.Get.WithRefresh(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &proto.APIResponse{
		StatusCode: int32(res.StatusCode),
		Content:    res.String(),
	}, nil
}

func (ss *ServiceServer) Update(_ context.Context, req *proto.APIUpdateRequest) (*proto.APIResponse, error) {
	reqJson, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	reqStr := fmt.Sprintf("{\"doc\": %s}", string(reqJson))

	res, err := ss.elasticClient.Update(
		"advertisement",
		strconv.FormatInt(req.Id, 10),
		strings.NewReader(reqStr),
		ss.elasticClient.Update.WithRefresh("true"),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &proto.APIResponse{
		StatusCode: int32(res.StatusCode),
		Content:    res.String(),
	}, nil
}

func (ss *ServiceServer) Delete(_ context.Context, id *wrapperspb.StringValue) (*proto.APIResponse, error) {
	res, err := ss.elasticClient.Delete(
		"advertisement",
		id.GetValue(),
		ss.elasticClient.Delete.WithRefresh("true"),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &proto.APIResponse{
		StatusCode: int32(res.StatusCode),
		Content:    res.String(),
	}, nil
}

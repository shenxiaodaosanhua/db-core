package core

import (
	"context"
	"db-core/helpers"
	"db-core/pbfiles"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DbService struct {
	pbfiles.UnimplementedDBServiceServer
}

func (s *DbService) Query(ctx context.Context, request *pbfiles.QueryRequest) (response *pbfiles.QueryResponse, err error) {
	api := SysConfig.FindAPI(request.Name)
	if api == nil {
		return nil, status.Error(codes.Unavailable, "error api name")
	}
	ret, err := api.Query(request.Params) // 返回值是一个map[string]interface{}
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}
	// 把map 转化为 StructList
	structList, err := helpers.MapListToStructList(ret)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pbfiles.QueryResponse{
		Message: "success",
		Result:  structList,
	}, nil
}

func (s *DbService) Exec(ctx context.Context, request *pbfiles.ExecRequest) (*pbfiles.ExecResponse, error) {
	api := SysConfig.FindAPI(request.Name)
	if api == nil {
		return nil, status.Error(codes.Unavailable, "error api name")
	}

	rows, selectKey, err := api.ExecBySql(request.Params)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}
	selectKeyValue, err := helpers.MapToStruct(selectKey)

	return &pbfiles.ExecResponse{
		Message:      "success",
		RowsAffected: rows,
		Select:       selectKeyValue,
	}, nil
}

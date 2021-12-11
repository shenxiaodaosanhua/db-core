package core

import (
	"context"
	"db-core/helpers"
	"db-core/pbfiles"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
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

func (s *DbService) Tx(server pbfiles.DBService_TxServer) error {
	tx := GormDB.Begin()
	for {
		txRequest, err := server.Recv()
		if err == io.EOF {
			tx.Commit()
			return nil
		}

		if err != nil {
			tx.Rollback()
			return err
		}
		api := SysConfig.FindAPI(txRequest.Name)
		if api == nil {
			tx.Rollback()
			return status.Error(codes.Unavailable, "api not found")
		}

		api.SetDB(tx)
		ret := make(map[string]interface{})
		if txRequest.Type == "query" {
			result, err := api.QueryBySql(txRequest.Params)
			if err != nil {
				tx.Rollback()
				return err
			}
			ret["query"] = result
		} else {
			rows, selectKey, err := api.ExecBySql(txRequest.Params)
			if err != nil {
				tx.Rollback()
				return err
			}

			ret["exec"] = []interface{}{
				rows,
				selectKey,
			}
		}
		m, _ := helpers.MapToStruct(ret)
		err = server.Send(&pbfiles.TxResponse{
			Message: "操作成功",
			Result:  m,
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
}

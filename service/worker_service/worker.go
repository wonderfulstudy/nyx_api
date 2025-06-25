package worker

import (
	"errors"
	"nyx_api/models"
	"nyx_api/pkg/log"
	"strconv"
)

func ListService(req WorkerListRequest) (WorkerListResponse, error) {
	var response WorkerListResponse
	page, err := strconv.Atoi(req.Page)
	if err != nil {
		return WorkerListResponse{}, errors.New("请求体中的分页参数page不是数值类型")
	}
	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		return WorkerListResponse{}, errors.New("请求体中的分页参数limit不是数值类型")
	}
	workers, err := models.ListWorkers(page, limit)
	if err != nil {
		return WorkerListResponse{}, errors.New("从数据库获取所有用户数据失败" + err.Error())
	}

	for _, worker := range workers {
		log.Log.Debug("worker", worker)
		response.Workers = append(response.Workers, WorkerInfoResponse{
			Id:            worker.Id,
			WorkerAddress: worker.WorkerAddress,
			OwnerAddress:  worker.OwnerAddress,
			Status:        worker.Status,
			Uptime:        worker.Uptime,
		})
	}
	response.Total = models.GetWorkerCount()

	return response, nil
}

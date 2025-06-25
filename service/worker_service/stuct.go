package worker

type WorkerListRequest struct {
	Page  string `form:"page" valid:"Required"`
	Limit string `form:"limit" valid:"Required"`
}

type WorkerListResponse struct {
	Total   int                  `json:"total"`
	Workers []WorkerInfoResponse `json:"workers"`
}

type WorkerInfoResponse struct {
	Id            int    `json:"id"`
	WorkerAddress string `json:"workerAddress"`
	OwnerAddress  string `json:"ownerAddress"`
	Status        string `json:"status"`
	Uptime        int    `json:"uptime"`
}

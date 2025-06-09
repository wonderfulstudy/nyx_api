package models

import (
	"nyx_api/pkg/setting"
	"sort"
	"time"
)

type Worker struct {
	Id            int       `gorm:"primary_key,int" json:"id"`
	WorkerAddress string    `json:"workerAddress"`
	OwnerAddress  string    `json:"ownerAddress"`
	Status        string    `json:"status"`
	Uptime        int       `json:"uptime"`
	ChainCreateAt time.Time `gorm:"timestamp" json:"chainCreateAt"`
	ChainUpdateAt time.Time `gorm:"timestamp" json:"chainUpdateAt"`
	CreatedAt     time.Time `gorm:"timestamp" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"timestamp" json:"updatedAt"`
}

// 获取所有矿机信息
func GetWorkers() (workers []Worker) {
	db.Find(&workers)
	return
}

// 根据矿机id获取矿机信息
func GetWorkerById(Id int) (worker Worker) {
	db.Where("id = ?", Id).First(&worker)
	return
}

// 根据矿机地址获取矿机信息
func GetWorkerByAddress(address string) (worker Worker) {
	db.Where("address = ?", address).First(&worker)
	return
}

func GetWorkerList(page, limit int) (workers []Worker) {
	db.Offset((page - 1) * setting.PageSize).Limit(limit).Find(&workers)
	return
}

func GetWorkerCount() (count int) {
	db.Model(&Worker{}).Count(&count)
	return
}

// desc 为 fasle 是进行正序排序 为true进行倒序排序
func SortWorkersByAddress(page, limit int, desc bool) (workers []Worker) {
	db.Offset((page - 1) * setting.PageSize).Limit(limit).Find(&workers)
	if desc {
		sort.SliceStable(workers, func(i, j int) bool {
			return workers[i].WorkerAddress > workers[j].WorkerAddress
		})
	} else {
		sort.SliceStable(workers, func(i, j int) bool {
			return workers[i].WorkerAddress < workers[j].WorkerAddress
		})
	}
	return workers
}

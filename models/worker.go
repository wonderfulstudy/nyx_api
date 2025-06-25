package models

import (
	"nyx_api/pkg/setting"
	"sort"
	"time"
)

type Worker struct {
	Id            int `gorm:"primary_key,int"`
	WorkerAddress string
	OwnerAddress  string
	Status        string
	Uptime        int
	OwnerUser     int       `gorm:"foreignkey:UserId;references:RESTRICT;onDelete:RESTRICT"`
	ChainCreateAt time.Time `gorm:"timestamp"`
	ChainUpdateAt time.Time `gorm:"timestamp"`
}

func GetWorkerCount() (count int) {
	db.Model(&Worker{}).Count(&count)
	return
}

func GetWorkerCountByOnweruser(userId int) (count int) {
	db.Model(&Worker{}).Where("owner_user = ?", userId).Count(&count)
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

func ListWorkers(page, limit int) ([]Worker, error) {
	var workers []Worker
	result := db.Model(&Worker{}).
		Offset((page - 1) * setting.PageSize).
		Limit(limit).
		Find(&workers)
	if result.Error != nil {
		return []Worker{}, result.Error
	}
	return workers, nil
}

package services

import (
	"SecKill_Product/repositories"
)

type ILogService interface {

	GetAllLogInfo()(map[int]map[string]string, error)

}

func NewlogService(repository repositories.ILogRepository)ILogService{
	return &logService{repository}
}

type logService struct {
	logRepository repositories.ILogRepository
}

func (t *logService) GetAllLogInfo()(map[int]map[string]string, error){
	return t.logRepository.SelectAllLog()
}
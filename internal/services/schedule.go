package services

import (
	"cinema-ticket-system/internal/models"
	"cinema-ticket-system/internal/repositories"
)

type ScheduleService struct {
	scheduleRepo *repositories.ScheduleRepository
}

func NewScheduleService(scheduleRepo *repositories.ScheduleRepository) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
	}
}

func (ss *ScheduleService) Create(schedule *models.Schedule) error {
	return ss.scheduleRepo.Create(schedule)
}

func (ss *ScheduleService) GetByID(id uint) (*models.Schedule, error) {
	return ss.scheduleRepo.GetByID(id)
}

func (ss *ScheduleService) Update(schedule *models.Schedule) error {
	return ss.scheduleRepo.Update(schedule)
}

func (ss *ScheduleService) Delete(id uint) error {
	return ss.scheduleRepo.Delete(id)
}

func (ss *ScheduleService) GetAll() ([]models.Schedule, error) {
	return ss.scheduleRepo.GetAll()
}
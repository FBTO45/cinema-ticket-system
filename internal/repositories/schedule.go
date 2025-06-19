package repositories

import (
	"cinema-ticket-system/internal/models"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (sr *ScheduleRepository) Create(schedule *models.Schedule) error {
	return sr.db.Create(schedule).Error
}

func (sr *ScheduleRepository) GetByID(id uint) (*models.Schedule, error) {
	var schedule models.Schedule
	if err := sr.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (sr *ScheduleRepository) Update(schedule *models.Schedule) error {
	return sr.db.Save(schedule).Error
}

func (sr *ScheduleRepository) Delete(id uint) error {
	return sr.db.Delete(&models.Schedule{}, id).Error
}

func (sr *ScheduleRepository) GetAll() ([]models.Schedule, error) {
	var schedules []models.Schedule
	if err := sr.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
package repository


import (
	"errors"

	"gorm.io/gorm"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/repository/gorm/models"
)


type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(task *domain.Task) error{
	taskModel := models.FromDomainTask(*task)

	if err:= r.db.Create(&taskModel).Error; err != nil {
		return err
	}


	task.ID = taskModel.ID

	return nil
}


func (r *taskRepository) GetByID(id uint)(*domain.Task,error){
	var taskModel models.Task

	if err:=r.db.First(&taskModel,id).Error; err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}

	domainTask := models.ToDomainTask(taskModel)

	return &domainTask,nil
}


func (r *taskRepository) Delete(id uint) error {
	var taskModel models.Task
	if err := r.db.First(&taskModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if err := r.db.Delete(&taskModel).Error; err != nil {
		return err
	}

	return nil
}


func (r *taskRepository) GetByUserId(userId uint) ([]domain.Task, error) {
	var taskModels []models.Task
	if err := r.db.Where("user_id = ?", userId).Find(&taskModels).Error; err != nil {
		return nil, err
	}

	var tasks []domain.Task
	for _, taskModel := range taskModels {
		tasks = append(tasks, models.ToDomainTask(taskModel))
	}

	return tasks, nil
}


func (r *taskRepository) Update(domainTask *domain.Task) error {
	taskModel := models.FromDomainTask(*domainTask)

	if err := r.db.Save(&taskModel).Error; err != nil {
		return err
	}

	return nil
}
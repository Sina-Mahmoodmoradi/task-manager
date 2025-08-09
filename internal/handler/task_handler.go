package handler

import (
	"net/http"
	"strconv"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"
	"github.com/gin-gonic/gin"
)


type TaskHandler struct{
	taskService domain.TaskService
}


func NewTaskHandler(taskService domain.TaskService) *TaskHandler{
	return &TaskHandler{
		taskService: taskService,
	}
}


type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type TaskResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type ListTasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

type TaskUpdateRequest struct {
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Done        *bool   `json:"done"`
}

func (h *TaskHandler)CreateTask(c *gin.Context){
	var req CreateTaskRequest
	
	if err:=c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid input"})
		return
	}

	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	task, err := h.taskService.CreateTask(userId.(uint),req.Title,req.Description)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
	})
}


func (h *TaskHandler)DeleteTask(c *gin.Context){
	idParam := c.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = h.taskService.DeleteTask(uint(idUint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (h *TaskHandler)GetTask(c *gin.Context){
	idParam := c.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.taskService.GetTaskByID(uint(idUint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
	})
}

func (h *TaskHandler)ListTasks(c *gin.Context){
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tasks, err := h.taskService.ListTasksByUser(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseTasks := make([]TaskResponse, 0, len(tasks))


	for _,task := range tasks{
		responseTasks = append(responseTasks, TaskResponse{
			ID: task.ID,
			Title: task.Title,
			Description: task.Description,
		})
	}

	c.JSON(http.StatusOK,ListTasksResponse{
		Tasks: responseTasks,
	})
}


func (h *TaskHandler)UpdateTask(c *gin.Context){
	idParam := c.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	} 

	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	var req TaskUpdateRequest
	if err:=c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid input"})
		return
	}

	updates := domain.TaskUpdate{
		Title:       req.Title,
		Description: req.Description,
		Done:       req.Done,
	}


	task, err := h.taskService.UpdateTask(uint(idUint64),userId.(uint),&updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
	})
}
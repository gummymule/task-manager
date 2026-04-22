package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gummymule/task-manager/internal/domain"
	"github.com/gummymule/task-manager/pkg/response"
)

type TaskHandler struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskHandler(taskUsecase domain.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase}
}

// GetAll godoc
// @Summary      Get all tasks
// @Description  Get all tasks for authenticated user with pagination
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Param        page   query    int  false  "Page number"
// @Param        limit  query    int  false  "Items per page"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Router       /tasks [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	userID := c.GetInt64("user_id")
	boardID, err := strconv.ParseInt(c.Param("board_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid board id")
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	tasks, err := h.taskUsecase.GetAll(userID, boardID, page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "success", tasks)
}

// GetByID godoc
// @Summary      Get task by ID
// @Description  Get a specific task by ID
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /tasks/{id} [get]
func (h *TaskHandler) GetByID(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.taskUsecase.GetByID(id, userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "success", task)
}

// Create godoc
// @Summary      Create new task
// @Description  Create a new task for authenticated user
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body domain.Task true "Task data"
// @Success      201  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /tasks [post]
func (h *TaskHandler) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")
	boardID, err := strconv.ParseInt(c.Param("board_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid board id")
		return
	}

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	task.UserID = userID
	task.BoardID = boardID
	result, err := h.taskUsecase.Create(&task)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "task created", result)
}

// Update godoc
// @Summary      Update task
// @Description  Update an existing task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path      int         true  "Task ID"
// @Param        request body      domain.Task true  "Task data"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /tasks/{id} [put]
func (h *TaskHandler) Update(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid task id")
		return
	}

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	task.ID = id
	task.UserID = userID
	result, err := h.taskUsecase.Update(&task)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "task updated", result)
}

// Delete godoc
// @Summary      Delete task
// @Description  Delete an existing task
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /tasks/{id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid task id")
		return
	}

	if err := h.taskUsecase.Delete(id, userID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "task deleted", nil)
}

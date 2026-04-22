package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gummymule/task-manager/internal/domain"
	"github.com/gummymule/task-manager/pkg/response"
)

type BoardHandler struct {
	boardUsecase domain.BoardUsecase
}

func NewBoardHandler(boardUsecase domain.BoardUsecase) *BoardHandler {
	return &BoardHandler{boardUsecase}
}

// GetAll godoc
// @Summary      Get all boards
// @Description  Get all boards for authenticated user
// @Tags         boards
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Router       /boards [get]
func (h *BoardHandler) GetAll(c *gin.Context) {
	userID := c.GetInt64("user_id")
	boards, err := h.boardUsecase.GetAll(userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "success", boards)
}

// GetByID godoc
// @Summary      Get board by ID
// @Description  Get a specific board by ID
// @Tags         boards
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Board ID"
// @Success      200  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /boards/{id} [get]
func (h *BoardHandler) GetByID(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, err := strconv.ParseInt(c.Param("board_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid board id")
		return
	}
	board, err := h.boardUsecase.GetByID(id, userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "success", board)
}

// Create godoc
// @Summary      Create new board
// @Description  Create a new board for authenticated user
// @Tags         boards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body domain.Board true "Board data"
// @Success      201  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /boards [post]
func (h *BoardHandler) Create(c *gin.Context) {
	userID := c.GetInt64("user_id")
	var board domain.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	board.UserID = userID
	result, err := h.boardUsecase.Create(&board)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "board created", result)
}

// Update godoc
// @Summary      Update board
// @Description  Update an existing board
// @Tags         boards
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path      int          true  "Board ID"
// @Param        request body      domain.Board true  "Board data"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /boards/{id} [put]
func (h *BoardHandler) Update(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, err := strconv.ParseInt(c.Param("board_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid board id")
		return
	}
	var board domain.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}
	board.ID = id
	board.UserID = userID
	result, err := h.boardUsecase.Update(&board)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "board updated", result)
}

// Delete godoc
// @Summary      Delete board
// @Description  Delete an existing board
// @Tags         boards
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Board ID"
// @Success      200  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /boards/{id} [delete]
func (h *BoardHandler) Delete(c *gin.Context) {
	userID := c.GetInt64("user_id")
	id, err := strconv.ParseInt(c.Param("board_id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid board id")
		return
	}
	if err := h.boardUsecase.Delete(id, userID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "board deleted", nil)
}
package handler

import (
	todo "github.com/TelitsynNikita"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllNewDeals(c *gin.Context) {
	deals, err := h.services.GetAllNew()
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, deals)
}

func (h *Handler) getOwnDeals(c *gin.Context) {
	var input todo.AllOwnDeal
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, ok := c.Get(userCtxId)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id not found")
		return
	}

	userRole, ok := c.Get(userCtxRole)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id not found")
		return
	}

	deals, err := h.services.GetAllOwnDeals(userId.(int), userRole.(string), input.Status)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, deals)
}

func (h *Handler) getOneDeal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	deal, err := h.services.GetOneDealById(id)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, deal)
}

func (h *Handler) createOneDeal(c *gin.Context) {
	id, ok := c.Get(userCtxId)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input []todo.Deal
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, value := range input {
		_, err := h.services.Create(id.(int), value)
		if err != nil {
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
	}
}

func (h *Handler) deleteDeal(c *gin.Context) {
	var input []todo.UpdateDealBookkeeperId
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, value := range input {
		err := h.services.Delete(value.DealId)
		if err != nil {
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Успешно удалено",
	})
}

func (h *Handler) updateDealStatus(c *gin.Context) {
	var input []todo.UpdateDealStatus
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, deal := range input {
		err := h.services.UpdateStatus(deal.Status, deal.Id)
		if err != nil {
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Успешно обновлено",
	})
}

func (h *Handler) updateBookkeeperId(c *gin.Context) {
	userId, ok := c.Get(userCtxId)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id not found")
		return
	}

	userRole, ok := c.Get(userCtxRole)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id not found")
		return
	}

	if userRole != "BOOKKEEPER" {
		newErrorMessage(c, http.StatusInternalServerError, "no access")
		return
	}

	var input []todo.UpdateDealBookkeeperId
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, value := range input {
		err := h.services.UpdateDealBookkeeperId(userId.(int), value.DealId)
		if err != nil {
			newErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Успешно обновлено",
	})
}

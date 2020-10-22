package controllers

import (
	"context"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/patiwatpanda/app/ent"
	"github.com/patiwatpanda/app/ent/pillorderitem"
	
)
 
//PillorderItemController defines the struct for the Pillorderitem controller
type PillorderItemController struct {
	client *ent.Client
	router gin.IRouter
}

type PillorderItem struct {
	 Pillorder int
	 PillorItemName string
}


// PillorderItemCreate handles POST requests for adding pillorderitem entities
// @Summary Create pillorderitem 
// @Description Create pillorderitem
// @ID create-pillorderitem
// @Accept   json
// @Produce  json
// @Param pillorderitem  body ent.PillorderItem true "PillorderItem entity"
// @Success 200 {object} ent.PillorderItem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pillorderItems [post]
func (ctl *PillorderItemController) PillorderItemCreate(c *gin.Context) {
	obj := ent.PillorderItem{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pillorderitem binding failed",
		})
		return
	}
	
	
	
	pi, err := ctl.client.PillorderItem.
		Create().
		SetPillorderItemName(obj.PillorderItemName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, pi)
 }
 


 // GetPillorderItem handles GET requests to retrieve a PillorderItem entity
// @Summary Get a PillorderItem entity by ID
// @Description get PillorderItem by ID
// @ID get-pillorderitem
// @Produce  json
// @Param id path int true "PillorderItem ID"
// @Success 200 {object} ent.PillorderItem
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pillorderitems/{id} [get]
func (ctl *PillorderItemController) GetPillorderItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pi, err := ctl.client.PillorderItem.
		Query().
		Where(pillorderitem.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pi)
}

// ListPillorderItem handles request to get a list of pillorderitem entities
// @Summary List pillorderitem entities
// @Description list pillorderitem entities
// @ID list-pillorderitem
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PillorderItem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pillorderitems [get]
func (ctl *PillorderItemController) ListPillorderItem(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	pillorderitems, err := ctl.client.PillorderItem.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pillorderitems)
}


// NewPillorderItemController creates and registers handles for the pillorderitem controller
func NewPillorderItemController(router gin.IRouter, client *ent.Client) *PillorderItemController {
	pic := &PillorderItemController{
		client: client,
		router: router,
	}
	pic.register()
	return pic
}
// InitPillorderItemController registers routes to the main engine
func (ctl *PillorderItemController) register() {
	pillorderitems := ctl.router.Group("/pillorderitems")

	pillorderitems.GET("", ctl.ListPillorderItem)
	 // CRUD
	 pillorderitems.POST("", ctl.PillorderItemCreate)
	 pillorderitems.GET(":id", ctl. GetPillorderItem)

}
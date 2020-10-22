package controllers

import (
	"context"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patiwatpanda/app/ent"
	"github.com/patiwatpanda/app/ent/pillorder"
	"github.com/patiwatpanda/app/ent/employee"
	"github.com/patiwatpanda/app/ent/patient"
	"github.com/patiwatpanda/app/ent/dentist"
	"github.com/patiwatpanda/app/ent/pillorderitem"

	
)
 
//PillorderController defines the struct for the Pillorder controller

type PillorderController struct {
	client *ent.Client
	router gin.IRouter
}
type Pillorder struct {
	Employee 		int
	Dentist   		int
	Patient 		int
	PillorderNameID  string
	PillorderItem   int
	PillorderDate   string
}
// PillorderCreate handles POST requests for adding pillorder entities
// @Summary Create pillorder 
// @Description Create pillorder 
// @ID create-pillorder 
// @Accept   json
// @Produce  json
// @Param pillorder  body ent.Pillorder true "Pillorder entity"
// @Success 200 {object} ent.Pillorder
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pillorders [post]
func (ctl *PillorderController) PillorderCreate(c *gin.Context) {
	obj := Pillorder{}
	 if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pillorder binding failed",
		})
		return
	 }
	e, err := ctl.client.Employee.
	Query().
	Where(employee.IDEQ(int(obj.Employee))). 
	Only(context.Background())

	if err != nil {
	c.JSON(400, gin.H{
		"error": "employee not found",
	})
	return
	}
	p, err := ctl.client.Patient.
	Query().
	Where(patient.IDEQ(int(obj.Patient))).
	Only(context.Background())

	if err != nil {
	c.JSON(400, gin.H{
	"error": "patient not found",
	})
	return
	}
	d, err := ctl.client.Dentist.
	Query().
	Where(dentist .IDEQ(int(obj.Dentist))).
	Only(context.Background())

	if err != nil {
	c.JSON(400, gin.H{
	"error": "dentist not found",
	})
	return
	}
	
	pi, err := ctl.client.PillorderItem.
	Query().
	Where(pillorderitem.IDEQ(int(obj.PillorderItem))). 
	Only(context.Background())

	if err != nil {
	c.JSON(400, gin.H{
		"error": "pillorderitem not found",
	})
	return
	}
	time, err := time.Parse(time.RFC3339, obj.PillorderDate)

	po, err := ctl.client.Pillorder.
		Create().
		SetPatient(p).
		SetEmployee(e).
		SetDentist(d).
		SetPillorderNameID(obj.PillorderNameID).
		SetPillorderitem(pi).
		SetPillorderDate(time).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}


	c.JSON(200, po)
		
}


 // GetPillorder handles GET requests to retrieve a Pillorder entity
// @Summary Get a Pillorder entity by ID
// @Description get pillorder by ID
// @ID get-pillorder
// @Produce  json
// @Param id path int true "Pillorder ID"
// @Success 200 {object} ent.Pillorder
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pillorders/{id} [get]
func (ctl *PillorderController) GetPillorder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	po, err := ctl.client.Pillorder.
		Query().
		Where(pillorder.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, po)
}

// ListPillorder handles request to get a list of pillorder entities
// @Summary List pillorder entities
// @Description list pillorder entities
// @ID list-pillorder
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Pillorder
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pillorders [get]
func (ctl *PillorderController) ListPillorder(c *gin.Context) {
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

	pillorders, err := ctl.client.Pillorder.
		Query().
		WithDentist().
		WithEmployee().
		WithPatient().
		WithPillorderitem().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pillorders)
}




// NewPillorderController creates and registers handles for the pillorder controller
func NewPillorderController(router gin.IRouter, client *ent.Client) *PillorderController {
	poc := &PillorderController{
		client: client,
		router: router,
	}
	poc.register()
	return poc
}
// InitPillorderController registers routes to the main engine

func (ctl *PillorderController) register() {
	pillorders := ctl.router.Group("/pillorders")

	pillorders.GET("", ctl.ListPillorder)
	 // CRUD
	pillorders.POST("", ctl.PillorderCreate)
	pillorders.GET(":id", ctl.GetPillorder)

}
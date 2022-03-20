//Package controller provides functions to translate incoming request
//into outcoming response
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wilorios/microservice-template-go2/internal/entity"
	"github.com/wilorios/microservice-template-go2/internal/service"
)

type iPersonController interface {
	Save(ctx *gin.Context) entity.Person
	FindAll() []entity.Person
}

type PersonControlle struct {
	personService service.PersonServ
}

//New function create a new controller
func New(ser service.PersonServ) *PersonControlle {
	return &PersonControlle{}
}

//FindAll function search all the entity person calling the service
func (c *PersonControlle) FindAll() []entity.Person {
	return c.personService.FindAll()
}

//Save function get the request and bind to entity and call the service
//to save the data
func (c *PersonControlle) Save(ctx *gin.Context) entity.Person {
	var e entity.Person
	ctx.BindJSON(&e)
	c.personService.Save(e)
	return e
}

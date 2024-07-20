package main

import (
	"log"
	"net/http"
	"reflect"
	"github.com/gin-gonic/gin"
)

func HandleRequest(c *gin.Context) {
	//1.get operation type
	operationType := c.GetHeader("operationType")

	//2.get facade
	f := FacadeRegister[operationType]
	fValue := reflect.ValueOf(f)

	//3.get param type
	paramType := fValue.Type().In(0)

	//4.instance param
	paramInstance := reflect.New(paramType).Interface()

	//5.bind request data
	c.ShouldBindJSON(paramInstance)

	//6.build reflect call param
	var args []reflect.Value
	args = append(args, reflect.ValueOf(reflect.ValueOf(paramInstance).Elem().Interface()))

	//7.print request start log
	log.Printf("start handle request.opretation type = %s, request = %+v", operationType, args[0])

	//8.invoke function
	results := fValue.Call(args)

	//9.write result
	c.JSON(http.StatusOK, results[0].Elem().Interface())

	//10.print request end log
	log.Printf("end handle request.opretation type = %s, result = %+v", operationType, results[0])

}

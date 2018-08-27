package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// curl http://127.0.0.1:3000/test1/_a1?a2=rr
	// {"a1":"_a1","a2":"rr","a3":"123"}
	router.GET(`/test1/:a1`, func(c *gin.Context) {
		a1 := c.Param(`a1`)
		a2 := c.Query(`a2`)
		a3 := c.DefaultQuery(`a3`, `123`)
		c.JSON(200, gin.H{
			`a1`: a1,
			`a2`: a2,
			`a3`: a3,
		})
	})

	// curl http://127.0.0.1:3000/test2 -X POST -i -H "Content-Type:application/json" -d '{"boolean" : false, "foo" : "bar"}'
	// {"input":{"boolean":false,"foo":"bar"}}
	router.POST(`/test2`, func(c *gin.Context) {
		inputJSON := map[string]interface{}{}
		c.BindJSON(&inputJSON)
		c.JSON(200, gin.H{
			`input`: inputJSON,
		})
	})

	router.Run(`:3000`)
}

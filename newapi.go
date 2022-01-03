package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    // handler is block of code that is executed for the specific endpoint
	//.POST  is the route,
	
	//  create a new git router
	 router := gin.Default()
     
	 // "/" is the endpoint 
	 
     
	 router.POST("/", Posthomepage)
     
     
     //run the server on port 3000
	_  = router.Run(":3000")
}



func Posthomepage(c *gin.Context) {
	fmt.Println("Posthomepage")
	c.JSON(200, gin.H{
		"message": "Post home page",
	    

	})
}
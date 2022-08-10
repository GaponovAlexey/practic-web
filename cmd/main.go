package main

import "log"

func main() {
	router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")

}



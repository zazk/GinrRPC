package main

import (
	"fmt"
	"grpc_tutorial/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a := getParam(ctx, "a")
		b := getParam(ctx, "b")
		req := &proto.Request{A: a, B: b}

		if response, err := client.Add(ctx, req); err == nil {
			success(ctx, response)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a := getParam(ctx, "a")
		b := getParam(ctx, "b")
		req := &proto.Request{A: a, B: b}

		if response, err := client.Multiply(ctx, req); err == nil {
			success(ctx, response)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}

func getParam(ctx *gin.Context, param string) int64 {
	fmt.Println("getParameters...")
	a, err := strconv.ParseUint(ctx.Param(param), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return 0
	}
	return int64(a)
}

func success(ctx *gin.Context, response *proto.Response) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(response.Result),
	})
}

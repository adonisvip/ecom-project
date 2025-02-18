package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"ecom-gateway/grpc"
	pbService "ecom-gateway/proto/service"
)

// GetProductHandler lấy thông tin sản phẩm
func GetProductHandler(c *gin.Context) {
	productID := c.Param("id")

	res, err := grpc.ServiceClient.GetProduct(context.Background(), &pbService.ProductRequest{ProductId: productID})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}

package router

import (
	db "go-api-crud/db"
	"net/http"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
    r.GET("/product", getProducts)
    r.GET("/product/:id", getProduct)
    r.POST("/product", postProduct)
    r.PUT("/product/:id", putProduct)
    r.DELETE("/product/:id", deleteProduct)
    return r
}

func postProduct(ctx *gin.Context) {
    var product db.Product
    err := ctx.Bind(&product)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    res, err := db.CreateProduct(&product)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{
        "product": res,
    })
}

func getProducts(ctx *gin.Context) {
    res, err := db.GetProducts()
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "products": res,
    })
 }
  
 func getProduct(ctx *gin.Context) {
    id := ctx.Param("id")
    res, err := db.GetProduct(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "product": res,
    })
 }

 func putProduct(ctx *gin.Context) {
    var updateProduct db.Product
    err := ctx.Bind(&updateProduct)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    id := ctx.Param("id")
    dbProduct, err := db.GetProduct(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    dbProduct.ProductName = updateProduct.ProductName
    dbProduct.Desc = updateProduct.Desc
  
    res, err := db.UpdateProduct(dbProduct)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "task": res,
    })
 }

 func deleteProduct(ctx *gin.Context) {
    id := ctx.Param("id")
    err := db.DeleteProduct(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "task deleted successfully",
    })
 }
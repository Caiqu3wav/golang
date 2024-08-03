package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupProductRoutes(r *gin.Engine, client *prisma.Client) {
	r.GET("/products", func(c *gin.Context) {
		products, err := client.Product.FindMany().Exec()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.POST("/products", func(c *gin.Context) {
        var product prisma.ProductCreateInput
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        newProduct, err := client.Product.Create(product).Exec()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, newProduct)
    })

    r.GET("/products/:id", func(c *gin.Context) {
        id := c.Param("id")
        product, err := client.Product.FindOne(prisma.ProductID.Equals(id)).Exec()
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "Produto não encontrado"})
            return
        }
        c.JSON(http.StatusOK, product)
    })

    r.PUT("/products/:id", func(c *gin.Context) {
        id := c.Param("id")
        var product prisma.ProductUpdateInput
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        updatedProduct, err := client.Product.UpdateOne(prisma.ProductID.Equals(id)).Data(product).Exec()
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "Produto não encontrado"})
            return
        }
        c.JSON(http.StatusOK, updatedProduct)
    })

    r.DELETE("/products/:id", func(c *gin.Context) {
        id := c.Param("id")
        err := client.Product.DeleteOne(prisma.ProductID.Equals(id)).Exec()
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "Produto não encontrado"})
            return
        }
        c.JSON(http.StatusNoContent, nil)
    })
}
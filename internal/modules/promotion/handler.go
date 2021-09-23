package promotion

import (
	"github.com/gin-gonic/gin"
)

type PromotionHandler struct{
	repository *PromotionRepository
}
func NewPromotionHandler(repo *PromotionRepository) *PromotionHandler{
	return &PromotionHandler{repository: repo}
}
func(handler *PromotionHandler) ListPromotions(c *gin.Context){
	data, err := handler.repository.List()
	if err != nil{
		c.JSON(500,gin.H{
			"error": err,
		})
	}
	c.JSON(200,&data)
}
func(handler *PromotionHandler) CreatePromotion(c *gin.Context){
	var promotion Promotion
	err := c.ShouldBindJSON(&promotion)
	if err != nil{
		c.JSON(500,gin.H{
			"error": "Incorrect Format",
		})
		return
	}
	if promotion.Description == ""{
		c.JSON(500,gin.H{
			"error": "Empty Description",
		})
		return
	}
	if len(promotion.Description) > 120{
		c.JSON(500,gin.H{
			"error": "Description max 80 characters",
		})
		return
	}
	if promotion.Percentage < 0.1{
		c.JSON(500,gin.H{
			"error": "Null Percentaje",
		})
		return
	}
	if promotion.Percentage > 70{
		c.JSON(500,gin.H{
			"error": "Percentaje can't above to 70%",
		})
		return
	}
	if promotion.StartDate == ""{
		c.JSON(500,gin.H{
			"error": "Empty StartDate",
		})
		return
	}
	if promotion.FinishDate == ""{
		c.JSON(500,gin.H{
			"error": "Empty FinishDate",
		})
		return
	}
	existDates, err := handler.repository.GetPromotionForDate(promotion.StartDate, promotion.FinishDate)
	if err != nil{
		c.JSON(500,gin.H{
		   "error": err,
	   })
	   return
   	}
	if len(*existDates) > 0{
		c.JSON(200,gin.H{
			"error": "Exist Date",
		})
		return
	} 
	data, err := handler.repository.Create(promotion)
	if err != nil{
	 	c.JSON(500,gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200,&data)
}
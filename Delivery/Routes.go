package Delivery

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("karyawan", GetKaryawan)
		grp1.POST("karyawan", CreateKaryawan)
		grp1.GET("karyawan/:id", GetKaryawanByID)
		grp1.PUT("karyawan/:id", UpdateKaryawan)
		grp1.DELETE("karyawan/:id", DeleteKaryawan)
	}
	return r
}


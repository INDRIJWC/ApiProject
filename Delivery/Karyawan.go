package Delivery

import (
	"ApiProject/Entity"
	"ApiProject/Usescase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetKaryawan(c *gin.Context) {
	var user []Entity.Karyawan
	err:=Usescase.GetUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateKaryawan(c *gin.Context) {
	var karyawan Entity.Karyawan
	c.BindJSON(&karyawan)
	err := Usescase.CreateUser(&karyawan)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, karyawan)
	}
}

func GetKaryawanByID(c *gin.Context)  {
	id := c.Params.ByName("id")
	var karyawan Entity.Karyawan
	err := Usescase.GetKaryawanByID(&karyawan, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, karyawan)
	}
}

func  UpdateKaryawan(c *gin.Context)  {
	var user Entity.Karyawan
	id := c.Params.ByName("id")
	err := Usescase.GetKaryawanByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Usescase.UpdateNasabah(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteKaryawan(c *gin.Context)  {
	var user Entity.Karyawan
	id := c.Params.ByName("id")
	err := Usescase.DeleteKaryawan(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

}






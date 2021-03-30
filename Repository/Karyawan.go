package Repository

import (
	"ApiProject/Config"
	"ApiProject/Entity"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllNasabahs(karyawan *[]Entity.Karyawan) (err error) {
	if err = Config.DB.Find(&karyawan).Error; err != nil {
		return err
	}
	return nil
}

func CreateNasabah(karyawan *Entity.Karyawan) (err error) {
	if err = Config.DB.Create(&karyawan).Error; err != nil {
		return err
	}
	return nil
}

func GetNasabahByID(karyawan *Entity.Karyawan, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(karyawan).Error; err != nil {
		return err
	}
	return nil
}

func UpdateNasabah(karyawan *Entity.Karyawan, id string) (err error) {
	fmt.Println(karyawan)
	Config.DB.Save(karyawan)
	return nil
}


func DeleteNasabah(karyawan *Entity.Karyawan, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(karyawan)
	return nil
}

package Usescase

import (
	"ApiProject/Entity"
	"ApiProject/Repository"
)

func GetUsers(karyawan *[]Entity.Karyawan) error {
	return Repository.GetAllNasabahs(karyawan)
}

func CreateUser(karyawan *Entity.Karyawan)error {
	return Repository.CreateNasabah(karyawan)
}

func GetKaryawanByID(karyawan *Entity.Karyawan, id string)error {
	return Repository.GetNasabahByID(karyawan, id)
}

func UpdateNasabah(karyawan *Entity.Karyawan, id string)error {
	return Repository.UpdateNasabah(karyawan, id)
}

func DeleteKaryawan(karyawan *Entity.Karyawan, id string)error {
	return Repository.DeleteNasabah(karyawan, id)
}

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Usuario struct {
	ID     uint `gorm:"primaryKey"`
	Nombre string
	Edad   int
}

func main() {
	dsn := "usuario:contraseña@tcp(127.0.0.1:3306)/basedatos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}

	db.AutoMigrate(&Usuario{})

	// Crear un nuevo registro
	nuevoUsuario := Usuario{Nombre: "Juan", Edad: 30}
	db.Create(&nuevoUsuario)
	fmt.Println("Nuevo usuario creado:", nuevoUsuario)

	// Leer registros
	var usuarios []Usuario
	db.Find(&usuarios)
	fmt.Println("Usuarios:", usuarios)

	// Actualizar un registro
	db.Model(&nuevoUsuario).Update("Edad", 31)
	fmt.Println("Usuario actualizado:", nuevoUsuario)

	// Eliminar un registro
	db.Delete(&nuevoUsuario)
	fmt.Println("Usuario eliminado")
}

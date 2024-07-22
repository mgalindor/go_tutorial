### 14. Trabajar con Bases de Datos

Trabajar con bases de datos es una parte fundamental del desarrollo de aplicaciones modernas. En Go, se pueden utilizar varias bibliotecas y paquetes para interactuar con bases de datos SQL y NoSQL. En este tema, cubriremos cómo trabajar con bases de datos SQL utilizando la biblioteca `database/sql` estándar y el ORM GORM para simplificar las operaciones con bases de datos.

## Contenido

1. Conexión a una Base de Datos SQL
2. Consultas Básicas con `database/sql`
3. Uso del ORM GORM
4. Operaciones CRUD con GORM

### Conexión a una Base de Datos SQL

Para conectarse a una base de datos SQL en Go, se utiliza el paquete `database/sql` junto con un driver específico de la base de datos (por ejemplo, `github.com/go-sql-driver/mysql` para MySQL).

#### Instalación del Driver de MySQL

```sh
go get -u github.com/go-sql-driver/mysql
```

#### Ejemplo de Conexión a una Base de Datos MySQL

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    dsn := "usuario:contraseña@tcp(127.0.0.1:3306)/basedatos"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error al abrir la conexión:", err)
        return
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println("Error al conectar a la base de datos:", err)
        return
    }

    fmt.Println("Conexión exitosa a la base de datos")
}
```

**Explicación:**
- `sql.Open("mysql", dsn)`: Abre una conexión a la base de datos utilizando el Data Source Name (DSN) especificado.
- `db.Ping()`: Verifica la conexión a la base de datos.

### Consultas Básicas con `database/sql`

Una vez establecida la conexión, se pueden realizar consultas SQL utilizando los métodos `Query` y `Exec` de `*sql.DB`.

#### Ejemplo de Consultas Básicas

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    dsn := "usuario:contraseña@tcp(127.0.0.1:3306)/basedatos"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error al abrir la conexión:", err)
        return
    }
    defer db.Close()

    // Insertar un nuevo registro
    insertQuery := "INSERT INTO usuarios(nombre, edad) VALUES (?, ?)"
    result, err := db.Exec(insertQuery, "Juan", 30)
    if err != nil {
        fmt.Println("Error al insertar el registro:", err)
        return
    }
    id, _ := result.LastInsertId()
    fmt.Println("Nuevo registro insertado con ID:", id)

    // Seleccionar registros
    selectQuery := "SELECT id, nombre, edad FROM usuarios"
    rows, err := db.Query(selectQuery)
    if err != nil {
        fmt.Println("Error al seleccionar los registros:", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var nombre string
        var edad int
        err := rows.Scan(&id, &nombre, &edad)
        if err != nil {
            fmt.Println("Error al escanear el registro:", err)
            return
        }
        fmt.Printf("ID: %d, Nombre: %s, Edad: %d\n", id, nombre, edad)
    }
}
```

**Explicación:**
- `db.Exec(insertQuery, "Juan", 30)`: Ejecuta una consulta de inserción.
- `db.Query(selectQuery)`: Ejecuta una consulta de selección.
- `rows.Scan(&id, &nombre, &edad)`: Escanea los resultados de la consulta en las variables especificadas.

### Uso del ORM GORM

GORM es un ORM (Object-Relational Mapping) para Go que simplifica las operaciones con bases de datos al mapear estructuras Go a tablas de bases de datos.

#### Instalación de GORM

```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

#### Ejemplo de Conexión y Migración con GORM

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Usuario struct {
    ID     uint   `gorm:"primaryKey"`
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

    // Migrar el esquema
    db.AutoMigrate(&Usuario{})
    fmt.Println("Conexión exitosa y migración completada")
}
```

**Explicación:**
- `gorm.Open(mysql.Open(dsn), &gorm.Config{})`: Abre una conexión a la base de datos utilizando GORM.
- `db.AutoMigrate(&Usuario{})`: Realiza la migración del esquema de la base de datos para la estructura `Usuario`.

### Operaciones CRUD con GORM

GORM facilita las operaciones CRUD (Crear, Leer, Actualizar, Eliminar) con métodos sencillos y expresivos.

#### Ejemplo de Operaciones CRUD

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Usuario struct {
    ID     uint   `gorm:"primaryKey"`
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
```

**Explicación:**
- `db.Create(&nuevoUsuario)`: Crea un nuevo registro en la base de datos.
- `db.Find(&usuarios)`: Lee todos los registros de la tabla `usuarios`.
- `db.Model(&nuevoUsuario).Update("Edad", 31)`: Actualiza el campo `Edad` del registro `nuevoUsuario`.
- `db.Delete(&nuevoUsuario)`: Elimina el registro `nuevoUsuario` de la base de datos.

### Resumen

En este tema, hemos cubierto cómo trabajar con bases de datos en Go utilizando tanto la biblioteca estándar `database/sql` como el ORM GORM. Hemos visto cómo conectar a una base de datos, realizar consultas básicas, y llevar a cabo operaciones CRUD utilizando GORM. Estas habilidades son esenciales para desarrollar aplicaciones que interactúan con bases de datos de manera eficiente y estructurada.
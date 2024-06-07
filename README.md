# go-back-project

#### [Memoria: ](memoria/memoria.md)
#### [Pequeña referencia: *Work in progress*](HEEDME.md)
#### [Frontend: *Not fully functional*](https://github.com/miguel010939/react-front)

# Instalar Go

La instalación del SDK de Go se puede hacer desde su [web oficial](https://go.dev/dl/). El proceso para Windows es particularmente
sencillo, consiste sólo en ejecutar el instalador.


# Base de Datos

Como base de datos utilizo una imagen de postgres desplegada en un contenedor Docker.

Para configurar una equivalente, se descarga del registro público una imagen de `postgres:latest` y se ejecuta el contenedor.
La configuración de este debe ser consistente con la *string* de conexión *hardcodeada* en `config.go` [<=](config/config.go). 
Así, mi usuario de postgres es `postgres`, mi contraseña es `password` y conecto en el puerto predeterminado de postgres `5432`.

# Ejecutar el Servidor

El punto de entrada es la función `main()` en `main.go` del paquete `main`. Para correr el servidor, ejecutar en una terminal en 
el directorio raíz del repo 

```bash
go run main.go
```

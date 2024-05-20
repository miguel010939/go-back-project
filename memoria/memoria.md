
## Indice

1. Descripción del proyecto y ámbito de implantación
2. Temporalización del proyecto y fases de desarrollo
3. Recursos de hardware y software
4. Arquitectura de software y sistemas
5. Descripción de datos

![](imgs/go_mascot.png)

## Descripción del proyecto y ámbito de implantación
El proyecto consiste en un API REST en Go para una web de subastas online. Los usuarios pueden registrarse, subastar sus productos y comprar los de otros.

La idea es poder simular una subasta en tiempo real: el servidor notifica de forma concurrente a todos los usuarios pujando por un producto dado a través de un SSE (Server Sent Event)

La base de datos PostgreSQL se aloja en un container Docker. 

## Temporalización del proyecto y fases de desarrollo



## Recursos de hardware y software

### Requisitos de Hardware
Ni idea.

Los máximos de los mínimos requisitos para el software requerido.

### Requisitos de Software

+ Docker
+ PostgreSQL (en un container de Docker)
+ Go SDK

## Arquitectura de software y sistemas

Diagrama Entidad-Relación
```mermaid
---
title: Diagrama E-R Subastas
---
erDiagram
    USUARIO }o--o{ PRODUCTO : puja
    USUARIO ||--|| SESION : inicia
    USUARIO }o--o{ PRODUCTO : favorito
    USUARIO }o--o{ USUARIO : sigue

```
Diagrama de la Estructura de Base de Datos
```mermaid
---
title: DB
---
classDiagram
    Session ..> User
    Follower ..> User
    Product ..> User
    Bid ..> User
    Favorite ..> User
    Bid ..> Product
    Favorite ..> Product
    class User{
        INT id (PK)
        VARCHAR(50) username
        VARCHAR(256) hashedpswd
        VARCHAR(100) email
    }
    class Session{
        INT id (PK)
        VARCHAR(32) token
        INT user (FK)
        TIMESTAMP time
    }
    class Follower{
        INT id (PK)
        INT usera (FK)
        INT userb (FK)
    }
    class Product{
        INT id (PK)
        VARCHAR(50) name
        VARCHAR(300) description
        VARCHAR(150) imageurl
        INT user (FK)
    }
    class Bid{
        INT id (PK)
        INT user (FK)
        INT product (FK)
        DECIMAL(8,2) ammount 
        TIMESTAMP time
    }
    class Favorite{
        INT id (PK)
        INT user (FK)
        INT product (FK)
        
    }
```


### Estructura del Proyecto
El proyecto está estructurado en 4 capas: el *router*, los *handlers*, los repositorios y los modelos.

| Router                           | Handler              | Repository                                | Model                                 |
|----------------------------------|----------------------|-------------------------------------------|---------------------------------------|
| Asocia las URLs a sus *handlers* | Gestiona la petición | Conecta directamente con la Base de Datos | Structs & Métodos asociados al objeto |

#### Router
El *router* se encarga de asociar a cada URL su *handler* correspondiente

#### Handlers
Los *handlers* procesan la petición HTTP recuperando de la *request* los parámetros (*query* & *path*) que sean necesarios.
Los *handlers* llaman a los repositorios y presentan los datos obtenidos por ellos o devuelven los errores oportunos, según sea el caso.

#### Repositories
Los repositorios son los que conectan directamente con la base de datos. Devuelven los datos de las consultas, ejecutan actualizaciones y eliminaciones.

#### Models
Representan en el programa objetos asociados a las tablas de base de datos.


```javascript
// TODO Diagrama & Explicación de las tablas de la base de datos
```
![Diagrama de las tablas de base de datos](imgs/placeholder1.jpg)

```javascript
// TODO Detallar endpoints del API REST
```



## Descripción de datos



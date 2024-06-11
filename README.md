# Backend Golang Hexagonal

### Tecnologias Empleadas

- Go (Lenguaje de programación)
- Gin Framework (Framework para Go)
- PostgreSQL (Base de datos)

La estructura basica de los datos proporcionados para realizar es la siguiente:
Un array que dentro contiene multiples objetos y a su ves estos tienen las siguientes propiedades.
- que tienen un "id" del tipo numero
- un "average_price" igualmente del tipo numero
- tambien un "name"
- y un "brand_name"
![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/68f02b44-8b9c-4279-b674-1690ce42b60c)

### Endpoints de la Aplicación
Aqui se encuentran todos los endpoints que deben existir como se puede ver en la imagen son los siguientes, GET */brands*, GET */brands/:id/models*, POST */brands*, POST */brands/:id/models*, PUT */models/:id* y GET */models*.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/7e1ff56d-b341-4873-9989-30fc47e5710d)

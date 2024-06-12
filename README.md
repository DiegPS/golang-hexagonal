# Backend Golang

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

### explicacion muy por encima de la app

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/4651a825-9053-4b9c-9f41-2aebbe1ac423)


### Endpoints de la Aplicación
Aqui se encuentran todos los endpoints que deben existir como se puede ver en la imagen son los siguientes, GET */brands*, GET */brands/:id/models*, POST */brands*, POST */brands/:id/models*, PUT */models/:id* y GET */models*.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/7e1ff56d-b341-4873-9989-30fc47e5710d)


### Recuerda crear un archivo `.env`

con la siguiente estructura para ejecutar el proyecto:

```
PORT=8080
DATABASE_URL=postgresql://postgres:1234@localhost:5432/car_database?sslmode=disable
RUN_MIGRATIONS=true
```

En ese ejemplo, el usuario se llama `postgres` la contraseña es `1234` y el nombre de la bases de datos es `car_database` pero logicamente puedes colocar los nombres que tu desees, tambien tenemos correr migraciones por defecto en true, pero una vez las hayas creado te recomiendo que lo pases a false.

### Estructura de la aplicación

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/7f6f64d0-8ef9-4555-b353-b6dbc32a7419)

#### Desplegado en usando Serverless technologies

Zeabur configurado para que despliegue todo en funciones empleando el `zbpack` que es una configuracion que permite realizar esto y evitar la necesidad de tener una maquina corriendo `24/7`, y `Neon` para desplegar PostgreSQL, de forma `serverless` sin la necesidad de pagar tanto en el caso de que no la estes usando.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/46ecec85-1e96-4d06-944f-72528268f1e6)






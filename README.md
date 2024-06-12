# Backend Golang

¿Quieres probar la aplicación? Puedes hacerlo en el siguiente enlace: https://golang-hex.zeabur.app/

### Tecnologias Empleadas

- Go (Lenguaje de programación)
- Gin Framework (Framework para Go)
- PostgreSQL (Base de datos)

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/447d451f-8830-4d2b-9b75-01496c6683d4)


Se emplea un archivo llamado `models.json` en el cual se encuentran los siguientes campos que ves en la imagen, `id`, `name`, `average_price` y `brand_name` que son los campos que se emplean en la aplicacion.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/68f02b44-8b9c-4279-b674-1690ce42b60c)

### Explicacion por encima de la arquitectura de la aplicación(no es hexagonal)

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/4651a825-9053-4b9c-9f41-2aebbe1ac423)

### Endpoints de la Aplicación
Aqui se encuentran todos los endpoints que deben existir como se puede ver en la imagen son los siguientes, GET */brands*, GET */brands/:id/models*, POST */brands*, POST */brands/:id/models*, PUT */models/:id* y GET */models*.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/7e1ff56d-b341-4873-9989-30fc47e5710d)


### Estructura de la aplicación

Hablando un poco mas de la estructura de la aplicacion, se puede identificar que se tiene un archivo `main.go` que es el archivo principal de la aplicacion, en el cual se cargan las variables de entorno en el caso de que este en modo desarrollo o produccion, se conecta a la base de datos y se ejecutan las migraciones en el caso de que se haya configurado para que se ejecuten, tambien se ejecuta el servidor en el puerto que se haya configurado.

posteriormente, tenemos la carpeta `routes` en la cual se encuentran los diferentes endpoints que se han creado, en este caso tenemos los endpoints de `brands` y `models` que son los que se han creado en la aplicacion.

luego la carpeta `controllers` que se encargan de gestionar las peticiones que se realizan a los diferentes endpoints, para luego ir al archivo `services` que se encargan de gestionar la logica de negocio de la aplicacion, y luego al archivo `repositories` que se encargan de gestionar las peticiones a la base de datos.

¿Pero por que tenemos una carpeta llamada `errors`, otra llamada `models` y hasta una llamada `database`? `errors` nos permite tener los diferentes errores que se pueden presentar en la aplicacion centralizados en un solo lugar, `models` nos permite tener los modelos que se emplean en la aplicacion y que pertenecen al dominio de la aplicacion, y `database` nos permite tener la configuracion de la misma y la conexion en un solo lugar sin contaminar el resto de la aplicacion.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/7f6f64d0-8ef9-4555-b353-b6dbc32a7419)

#### Desplegado usando Serverless technologies

Zeabur configurado para que despliegue todo en funciones empleando el `zbpack` que es una configuracion que permite realizar esto y evitar la necesidad de tener una maquina corriendo `24/7`, y `Neon` para desplegar PostgreSQL, de forma `serverless` es decir que solo se paga por lo que uses y no por tener una maquina corriendo todo el tiempo al igual que en el caso de `Zeabur`.

![image](https://github.com/DiegPS/golang-hexagonal/assets/88301232/46ecec85-1e96-4d06-944f-72528268f1e6)

¿Pero por que serverless? es una forma de desplegar aplicaciones de forma mas eficiente y economica, ya que solo pagas por lo que usas, aunque puede tambien tener sus inconvenientes como todo, ademas de que es mas facil de escalar y de mantener, ya que no tienes que preocuparte por la infraestructura, ya que el proveedor de servicios se encarga de eso.

## Ejecutar esto en local

Para ejecutar este proyecto en local, necesitas tener instalado `Go` y `PostgreSQL` en tu maquina.

tambien puedes hacer uso de `Docker` para ejecutar la base de datos en un contenedor, si no deseas instalar PostgreSQL en tu maquina y por lo tanto tambien de `Docker Compose` para ejecutar la base de datos y la aplicacion en contendores sin instalar ninguna de estas.

### Configuración de la base de datos

Primero, necesitas crear una base de datos en PostgreSQL. Puedes hacerlo ejecutando el siguiente comando en tu terminal:

```bash
CREATE DATABASE car_database;
```

recuerda que debes crear las tablas y las relaciones necesarias para que la aplicacion funcione correctamente.

para crear la tabla de `brands` puedes ejecutar el siguiente comando:

```bash
CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE
);
```

Esto crear una tabla con un campo `id` que es un numero autoincrementable y un campo `name` que es un string unico.

para crear la tabla de `models` puedes ejecutar el siguiente comando:

```bash
CREATE TABLE models (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    average_price INTEGER,
    brand_id INTEGER REFERENCES brands(id)
);
```

Esto creara una tabla con un campo `id` que es un numero autoincrementable, un campo `name` que es un string, un campo `average_price` que es un numero entero y un campo `brand_id` que es un numero entero que hace referencia a la tabla `brands`.

por lo tanto la relacion entre las tablas es de uno a muchos, es decir una marca puede tener muchos modelos, pero un modelo solo puede tener una marca.

Esto es un configuración basica, por lo que puede que si envias numeros negativos o strings en los campos de `average_price` o `brand_id` la aplicacion falle, pero es normal son errores que sabemos que pueden suceder, pero no estamos comprobando en este caso.

### configuracion de las variables de entorno

Necesitas crear un archivo `.env` en la raiz del proyecto, con las siguientes variables de entorno:

```
PORT=8080
DATABASE_URL=postgresql://postgres:1234@localhost:5432/car_database?sslmode=disable
RUN_MIGRATIONS=true
PROD=false
```

En ese ejemplo, el usuario se llama `postgres` la contraseña es `1234` el puerto es `5432` y el nombre de la bases de datos es `car_database` pero logicamente puedes colocar los nombres que tu desees, tambien tenemos correr migraciones por defecto en true, pero una vez las hayas creado te recomiendo que lo pases a false.

tambien si te fijas en la variable `PROD` esta en false, esto es para que la aplicacion se ejecute en modo desarrollo, si deseas que se ejecute en modo produccion cambia el valor a true.

### Ejecutar el proyecto

Para ejecutar el proyecto, necesitas ejecutar el siguiente comando en tu terminal:

```bash
go run .
```

Este comando deberia ser suficiente para que la aplicacion se ejecute en tu maquina local, en el puerto `8080` si no has cambiado la configuracion.

es decir puedes ir a tu navegador, y colocar la siguiente url `http://localhost:8080/brands` y deberias ver un json con todas las marcas que has creado en tu base de datos.

si deseas comprobar que la aplicacion funciona correctamente puedes hacer uso de la herramienta `Postman` o `Insomnia` para realizar las peticiones a los diferentes endpoints.

### Testing en la aplicacion

para el testing usamos una variable de entorno `.env.test` en donde tenemos un `DATABASE_URL` diferente al de desarrollo, ya que no queremos que los test afecten a la base de datos de desarrollo, por lo que creamos una base de datos de test.

ejemplo de `.env.test`:

```bash
PORT=8080
DATABASE_URL=postgresql://postgres:1234@localhost:5432/car_database_test?sslmode=disable
RUN_MIGRATIONS=true
```

Para ejecutar los test de la aplicacion, necesitas ejecutar el siguiente comando en tu terminal:

```bash
go test ./...
```

Este comando deberia ser suficiente para que los test de la aplicacion se ejecuten correctamente, y puedas ver si la aplicacion funciona correctamente.

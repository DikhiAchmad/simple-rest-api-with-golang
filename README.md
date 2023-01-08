<h1 align="center" style="font-size: 30px;margin:0;"><b>Simple REST API CRUD Golang</b></h1>
<p align="center" style="font-size: 15px;margin:20px 0;">at this time I want to do research on golang. when learning about frameworks, what was in my mind was always learning simple crud xixi. the package that I use is gin and gorm. the explanation is below here:</p>

### What is Gin?

```
    Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API, but with performance up to 40 times faster than Martini. If you need smashing performance, get yourself some Gin.
```

source: https://gin-gonic.com/docs/

### What is Gorm?

```
   GORM is an ORM (Object Relational Mapping) library for Golang. Let's start writing the CRUD API now. Create a new GO project and install related dependencies. We install GORM, Gin and Mysql by running the following command, in cmd or terminal that is already running in our project directory/folder.
```

source: https://gorm.io/

## How to use ?

The API you create should be able to store users via route :

- Method : POST

- URL : /user

- Body Request:

```
{
    "Name": string,
    "Email": string,
    "Password": string
}
```

Response to be returned :

- Status Code : 200
- Response Body :

```json
{
  "data": {
    "id": 1,
    "email": "dani@dani.com",
    "name": "dani",
    "password": "$2a$14$Y1/lKHMzLqUSSS.dGR3NnO6XZVPjzUbLfTUARV9WVkannBEcKZpxi",
    "created_at": "2023-01-08T18:40:06.5953625+07:00",
    "updated_at": "2023-01-08T18:40:06.5953625+07:00"
  }
}
```

## API can display all Users

The API you create should be able to display all Users stored via route :

- Method : GET

- URL : /user

## API can display one User

The API you create should be able to display Users stored via route :

- Method : GET

- URL : /user/:id

## API can change User data

The API you create must be able to change user data based on id via route:

- Method : PATCH

- URL : /user/:id

- Body Request:

```
{
    "Name": string,
    "Email": string,
    "Password": string
}
```

## API can delete User data

The API that you create must be able to change user data based on id via route:

- Method : DELETE

- URL : /user/:id

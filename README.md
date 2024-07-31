# Gin GORM REST APIs

This repository contains a RESTful API built using the Gin framework and GORM for database operations.

## Features

- CRUD operations
- Database interaction with GORM
- RESTful API structure

## Requirements

- Go 1.15 or higher
- A running PostgreSQL database

## Getting Started

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/muthukumar89uk/gin-RESTAPI-postgres-gorm.git
   ```
Click here to directly [download it](https://github.com/muthukumar89uk/gin-RESTAPI-postgres-gorm/zipball/master).

## Install dependencies:

          go mod tidy

## Run the Application
  1. Run the Server
   
       ```
          go run .
       ```   
  2. The server will start on `http://localhost:8080`.

## API Endpoints
 
- `POST /user`      - Create a new user
- `GET /users`      - Retrieve all user
- `GET /user/:id`   - Retrieve an user by ID
- `PUT /user/:id`   - Update an existing user
- `DELETE /user/:id`- Delete an user

## Refer
  - [Gin Web Framework](https://github.com/gin-gonic/gin) 
  - [GORM](https://gorm.io/)
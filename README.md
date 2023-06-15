# Monolith Application - README

## Project Overview

This project is a monolithic application written in Go that provides a secure way to handle user authentication using JWT (JSON Web Tokens). It includes functionality for user registration, login, and token refresh, along with secure routes that require authentication to access. The application uses the Gin web framework and Gorm for ORM. For storage, it utilizes a PostgreSQL database.

This project also includes a python script as an API client that communicates with the application, performing functions like user registration, login, token refreshing, and retrieving user information.

## Project Structure

The application codebase is divided into several packages that handle different functionalities of the application:

- `cmd`: Contains the main application execution setup, including database connection setup and server routing setup.
- `database`: Manages the connection to the database, including handling retries in case of a failure.
- `handlers`: Contains the HTTP handlers that deal with user registration, login, and token refresh, and it also exposes user information for authenticated routes.
- `middlewares`: Houses the middleware for authorization of secure routes and logging of HTTP requests.
- `models`: Defines the data models used in the application, including the User, File, and Refresh Token models.
- `security`: Provides functions to create and validate JWT tokens.
- `utils`: Contains utility functions like loading environment variables.

## Setup

Ensure you have a Postgres database set up and ready to connect to. For this, you can use the provided Docker compose file. It sets up two services, a Postgres database, and an Adminer service for database management.

## Environment Variables

The application requires the following environment variables:

- `DB_HOST`: The hostname of your PostgreSQL server.
- `DB_PORT`: The port your PostgreSQL server is running on.
- `POSTGRES_USER`: The PostgreSQL user.
- `POSTGRES_DB`: The PostgreSQL database name.
- `POSTGRES_PASSWORD`: The PostgreSQL password.
- `SECRET`: The secret key used for signing JWT tokens.
- `ENV`: (Optional) Can be used to determine the running environment of the application.

These should be placed in a `.env` file at the root of your project. Note that this file is included in the `.gitignore` to prevent accidental commit of sensitive information.

## Running the application

To run the application, navigate to the root directory and execute the following command:

```sh
docker compose up -d;
go run main.go;
```

This starts the application, sets up the required routes, and begins listening for incoming requests on port 8080.

## Usage

The application provides several endpoints for interaction:

- `/auth/register`: Register a new user (POST request expecting 'email' and 'password' in the request body).
- `/auth/login`: Login as a user and get access and refresh tokens (POST request expecting 'email' and 'password' in the request body).
- `/auth/refresh`: Refresh the access token using the refresh token (POST request expecting the refresh token in the cookies).
- `/api/me`: Get the authenticated user's information (GET request expecting the 'Authorization' header with the access token).

In the Python API client script, methods are provided for each of these endpoints. You can use it to interact with the application.

## Logging

By default, the application logs all incoming requests to a file named `gin.log`. This logging includes basic request information and any errors that occurred while processing the request.

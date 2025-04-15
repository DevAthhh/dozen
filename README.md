# DoZen

# About
This is simple REST API for database interactions. In this project, I used the jwt token authorization system, and on the database side, I used GORM postgres.

# What did I use?
1. `gorm.io/gorm` - for database
2. `golang-jwt/jwt/v5` - for authorization
3. `labstack/echo` - for routing
4. `viper` - for config
5. `godotenv` - for load .env file
6. `zap` - for logging

# Endpoint's
## POST `api/v1/u/login` 
This route represents the entry point to the applications. It accepts email and password input. And returns the user's token and ID

## POST `api/v1/u/register`
The same route is a registration point. It accepts three input values: username, password, and email.

## GET `api/v1/u/:id`
This endpoint returns information about the user, his tasks, his groups etc.

## POST `api/v1/t`
This endpoint allows you to create a task by accepting the input of the ID of the group for which the task will be created.

## PUT `api/v1/t`
This route allows you to change the status of the task (If status = done, then the task is marked as completed)

## DELETE `api/v1/t`
And the same route deletes the task.

## POST `api/v1/g`
This endpoint allows to you create a group.

#!/bin/bash

###############
####
# This example structure shows a typical 3-layered Clean Architecture design in Go, 
# where each layer is in its own subdirectory. The cmd directory contains the entry 
# point of the application, internal directory contains the internal components of the application, 
# and pkg directory contains the reusable library components.

# Overall, the goal of a Clean Architecture is to promote
# maintainability, flexibility, and testability by creating clear separation of concerns 
# between the different layers of the application.

#cmd/
mkdir -p ./cmd/server/

# The entry point of the application that initializes and starts the server.
touch ./cmd/server/main.go
############
#internal/
##api
mkdir -p ./internal/api/handlers
##HTTP handler functions for handling user-related requests.
touch ./internal/api/handlers/user.go

#Defines the routes and their corresponding handlers.
touch ./internal/api/routes.go
##biz
##
mkdir -p ./internal/biz/user/

# Defines the user service interface and its implementation.
touch ./internal/biz/user/service.go
# Defines the user model object.
touch ./internal/biz/user/model.go

##data
mkdir -p ./internal/data/repository
# Defines the user repository interface and its implementation.
touch ./internal/data/repository/user.go

# Defines the database connection object and helper functions for executing SQL queries.
touch ./internal/data/repository/db.go
mkdir -p ./internal/data/database
# Defines the MySQL database connector object.
touch ./internal/data/database/mysql.go
###########
#pkg/
mkdir -p ./pkg/server
# Defines the HTTP server object.
touch pkg/server/server.go
mkdir -p ./pkg/config
# Defines the configuration object for the application.
touch ./pkg/config/config.go




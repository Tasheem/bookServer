# Library Management System - Book Service
This service is a part of the Library Management System application.  The purpose of this service is to offer CRUD functionality for the books associated with the book store. This service only accepts requests from the Node.js authServer application.

## Tools Used For This Project
---
* Golang
* SQL
* JSON

## Project Structure
---
* A standard layered structure.
  * The handler functions are in the main file.
  * The handler functions interact with the service layer and the service layer interacts with the data access layer.
# go-music
This repository contains the source code of GoMusic web application which is the example in the book "[Hands-On Full Stack Development with Go(Mina Andrawos)](https://www.packtpub.com/web-development/hands-full-stack-development-go)".

## Content List
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)

## Technology Stack
- **Server**
   - Frontend: React
   - Backend: Gin (Go web framework)
- **Database**
   - MySQL

## Getting Started
### MySQL (Database)
- Install MySQL
- Create [database and tables](sql/create_schema.sql)
- Insert [data](sql/insert_data.sql) into database

### Gin (Backend RESTful API Server)
- Install Golang
- Install Gin
- Start Gin Server

### React (Frontend Server)
- Install [Node.js](https://nodejs.org/en/)
- Install all dependencies.
  ```bash
  cd go-music
  npm install
  ```
- Start react server.
  ```bash
  npm start
  ```

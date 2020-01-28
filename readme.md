# go-music
This repository contains the source code of GoMusic web application which is the example in the book "[Hands-On Full Stack Development with Go(Mina Andrawos)](https://www.packtpub.com/web-development/hands-full-stack-development-go)".

## Content List
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
- [Design](#design)
- [Configuration](#configuration)
- [Differences to The Original Source Code](#differences-to-the-original-source-code)
- [Study Nofe of React](doc/react.md)

## Technology Stack
- **Server**
   - Frontend: React
   - Backend: Gin (Go web framework)
- **Database**
   - MySQL
- **Web Design**
   - Bootstrap
- **Online Payment**
   - [Stripe](https://stripe.com/)

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
- Start react server (development mode).
  ```bash
  npm start
  ```
  (For production, use command: `npm run build`)

## Design
Please see [here](doc/design.md)

## Configuration
Please see [here](doc/configuration.md)

## Differences to The Original Source Code
- Flatten the source code by chapter into one application.
- Make sure the code is runnable (Fix some issues in the original source code).
- Better documentation.
- Add comments to make the code easy to read.


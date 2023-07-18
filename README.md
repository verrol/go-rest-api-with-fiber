# Go Fiber Web Framework Demo

A Simple application to explore some of the features of using Fiber to build micro-services. This app is inspired by and mostly based off of https://blog.logrocket.com/express-style-api-go-fiber/. NOTE: Since this article is for Go Fiber v1, a few things won't work as per the article.

The goal is to be able to demonstrates the following:
* Handling routes that are fixed, parameterized, or optional
* Server static files (images, .js, .css, etc)
* Using middlewares for logging, auth, etc
* Calling out to other services like a DB and keyvalue store
* Handling DTOs

## Approach

Evolve the directory structure as needed. Commit small changes that demonstrate a feature or a goal.

## Getting Started

This project uses 'go-task' to make it easier to rerun the app during development. It also have tasks defined to build Docker image and start dependent services using Docker Compose.

### Setup and Installation

#### Required

1. Go toolchain: https://go.dev/dl/
2. Code Editor or IDE:
  * Codium: https://github.com/VSCodium/vscodium/releases
  * VSCode: https://code.visualstudio.com/download
  * Jetbrain GoLand: https://www.jetbrains.com/go/download/
2. Docker: https://docker.com
3. Go-Task: https://taskfile.dev/installation

#### Optional

1. Git: https://https.gitscm.com
2. DB managment tool
  * DBeaver: https://dbeaver.io

### Developing

In the project directory, run the following to watch for changes and rerun the task.

`task -w dev`

### Running

Use the task 'up' to both build the Docker container and then start the application with all depending services.

`task up`

## Monitoring

This app uses Fiber's Monitor Middle. Access monitor graphs at http://<your-host/ip>:<port>/metrics.
# Go Chat App

This is a simple chat application that allows multiple users to chat together. It's built using:

- [GO](https://go.dev/) for backend
- [Vue3](https://vuejs.org/) for frontend
- [Postgresql](https://www.postgresql.org/) as a database

## Installation

This project was set up to be previewed using docker. To run the project:

1. Install [Docker](https://docs.docker.com/engine/install/). Make sure it's working.
2. While in the root of the project. Run `docker compose up`
3. Navigate to app at <http://localhost:8080>

## Ideas to improve the app with

- Hide BE and FE behind an nginx reverse proxy
- Add HTTPS support
- Adjust docker compose to support development mode
- Install nodejs version dynamically using version in `.nvmrc`
- Add unit testing in BE using the `testing` package
- Add unit testing in FE using `Vitest`
- Add e2e testing using `Cypress`
- Add Github Action to run tests
- Separate ChatView logic into multiple hooks with single responsibilities
- Separate FE to BE communication into separate data drivers

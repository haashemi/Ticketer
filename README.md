# Ticketer

A dead simple ticket service web app. Made in about a month from scratch as a university project.

![Home Page](./_screenshots/1-home.png)

> More screenshots are [here](./_screenshots)

## Tech Stack

Ticketer is built on top of these languages, frameworks, and tools:

- Backend:
  - Go
  - iris web-framework
- Frontend:
  - Typescript
  - Sveltekit
- Database:
  - PostgreSQL

## Getting started

To get started deploying Ticketer for whatever reason. You can follow the steps below.

### Requirements:

- [Golang](https://go.dev/)
- [Postgres](https://www.postgresql.org/)
- Node.JS + PNPM

### Step 0: Setup your environment

After installing the requirements, you have to setup your environment for both the backend and frontend.

1. Create a new database in Postgres.

2. Rename `config.example.yaml` to `config.yaml` and update everything to match with your requirements. (Especially the `jwt_key` and `database`)

3. Rename `.env.example` from the `frontend/` folder to `.env.local` and update `API_HOST` with `api_addr` you've set from step 0.2.

### Step 1: Start the backend

```bash
# Build the backend.
go build .
# Run the server.
./Ticketer
# or .\Ticketer.exe if you're on Windows.
```

### Step 2: Add some mocked data.

Ticketer doesn't come with an admin panel or anything like that; it just reads the movies from the database. so you have to add them yourself.

With that said, I already added a few random movies in `sql/mock/movies.sql`. execute it in your database.

You can also see the movie assets used for this sample data in the `static/movie` folder.

### Step 3: Start the frontend:

```bash
# Install the dependencies
pnpm install
# Build the frontend
pnpm build
# Preview
pnpm preview --port 5173
```

Keep in mind that the backend is hardcoded to use port 5173 for proxying the frontend, so if you're not using it behind a reverse proxy, be sure to use this port for the frontend.

### Step 4: You're done!

Now open up the backend host (`api_addr` from step `0.2`) in your browser and explore this project!

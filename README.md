# Poker Helper/Evaluator

This project is a poker helper/evaluator that evaluates poker hands and provides winning combinations based on the user's cards and the community cards.

## Tech Stack

### Backend
- Go (Golang)
- PostgreSQL
- Docker
- Gofiber (Web Framework)
- pgx (PostgreSQL driver)
- Redis (for caching)

### Frontend
- React
- TypeScript
- CSS

## Getting Started with Docker Compose

The easiest way to run the entire project is using Docker Compose. This method sets up the backend, PostgreSQL database, and Redis cache in separate containers.

### Prerequisites

- Docker
- Docker Compose

### Running the Project

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/poker-evaluator.git
   cd poker-evaluator
   ```

2. Start the project using Docker Compose:
   ```
   docker-compose up --build
   ```

   This command will build the backend image, start all the services defined in the `docker-compose.yml` file, and initialize the database with the required schema.

3. The backend API will be available at `http://localhost:8080`.

4. To stop the project, use:
   ```
   docker-compose down
   ```

## Manual Setup (Without Docker)

If you prefer to run the components separately or without Docker, follow these instructions:

### Prerequisites

- Go 1.22.5 or later
- PostgreSQL
- Redis
- Node.js and npm

### Backend Setup

1. Navigate to the backend directory:
   ```
   cd backend
   ```

2. Create a `.env` file with the following content (adjust as needed):
   ```
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=poker_evaluator
   DB_PORT=5432
   REDIS_HOST=localhost
   REDIS_PORT=6379
   PORT=8080
   ```

3. Install backend dependencies:
   ```
   go mod download
   ```

### Database Setup

1. Install PostgreSQL on your local machine if you haven't already.

2. Create a new database and user:
   ```
   createdb poker_evaluator
   createuser -s postgres
   psql
   ALTER USER postgres WITH PASSWORD 'pokerpassword';
   GRANT ALL PRIVILEGES ON DATABASE poker_evaluator TO postgres;
   ```

3. Initialize the database schema:
   ```
   psql -U postgres -d poker_evaluator -f backend/init.sql
   ```

### Running the Backend

1. Start the Redis server.

2. Run the backend application:
   ```
   cd backend
   go run main.go
   ```

The server will start on `http://localhost:8080`.

### Frontend Setup and Running (if applicable)

1. Navigate to the frontend directory:
   ```
   cd frontend
   ```

2. Install frontend dependencies:
   ```
   npm install
   ```

3. Start the frontend development server:
   ```
   npm start
   ```

The frontend application will be available at `http://localhost:3000`.

## API Endpoints

- `POST /evaluate`: Evaluate a poker hand
- `GET /recent-results`: Get recent game results

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
# Warehouse Inventory Management System

A full-stack warehouse inventory management application built to helps manage stock, track purchases and sales, and provides real-time inventory monitoring.

## Tech Stack

### Backend (warehouse-api)
- **Go** 1.24.6
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Primary database
- **JWT** - Authentication
- **Docker** - Containerization

### Frontend (warehouse-client)
- **Next.js** 16.1.6
- **React** 19.2.3
- **TypeScript** 5
- **Tailwind CSS** 4
- **Shadcn UI** - Component library
- **React Hot Toast** - Notifications

## Getting Started

### Option 1: Running with Docker (Recommended)

This is the easiest way to get started. Docker will handle all the dependencies for you.

#### 1. Clone the repository
```bash
git clone <your-repo-url>
cd "WAREHOUSE INVENTORY"
```

#### 2. Set up environment variables
Navigate to the backend directory and configure your environment:
```bash
cd warehouse-api
```

Create or edit the `.env` file with your configuration:
```env
SERVER=8000
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=warehouse
JWT_SECRET=your-secret-key-here
```

#### 3. Start the services
```bash
docker-compose up -d
```

This will start both the PostgreSQL database and the API server. The API will be available at `http://localhost:8000`.

#### 4. Set up the frontend
Open a new terminal and navigate to the frontend directory:
```bash
cd warehouse-client
```

Install dependencies:
```bash
npm install
```

Create a `.env` file:
```env
NEXT_PUBLIC_API_URL=http://localhost:8000
```

Start the development server:
```bash
npm run dev
```

The frontend will be available at `http://localhost:3000`.

---

### Option 2: Running Locally (Without Docker)

If you prefer to run everything locally without Docker:

#### 1. Set up PostgreSQL Database

First, make sure PostgreSQL is running on your machine. Then create the database:

```bash
psql -U postgres
CREATE DATABASE warehouse;
\q
```

#### 2. Set up the Backend

Navigate to the backend directory:
```bash
cd warehouse-api
```

Install Go dependencies:
```bash
go mod download
```

Configure the `.env` file:
```env
SERVER=8000
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your-password
POSTGRES_DB=warehouse
JWT_SECRET=your-secret-key-here
```

Run the API server:
```bash
go run main.go
```

The API will start on `http://localhost:8000`. GORM will automatically create the necessary tables on first run.

#### 3. Set up the Frontend

Open a new terminal and navigate to the frontend directory:
```bash
cd warehouse-client
```

Install dependencies:
```bash
npm install
```

Create a `.env` file:
```env
NEXT_PUBLIC_API_URL=http://localhost:8000
```

Start the development server:
```bash
npm run dev
```

Visit `http://localhost:3000` to see the application.

## API Documentation

Complete API documentation is available on Postman:

**[View API Documentation](https://documenter.getpostman.com/view/45402659/2sBXcBo2uV)**

The documentation includes:
- Authentication endpoints (register, login)
- Barang (items) management
- Stok (stock) tracking
- Pembelian (purchases) records
- Penjualan (sales) records
- Request/response examples
- Error handling


## Project Structure

```
WAREHOUSE INVENTORY/
├── warehouse-api/          # Backend API (Go)
│   ├── config/            # Configuration files
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # JWT & ACL middleware
│   ├── models/            # Database models
│   ├── repositories/      # Data access layer
│   ├── services/          # Business logic
│   ├── main.go           # Application entry point
│   ├── Dockerfile
│   └── docker-compose.yaml
│
└── warehouse-client/      # Frontend (Next.js)
    ├── app/              # Next.js app directory
    ├── components/       # React components
    ├── context/          # React context (Auth)
    ├── hooks/            # Custom React hooks
    ├── lib/              # Utility functions
    └── public/           # Static assets
```




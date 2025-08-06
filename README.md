# Rentacar API 🚗

Rentacar is a peer-to-peer car rental platform built as a RESTful API using Go. The application enables users to rent vehicles, host their own cars, or do both—offering flexibility for both customers and car owners. Designed with clean architecture and robust authentication, Rentacar handles user management, car listings, bookings, transactions, and more.

## 🔧 Features

- 🔐 JWT-based user authentication
- 🧾 Car listing and booking system
- 💳 User balance top-up and transaction tracking
- 🧠 Email verification via [Verifyright](https://verifyright.id/)
- 📦 Image/file upload using Supabase
- 🧪 Unit tests using Testify
- 📄 API documentation with Swagger
- ☁️ Deployment-ready for Heroku

## 🛠️ Tech Stack

- **Language**: Go
- **Framework**: Echo
- **Database**: PostgreSQL + GORM
- **Testing**: Testify
- **Docs**: Swagger
- **Storage**: Supabase
- **Deployment**: Heroku

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- Supabase account (for storage)
- Verifyright API key

### Environment Variables

Create a `.env` file and define:

JWT_SECRET=your_jwt_secret
DATABASE_URL=your_postgres_connection_url
VERIFYRIGHT_API_KEY=your_api_key
SUPABASE_PROJECT_URL=your_supabase_url
SUPABASE_API_KEY=your_supabase_api_key

### Run Locally

```bash
go mod tidy
go run main.go
```

API Routes
Here are some of the available endpoints:
| Method | Endpoint                    | Description                 |
| ------ | --------------------------- | --------------------------- |
| POST   | `/auth/register`            | Register a new user         |
| POST   | `/auth/login`               | Login and receive JWT token |
| GET    | `/users/me`                 | Get current user info       |
| POST   | `/users/topup`              | Top up balance              |
| POST   | `/cars`                     | Host a new car              |
| GET    | `/cars/available`           | View all rentable cars      |
| POST   | `/bookings`                 | Book a car                  |
| POST   | `/bookings/return`          | Return a booked car         |
| GET    | `/users/rentalhistory`      | Rental history              |
| GET    | `/users/transactionhistory` | Transaction history         |


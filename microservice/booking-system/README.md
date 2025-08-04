# Booking System Microservices (Go + Azure SQL + Docker)

This project is a simple microservices-based **Booking System** built with **Golang**, using **Clean Architecture**, **Gin**, **Azure SQL Server**, and **Docker**.

---

## 🧱 Architecture Overview

The system includes:

- **User Service**: Handles registration, login, JWT generation.
- **Booking Service**: Manages bookings of items and requires JWT authentication.

---

## 🧰 Tech Stack

| Layer        | Technology              |
|--------------|------------------------ |
| Language     | Golang (1.21+)          |
| Framework    | Gin                     |
| Auth         | JWT (Custom Middleware) |
| ORM          | GORM                    |
| Database     | Azure SQL Server        |
| Config       | Viper + YAML/ENV        |
| Container    | Docker                  |
| Orchestration | Docker Compose         |



## Authentication Flow
- User-Service authenticates user → returns JWT

- Booking-Service uses JWT middleware to authorize and inject user_id into context


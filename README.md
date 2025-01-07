# **E-commerce Dashboard Application**

A collaborative e-commerce dashboard application developed using **Go (Golang)**. This project provides an intuitive and efficient management system for e-commerce operations, focusing on modular architecture, secure authentication, and robust database interaction.

---

## **Features**

- **RESTful API**: Built with Golang to deliver efficient routing and functionality.
- **Modular Architecture**: Divided into multiple layers such as handler, service, and repository for scalability and ease of maintenance.
- **Relational Database**: Supports a well-structured relational database with optimized tables and relationships.
- **Authentication**: Secure access control implemented for user authentication.
- **Logging**: Effective monitoring and debugging with logging mechanisms.
- **CRUD Functionality**: Comprehensive Create, Read, Update, and Delete operations.
- **Clear Documentation**: Includes detailed setup instructions and an example `.env` file for seamless project configuration.
- **Database Design**: Designed an ERD (Entity Relationship Diagram) to support complex relationships.
- **Maintainability**: Improved code maintainability with proper directory organization and reusable components.

---

## **Main Contributions**

- Modular architecture implementation for better scalability.
- Database schema optimization with an ERD.
- Secure authentication and access control.
- Development of CRUD operations.
- Logging for effective monitoring.

---

## **Technologies Used**

- **Programming Language**: Golang
- **Database**: PostgreSQL
- **ORM**: GORM
- **Version Control**: Git

---

## **Project Setup**

### **Prerequisites**

Ensure you have the following installed:

- **Go (Golang)**: Version 1.18 or higher.
- **PostgreSQL**: Version 12 or higher.
- **Git**: For version control.

### **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/ecommerce-dashboard-golang.git
   cd ecommerce-dashboard-golang
   ```

2. Create a `.env` file in the project root and configure the following environment variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=ecommerce_db
   JWT_SECRET=your_jwt_secret
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run database migrations (if applicable):
   ```bash
   go run cmd/migrate.go
   ```

5. Start the application:
   ```bash
   go run main.go
   ```

---

## **Project Structure**

```
project-root/
├── cmd/               # Command-line specific tools (e.g., migrations)
├── config/            # Configuration files (e.g., .env parsing)
├── handlers/          # HTTP handlers for request routing
├── models/            # Database models and schemas
├── repository/        # Database interaction layer
├── service/           # Business logic layer
├── utils/             # Utility functions (e.g., token generation, logging)
├── main.go            # Application entry point
```

---

## **Usage**

1. Access the API at `http://localhost:8080` (default).
2. Use tools like **Postman** or **cURL** for testing the API endpoints.
3. Secure access by generating and using JWT tokens for authentication.

---

## **Future Enhancements**

- Implement a GraphQL API.
- Add real-time features using WebSockets.
- Enhance frontend integration with a React or Angular dashboard.
- Optimize database queries for large-scale operations.

---

Feel free to contribute to the project or report issues via the GitHub repository!


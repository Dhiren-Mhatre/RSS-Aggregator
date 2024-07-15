# RSS Aggregator

RSS Aggregator is a web application built using Go and PostgreSQL. It provides a platform for users to follow RSS feeds, collect posts from these feeds, and manage their subscriptions. The application is designed with a robust backend, utilizing the Chi router for handling HTTP requests and SQLC for database queries.

## Features

- **User Management**: Users can create accounts and retrieve their details.
- **RSS Feed Management**: Users can add, retrieve, and delete RSS feeds they follow.
- **Feed Follows**: Users can manage their subscriptions to various RSS feeds.
- **Post Management**: The system periodically fetches posts from subscribed RSS feeds and stores them in the database.
- **Health Check**: Endpoint to check the readiness of the application.

## Use Cases

1. **Personal News Aggregator**: Users can aggregate their favorite news sources into a single platform.
2. **Content Management**: Collect and organize content from multiple RSS feeds for easy access and reading.
3. **Learning Tool**: A practical example for learning Go and PostgreSQL, showcasing how to build a web application with these technologies.

## Technology Stack

- **Go**: The primary programming language used for the application.
- **PostgreSQL**: The database used for storing user, feed, and post data.
- **Chi Router**: A lightweight, idiomatic router for Go.
- **SQLC**: Generates type-safe code from SQL queries.
- **GoDotEnv**: Loads environment variables from a `.env` file.
- **CORS**: Middleware for handling Cross-Origin Resource Sharing.

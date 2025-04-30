# VOTE-BROADCAST-SERVER

A scalable microservice application for creating, managing, and conducting polls with real-time voting result updates. Built with modern technologies to ensure security, performance, and ease of deployment.

## Overview

**VOTE-BROADCAST-SERVER** simplifies the process of creating and managing polls, providing an interactive voting experience. The system supports secure user authentication, instant result updates, and efficient inter-service communication, making it ideal for online events, public voting, or internal surveys.

## Key Features

- Poll creation via API
- JWT-based authentication
- Real-time data updates (displaying polls and votes) using WebSockets
- Efficient inter-service communication with gRPC
- Containerization and deployment with Docker

## Tech Stack

- Go
- JWT
- PostgreSQL
- gRPC
- WebSockets
- Docker

## Getting Started

### Requirements

- Go
- Docker and Docker Compose
- Git

### Installation

```shell
git clone https://github.com/semesoff/vote-broadcast-server.git
cd vote-broadcast-server
```

### Build and Run

```shell
docker-compose up --build
```

### API Endpoints

Below are the main REST API and WebSocket endpoints with request examples.

#### REST API

**User Registration**
- `POST /api/register` - Register a new user
- Request Body
```json
{
  "username": "string",
  "password": "string"
}
```

**User Login**
- `POST /api/login` - Authenticate a user
- Request Body
```json
{
  "username": "string",
  "password": "string"
}
```

**Get List of Polls**
- `GET /api/polls` - Retrieve a list of all polls

**Get Poll Information**
- `GET /api/polls/{poll_id}` - Retrieve poll information by ID

**Create a Poll**
- `POST /api/polls` - Create a new poll
- Request Body
```json
{
  "title": "Poll Title",
  "options": [
    {
      "id": 0,
      "text": "Option 1"
    },
    {
      "id": 0,
      "text": "Option 2"
    }
  ],
  "type": 1
}
```
- Headers
```json
Authorization: Bearer <your_jwt_token>
```
- Notes
    - `id` in options should not be specified; it will be generated automatically.
    - `type` - poll type (1 - single choice, 2 - multiple choice).

**Get Votes**
- `GET /api/votes/{poll_id}` - Retrieve votes by poll ID

**Vote in a Poll**
- `POST /api/votes` - Submit a vote for a poll
- Request Body
```json
{
  "poll_id": 24,
  "options_id": [
    40
  ]
}
```
- Notes
    - If the poll type is 1, `options_id` can contain only one element.
    - If the poll type is 2, `options_id` can contain multiple elements.

#### WebSocket Endpoints

**Default Address** `ws://localhost:5005/`

**Get List of Polls**
- `/getPolls`

**Get Votes for a Specific Poll**
- `/getVotes/{poll_id}`

## Project Structure

```text
vote-broadcast-server/
├── init-scripts/ - Database initialization scripts
├── proto/ - Proto files for contracts
├── services/ - Services
│   ├── auth/ - Authentication service
│   ├── gateway/ - Service for handling HTTP requests
│   ├── poll/ - Poll management service
│   ├── vote/ - Voting service
│   └── websocket/ - WebSocket service
├── .env - Common configuration file for services
```
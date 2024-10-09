# Distributed File Storage Server

This project implements a distributed file storage system using Golang, PostgreSQL, and Docker. It allows files to be uploaded, split into smaller chunks, and stored in a database. Files can be retrieved by merging chunks back into the original file. The solution is Dockerized for ease of deployment.

## Features

1. **File Upload (API 1)**: Upload a file, split it into chunks, and store the chunks in a database.
2. **Get File Metadata (API 2)**: Retrieve metadata of uploaded file chunks.
3. **Download and Merge File (API 3)**: Retrieve and merge the file chunks into the original file.

## Requirements

- Golang 1.20+
- Docker & Docker Compose
- PostgreSQL

## API Endpoints

### 1. Upload File
- **URL**: `/upload`
- **Method**: `POST`
- **Request**: Multipart form-data with a `file` field.
- **Response**: Returns a unique `file_id` for the uploaded file.

### 2. Get File Metadata
- **URL**: `/files/{file_id}`
- **Method**: `GET`
- **Response**: Returns metadata of the file chunks stored in the database.

### 3. Download and Merge File
- **URL**: `/download/{file_id}`
- **Method**: `GET`
- **Response**: Returns the original file by merging all chunks.

## Running the Application

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/distributed-file-storage.git
cd distributed-file-storage
```

### 2. Build and Run using Docker Compose

```bash
docker-compose up --build
```
This will build and start the database and the application.

### 3. Access the Application

API server runs on http://localhost:8080
PostgreSQL database runs internally in Docker

### Database Schema
Table: file_chunks

id: Auto-incrementing primary key.
file_id: Unique identifier for each uploaded file.
chunk_number: Sequential number for each chunk.
chunk_data: The binary data of the chunk.

### License
Distributed under the MIT License. See LICENSE for more information.
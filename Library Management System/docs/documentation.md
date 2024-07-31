# Library Management System Documentation

## Introduction
The Library Management System is a comprehensive solution designed to streamline the management of a library's resources, including books, users, and lending activities. This system provides a user-friendly interface for librarians to efficiently handle all aspects of library operations.

## Installation Steps
1. Download and extract/clone the latest release of the Library Management System from the project's GitHub repository.

2. Run the file main.go which is found in the root folder.

## Features
- Add a new book to the library.
- Remove an existing book from the library.
- Borrow a book if it is available.
- Return a borrowed book.
- List all available books in the library.
- List all books borrowed by a specific member.

## Folder Structure
```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod

```
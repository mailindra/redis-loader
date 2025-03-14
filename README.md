# Redis Loader

A high-performance Redis data loader utility that generates and loads random key-value pairs into Redis using parallel processing.

## Features

- Parallel data loading with configurable number of workers
- Batch processing for improved performance
- Real-time progress tracking with ETA
- Random key-value pair generation
- Performance metrics reporting

## Prerequisites

- Go 1.23 or higher
- Redis server

## Installation

```bash
git clone [repository-url]
cd redis-loader
```

## Configuration

The application uses the following default configurations:
- Batch size: 1000 items per batch
- Number of parallel workers: 8

## Usage

Run the application:

```bash
go run main.go
```

The program will:
1. Prompt you to enter the number of key-value pairs to generate
2. Connect to Redis
3. Generate and load random data in parallel
4. Display real-time progress including:
    - Current progress percentage
    - Loading speed (items/second)
    - Estimated time remaining
5. Show performance summary after completion

## Performance

The loader utilizes several optimization techniques:
- Parallel processing with multiple workers
- Batch operations to reduce network overhead
- Buffered channels for efficient data transfer
- Atomic operations for accurate progress tracking

## Architecture

The project follows a clean architecture pattern with the following components:

- `main.go`: Application entry point and dependency injection
- `service/loader_service.go`: Core business logic for data loading
- Repository interface for Redis operations
- Console input handling
- Utility functions for random data generation

## License

MIT License

Copyright (c) 2024 mailindra.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
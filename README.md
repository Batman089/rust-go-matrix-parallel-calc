# rust-go-matrix-parallel-calc or Parallel Matrix Multiplication in Rust and Go

![Matrix Multiplication](https://img.shields.io/badge/Parallelism-Rust%20%26%20Go-blue)

This project demonstrates matrix multiplication implemented in **parallel** using **Rust** and **Go**, highlighting how both languages handle concurrency and parallelism. 
Matrix multiplication is a computationally intensive task, making it ideal for parallel processing. 
This repository provides code, benchmarks, and insights into how Rust and Go manage parallel computation on multi-core systems.

## Project Structure

- **Rust**: Located in the `rust/` directory, leveraging `std::thread` for parallel processing.
- **Go**: Located in the `go/` directory, utilizing Go's built-in goroutines and `sync.WaitGroup` for concurrent execution.

## Key Concepts

- **Concurrency in Rust**: Rust's strict ownership model ensures thread safety, with threads managed using `std::thread`.
- **Concurrency in Go**: Go’s goroutines and channels make lightweight concurrency easy, using `sync.WaitGroup` for synchronization.

## Features

- Parallel computation using native concurrency models in Rust and Go.
- Row-wise parallelism: Each row of the resulting matrix is computed independently across threads or goroutines.
- Benchmarking for performance comparison across languages and thread counts.

## Requirements

- **Rust**: Version 1.50+ (Install from [rust-lang.org](https://www.rust-lang.org/))
- **Go**: Version 1.16+ (Install from [golang.org](https://golang.org/))

# Unique IP Address Counter

### Overview
This Go program reads a list of IP addresses from a file (`ip_addresses.txt`), processes them concurrently using multiple worker goroutines, and calculates the number of unique IPv4 addresses. It utilizes Go's concurrency model with goroutines and channels for efficient parallel processing.

### Features
- Utilizes multiple CPU cores for concurrent processing.
- Uses a custom IP address set for storing and counting unique addresses.
- Provides memory usage statistics before and after execution.
- Reads input file line by line to avoid excessive memory usage.

### Requirements
- Go 1.18+
- A text file (`ip_addresses.txt`) containing one IP address per line.

### Installation
1. Clone the repository or download the Go script.
2. Place a valid `ip_addresses.txt` file in the same directory.
3. Ensure Go is installed on your system.

### Usage
Run the program using:
```sh
 go run main.go
```

### Code Structure
- `main.go`:
  - Opens the input file and initializes workers.
  - Distributes IP address processing across worker goroutines.
  - Aggregates results and prints the number of unique IPs.
- `Worker(id int, wg *sync.WaitGroup, lines <-chan string, result chan<- *ipAddressSet)`: 
  - Parses and processes IP addresses.
  - Stores unique IPv4 addresses.
- `PrintMemUsage(title string)`: 
  - Outputs memory usage statistics.
- `ipAddressSet`: 
  - Manages a map-based set of unique IP addresses.

### Performance
The program efficiently utilizes all available CPU cores using `runtime.NumCPU()`. By leveraging goroutines, it significantly reduces processing time for large input files.

### Why `sync.Map` is Not Used
We do not use `sync.Map` because it consumes more memory than a standard `map` with a `sync.Mutex`. `sync.Map` is designed for high-concurrency scenarios where frequent reads and writes occur, but in our case, we primarily store unique IPs with occasional writes. Using a regular `map` with a `sync.Mutex` allows for lower memory usage while still maintaining thread safety.

### Why `ips` is a Global Variable
The `ips` map is defined globally to ensure all worker goroutines have access to a shared data structure without requiring complex coordination. This approach simplifies aggregation and avoids excessive memory allocation during processing.

### License
This project is licensed under the MIT License.

### Author
[Albert Karapetyan](https://github.com/AlbertKarapetyan)
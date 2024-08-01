# Rate Limiter in Go

## Overview

This project implements a simple rate limiter in Go, designed to control the rate of requests from clients. The rate limiter allows up to a specified number of requests within a given time window, providing a mechanism to prevent abuse or overuse of a service.

## Features

- **Rate Limiting**: Configurable limit on the number of requests a client can make within a time window.
- **Concurrency Safe**: Uses mutex locks to ensure thread-safe operations.
- **Automatic Reset**: Resets request counts automatically after the time window expires.

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/diegodazpeitia/rate-limiter.git
cd rate-limiter
```

### Usage

1. **Initialize the Rate Limiter**:

   Create a new rate limiter instance by specifying the request limit and time window. For example, set a limit of 5 requests per 10 seconds.

2. **Allow Requests**:

   Use the `AllowRequest` method to check if a request from a client is allowed. Pass the client ID as a parameter to identify different clients. The method returns a boolean indicating whether the request is permitted.

3. **Example**:

   The project includes a sample application in the `main` package demonstrating how to use the rate limiter.

## Contributing

Contributions are welcome! Please submit a pull request with any changes, enhancements, or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

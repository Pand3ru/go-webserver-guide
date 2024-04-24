# Go Webserver Example

## Overview

This repository contains the source code for a simple web server written in Go. The server is designed to demonstrate basic web serving capabilities, middleware usage, and session management using the `gorilla/sessions` package. It assigns a unique user number to each visitor and retains this number across their sessions.

This project is part of a detailed tutorial available on my blog, where I explain the step-by-step process of building this server and the principles behind web servers and Go programming.

## Blog Post

For a comprehensive guide and explanation of the code, visit the blog post:
[Building a Webserver in Go](https://panderu.org/posts/webserver)

## Key Features

- Basic HTTP server that listens on port 8080.
- Session management to track unique user visits.
- Middleware to log requests and manage sessions.

## Note

Please note that this code includes an example session key (`"test"`) which should be replaced with a secure, randomly generated key in a production environment. The settings for session handling, such as cookie security and expiry, are configured for educational purposes and might need adjustment for real-world applications.

## License

This project is open-sourced under the MIT License. See the LICENSE file for more details.

## Acknowledgments

This project is created for educational purposes as part of a blog post on [panderu.org](https://panderu.org/).

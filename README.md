# Project Mycelium RESTful service

# Description
This project is a Golang-based RESTful web service using GIN,
serving as the intermediary between the path planner and the frontend. This
service efficiently retrieves information from the planner, processes it, and
prepares the data for seamless integration with the Vue.js frontend.

# Key Features
- RESTful API: Exposes a clean and efficient API for communication with the frontend.
- Data Processing: Gathers information from the path planner and formats it for real-time consumption by the frontend.
- Scalable and Performant: Built on GIN, known for its speed and minimalistic design, ensuring optimal performance.

# Getting Started
- Clone the repository: git clone https://github.com/Mycelium-Lauzhack2023/restapi
- Install dependencies: go get ./...
- Build and run the server: go run main.go
- The server will be accessible at http://localhost:8081

# API Endpoints
- GET /api/path: Retrieve path information from the path planner.
- Add any additional endpoints as needed.

# Configuration
- Modify the config.yaml file to adjust server settings, planner connection details, etc.

# Dependencies
- GIN: github.com/gin-gonic/gin
- Add any other dependencies to go.mod

# License
- This project is licensed under the [MIT License](LICENSE).

# Acknowledgments
- Mention any third-party libraries, tools, or frameworks used.

# Contact
- For inquiries and support, contact us at [your-email@example.com](mailto:your-email@example.com).

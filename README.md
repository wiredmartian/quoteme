### Project Overview

This project is a web application that allows users to select various services for their app development needs and receive an instant estimate based on their selections. The application is built using Go for the backend and serves a simple HTML frontend.

#### Features
- **Service Selection**: Users can choose from a variety of services, each with a detailed description and pricing information.
- **Instant Pricing Estimates**: As users select services, they receive real-time estimates of the total cost.
- **Responsive Design**: The application is designed to work seamlessly on both desktop and mobile devices.

#### Technologies Used
- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS (Tailwind CSS), JavaScript (HTMX)
- **Data Storage**: JSON file for service data

#### Getting Started
To run the application locally, follow these steps:
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/wiredmartian/quoteme.git
   cd quoteme
    ```
2. **Install Go**: Ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
3. **Run the Application**: Start the Go server:
   ```bash
   go run main.go
   ```
4. **Access the Application**: Open your web browser and go to `http://localhost:8080` to see the application in action.

#### Project Structure
- `main.go`: The main Go application file that sets up the server and routes.
- `services.json`: A JSON file containing the list of services, their descriptions, and pricing.
- `templates/index.html`: The HTML template for the frontend interface.


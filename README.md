# Event Ticket Reservation System


## Features
✅ **Reserve a Ticket**: Users can reserve a ticket by providing their name, email, and preferred event date.
✅ **View Ticket Details**: Retrieve ticket details using an email.
✅ **View All Attendees**: Get a list of attendees for a specific event date.
✅ **Cancel a Reservation**: Allows users to cancel their reservation.
✅ **Modify Seat Reservation**: Users can change their seat reservation.

## Tech Stack
- **Golang** (Gin Framework)
- **PostgreSQL** (Database)
- **GORM** (ORM)
## Setup Instructions
### **1. Clone the Repository**
```sh
git clone https://github.com/sushantshankar2414/assigment.git
cd event-ticket-reservation
```

### **2. Set Up the Database**
Ensure PostgreSQL is running and create a database:
```sh
createdb event_ticket_db
```

### **3. Configure Environment Variables**
Create a `.env` file in the root directory:
```sh
echo "DB_HOST=localhost
DB_USER="your_user" AS I HAVE CREATED SUSHANT NAME 
DB_PASSWORD=your_password
DB_NAME=event_ticket_db
DB_PORT=5432
DB_SSLMODE=disable" > .env
```

### **4. Install Dependencies**
```sh
go mod tidy FOR INSTALLING ALL DEPENDENCY
```

### **5. Run Database Migrations**
```sh
go run main.go
```

### **6. Start the Server**
```sh
go run main.go
```
Server will start on **`http://localhost:8080`**

## API Endpoints

FOR API ENDPOINTS REFER THIS DOCUMENTATION ON POSTMAN
https://documenter.getpostman.com/view/41001727/2sAYX9n13u ALL API IS PRESENT THERE 


# CRUD-GO
A simple CRUD (Create, Read, Update, Delete) application built with Golang.

## 🚀 Features
- Create, Read, Update, and Delete (CRUD) operations.
- Uses MySQL as the database.
- Environment variable configuration for database credentials.

---

## 🛠️ Setup Instructions

### Clone the Repository
```sh
git clone https://github.com/muhammadadna/CRUD-GO.git
cd CRUD-GO
```
### Create Table employee
```sh
CREATE TABLE `employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(255) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
```

#### Clone the Repository
# For define database 
```sh
export DB_USER="root"
export DB_PASSWORD="yourpassword"
export DB_HOST="127.0.0.1"
export DB_PORT="3306"
export DB_NAME="CRUD-GO"
```

# Install Dependencies
```sh
go mod tidy
```

# For running the app
```sh
go run main.go
```

## 📝 License
This project is licensed under the MIT License.

By: [@muhammadadna](https://github.com/muhammadadna)

Semangat!!
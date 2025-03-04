---
### User Management API
This project is a User Management API built with Go and the Gin framework. It provides endpoints to create, retrieve, update, and delete user records in a PostgreSQL database. The project uses GORM for ORM functionality and includes features such as random 3-digit user ID generation and automatic timestamping for user creation and updates.  

### Features
- **Create User:** Add a new user to the database.  
- **Retrieve Users:** Fetch all users or a specific user by ID.  
- **Update User:** Modify user details.  
- **Delete User:** Remove a user from the database.  
- **Random ID Generation:** Each user is assigned a unique 3-digit ID.  
- **Automatic Timestamps:** CreatedAt and UpdatedAt fields are automatically managed.  

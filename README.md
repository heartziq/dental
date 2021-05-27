# Installation Guide

This a booking system web app. It provides features for customers to book dental appointments. Below is a summary on how to install locally.

## Installation

1.  Open  [mysql workbench](https://dev.mysql.com/downloads/workbench/) and run the following query - *include column **'IsAdmin'** of type **boolean**  as well!*

```sql
CREATE database MYSTOREDB;
USE MYSTOREDB;
CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256), IsAdmin boolean);
```

2.  From root folder, Install dependencies *(ensure GO111MODULE=on)*

```bash
go mod tidy
```

3. Create a folder ```cert/```

   ```bash
   mkdir cert
   ```

4. Generate own server certificate and respective private keys

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout cert/key.pem -out cert/cert.pem
```

4. From the root folder, run
```bash
go run server.go
```

## Register Admin
Please ensure all cookies and caches have been cleared.

* Go to [https://localhost:8080/register](https://localhost:8080/register)
* Register: User: admin and Password: 123
* Go to **mysql workbench** and escalate user 'admin' privilege with the following query:
```sql
UPDATE Users 
SET 
    IsAdmin = 1
WHERE
    UserName = "admin";
```

* Logout
* [Re-login](https://localhost:8080/login) as admin *(now, you will see all appointments from all customers)* 
* You may proceed with registering normal users *i.e. the customers*

## Delete User (Admin only)
This feature requires you to log in as an admin

* Go to [https://localhost:8080/dashboard](https://localhost:8080/dashboard)
* A list of users will be displayed
* Click "Delete" right next to the username that you want to delete

## Book
Requires you to log in as a normal user. Admin can only manage existing appointments

* From homepage, navigate to Book Appointment sub menu
* Select a date
* And time - only a "full hour" is permitted by design *(As mentioned in GoInAction1)*

## Search for a particular user
Only admin can view appointments made by other customers
Search string has to be **exact** *i.e. if looking for "John Smith", Enter exact string with exact cases and whitespaces

* Log on as admin
* Enter a full, complete and case-sensitive (with whitespace if any)
* Click submit

*using binary search tree; only output 1 result per search. 

## Thank you
If anything, feel free to contact heartziq@gmail.com
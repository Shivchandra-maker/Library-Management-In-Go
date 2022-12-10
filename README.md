# Library Management

##### Rest Api for books and Author details

___

#### Author Details :

```
  ID        int
  FirstName  string 
  LastName  string 
  Dob       string 
  PenName   string 
  
  ```
 ___
  #### Book Details: 

``` 
  ID            int
  Title         string 
  Author        Author 
  Publication   string 
  PublishedDate string 
  
``` 
Get Books and Author details

##### DataBase used MySQL

Commands to create Database and tables

``` CREATE DATABASE test; ```

``` USE test; ```

##### TABLE COMMANDS

```  
CREATE TABLE Authors(
id int NOT NULL AUTO_INCREMENT,
first_name varchar(255) NOT NULL,
last_name varchar(255) NOT NULL,
dob varchar(255) NOT NULL,
pen_name varchar(255) NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE Books(
id int NOT NULL AUTO_INCREMENT,
title varchar(255) NOT NULL,
publication varchar(255) NOT NULL,
publication_date varchar(255) NOT NULL,
author_id int NOT NULL,
PRIMARY KEY (id),
FOREIGN KEY (author_id) REFERENCES Authors(id)
);
```





To Start Server 

``` go run main.go```




Table Books
```
+------------------+--------------+------+-----+---------+----------------+
| Field            | Type         | Null | Key | Default | Extra          |
+------------------+--------------+------+-----+---------+----------------+
| id               | int          | NO   | PRI | NULL    | auto_increment |
| title            | varchar(255) | YES  |     | NULL    |                |
| publication      | varchar(255) | YES  |     | NULL    |                |
| publication_date | varchar(255) | YES  |     | NULL    |                |
| author_id        | int          | YES  | MUL | NULL    |                |
+------------------+--------------+------+-----+---------+----------------+
```
Table Authors
```
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int          | NO   | PRI | NULL    | auto_increment |
| first_name | varchar(255) | YES  |     | NULL    |                |
| last_name  | varchar(255) | YES  |     | NULL    |                |
| dob        | varchar(255) | YES  |     | NULL    |                |
| pen_name   | varchar(255) | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```

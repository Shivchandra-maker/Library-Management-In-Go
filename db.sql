DROP TABLE IF EXISTS Author;
create table Author (
                        authorId INT ,
                        firstName varchar(50),
                        lastName varchar(50),
                        dob varchar(50),
                        penName     varchar(50),
                        PRIMARY KEY (AuthorId)
)

DROP TABLE IF EXISTS Books;
CREATE TABLE Books(
                      bookId VARCHAR(10) ,
                      authorId INT,
                      title  VARCHAR(50),
                      Publications VARCHAR(50),
                      PublishedDate VARCHAR(50),
                      PRIMARY KEY (bookId),
                      FOREIGN KEY (authorId) REFERENCES Author(authorId)
)
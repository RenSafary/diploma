CREATE TABLE clients
(
    Id SERIAL PRIMARY KEY,
    FirstName VARCHAR(100),
    LastName VARCHAR(100),
    Email VARCHAR(250),
    Age INTEGER,
    Sex CHAR(1)
);
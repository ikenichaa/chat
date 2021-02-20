# chat
# Start project by running
```
docker-compose up
```

# mysql
username = user
password = password

# create table
CREATE TABLE chat (
    id int NOT NULL,
    name varchar(255) NOT NULL,
    message varchar(255),
    create_at datetime,
    PRIMARY KEY (ID)
);
# Backend
Insert user and message into DB
http://localhost:9999/chat/message (Post method)
{
    "user": "amy",
    "message": "hello ja"
}
http://localhost:9999/chat/message (Get method)
To get all message from DB

# Open front end:
http://localhost:8080/
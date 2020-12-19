CREATE TABLE IF NOT EXISTS projects (
    id serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    live_url VARCHAR(50) ,
    github VARCHAR(50)

)


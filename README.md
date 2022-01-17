# Allied/AirTrail Go API Demo

## Introduction

This GO program exposes an API that provides the following endpoints:
1. All Aircrafts
2. Aircrafts Asc Flight Hours
3. Aircrafts Desc Flight Hours
4. Aircrafts Filter by Type
5. Aircrafts Export CSV
6. Aircrafts Filter by Registration
7. Aircraft by Aircraft Id

The POSTMAN collection is also part of the repository.

## Setup

In order to run the program, the root folder should contain a .env file with the following properties

```azure
SEED_DATA=false
GENERATE_SCHEMA=false
LISTEN_SERVER=true
LISTEN_PORT=3000
PGHOST=127.0.0.1
PGPORT=5432
PGUSER=usman
PGDBNAME=databasename
PGPASSWORD=******
PGSSLMODE=disable
```

## Used Packages
Following are the only used packages that are not part of the GO standard library

 - github.com/lib/pq ([posgresql driver](https://github.com/lib/pq))
 - github.com/joho/godotenv ([for reading environment variables](https://github.com/joho/godotenv))
 - github.com/google/uuid ([for support of UUID in database columns](https://github.com/google/uuid))
 - entgo.io/ent ([go entity framework by Google](https://entgo.io))
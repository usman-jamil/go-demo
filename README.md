# Allied/AirTrail Go API Demo

## Introduction

`This program was written to showcase experience with core GoLang. Use of 3rd party packages and frameworks was avoided as much as possible`

This GO program exposes an API that provides the following endpoints:
1. All Aircrafts
2. All Aircrafts Paged
3. Aircrafts Asc Flight Hours
4. Aircrafts Desc Flight Hours
5. Aircrafts Filter by Type
6. Aircrafts Export CSV
7. Aircrafts Filter by Registration
8. Aircraft by Aircraft Id

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

Create an instance of postgres on your dev environment by running the following docker command. You will need to install docker desktop for this.

`docker run --name postgresql-container -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -d postgres`

In the project root directory install the following dependencies:
1. `go get github.com/joho/godotenv`
2. `go get -d entgo.io/ent/cmd/ent`
3. `go get github.com/google/uuid`

## Generate Dummy Data

The following Postgres SQL script generates 100,000 dummy records in the aircraft table

```
insert into aircrafts (aircraft_id, company_id, current_flight_hours, current_cycles, aircraft_registration,
                           base_airport_code, manufacturer, manufacturer_designator, common_designation, common_name,
                           pilots_required_to_fly, default_values, maximum_values, current_landings, fuel_details,
                           oil_details)
SELECT gen_random_uuid(), gen_random_uuid(), CAST(n as real) as current_flight_hours, n as current_cycles, 'aircraft_registration_' || n as aircraft_registration, 'base_airport_code_' || n as base_airport_code
       , 'manufacturer_' || n as manufacturer, 'manufacturer_designator_' || n as manufacturer_designator, 'common_designation_' || n as common_designation, 'common_name_' || n as common_name
       , n as pilots_required_to_fly, 'default_values_' || n as default_values, 'maximum_values_' || n as maximum_values, n
       , 'fuel_details_' || n as fuel_details, 'oil_details_' || n as oil_details
FROM generate_series(1, 100000) as n
```

## Used Packages
Following are the only used packages that are not part of the GO standard library

 - github.com/lib/pq ([posgresql driver](https://github.com/lib/pq))
 - github.com/joho/godotenv ([for reading environment variables](https://github.com/joho/godotenv))
 - github.com/google/uuid ([for support of UUID in database columns](https://github.com/google/uuid))
 - entgo.io/ent ([go entity framework by Google](https://entgo.io))

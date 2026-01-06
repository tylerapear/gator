# GATOR

A RSS feed aggregator build in Go.

Add RSS feeds to a "followed" list and run an aggregator in background that continuously fetches followed feeds and stores them in a SQL database. Browse the latest posts by running `gator browse`.

## Dependancies

- Go version 1.25.3 or later
- PostgreSQL 16.11 or later
- sqlc verison 1.30.0 or later
- goose version 3.26.0 or later

## Getting Started

1) Clone the repo:

    ```
    git clone https://github.com/tylerapear/gator.git
    ```

2) Cd into the project directory:

    ```
    cd gator
    ```

3) Create a PostgreSQL database called `gator`:

4) From the project directory, use goose to run all up migrations:

   ```
   goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
   ```

5) From the project directory, install the Go program:

    ```
    go install .
    ```

6) Set up the gator config file with your database connection:

   ```
   gator init postgres://postgres:postgres@localhost:5432/gator
   ```

7) Register a user:

   ```
   gator register <username>
   ```

8) Use any of the following commands!

## Commands
|Command|Description|Usage|
|-------|-----------|-----|
|addfeed|Add an RSS feed to the database and follow it|`gator addfeed <name> <url>`|
|agg|Run an aggregating service that continuously fetches updates to all feeds in the database|`gator agg <time_between_requests (eg. 1s, 1m, 1h, etc.)>`|
|browse|Browse posts aggregated from followed feeds|`gator browse [limit]`|
|feeds|List all feeds in the database|`gator feeds`|
|follow|Follow a feed|`gator follow <url>`|
|following|List feeds followed by the logged in user|`gator following`|
|login|Log in as a registered user|`gator`|
|register|Register a new user and log in as them|`gator register <username>`|
|reset|Reset all data in the database|`gator reset`|
|unfollow|Unfollow a feed|`gator unfollow <url>`|
|users|List all registered users|`gator users`|

# to-do-api

Description
----------- 

REST API for creating task lists in Go.
 
Installation instructions
--------------------------

To run, you need to install [Docker](https://docs.docker.com/).

Operating Instructions
-----------------------

Run application:

```
make build & make up
```

On first run, you need to apply migrations to database:

```
make migrate-up
```

Documentation
-------------

[Swagger docs](http://localhost:8080/swagger/index.html) for API will be available upon launch.

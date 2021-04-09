FROM mysql:8.0.15
COPY ./initialize.sql /docker-entrypoint-initdb.d/1-initialize.sql
COPY ./inittable.sql /docker-entrypoint-initdb.d/2-initdb.sql
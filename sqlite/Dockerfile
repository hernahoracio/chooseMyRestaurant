FROM apline:latest

RUN apk add --no-cache sqlite
WORKDIR /db

COPY schema.sql /db/schema.sql

CMD ["sqlite3", "chooseMyRestaurant.db"]

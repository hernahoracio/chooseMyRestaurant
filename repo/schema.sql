CREATE TABLE Restaurants (
  ID Integer Primary Key Autoincrement,
  Name text Not Null,
  Address text,
  Latitude float Not Null,
  Longitude float Not Null,
  cuisine text,
  rating float
)

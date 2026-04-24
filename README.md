# chooseMyRestaurant

chooseMyRestaurant finds nearby restaurants using:

- Your current coordinates from `CoreLocationCLI` (this currently only works for MacOS lol) 
- Google Places API (Nearby Search)
- Next feature will be to write the API response to a sqlite server and filter based on your current needs. So if you only want chinese food near you can filter that. (Yeah I could ask the Google API for that but I'm trying to use sqlite in a project)

## Project structure

```text
.
├── cmd/
│   ├── main.go
│   ├── config.go
│   └── Dockerfile
├── internal/
│   ├── coords.go
│   └── places.go
├── models/
│   └── model.go
├── repo/
│   ├── repository.go
│   ├── errors.go
│   ├── dto/restaurant.go
│   └── schema/schema.sql
├── server/
│   ├── server.go
│   └── restaurants/create.go
├── sqlite/
│   └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
```

## Requirements

- Go installed (1.26 or greater)
- A Google API key with Places API access
- `CoreLocationCLI` installed

## 1) Configure environment variables

1. Copy `.env.example` to `.env`.
2. Replace `GOOGLE_API_KEY` with your real key. You can also add more parameters if you want. (Take a look at the Google Places API docs for that)

Example:

```env
GOOGLE_API_KEY=your_real_google_api_key_here
PlacesBaseURL=https://places.googleapis.com/v1/places:searchNearby
Parameters=places.displayName,places.formattedAddress,places.types
```

## 2) Run the app

From repo root:

```bash
go run ./cmd
```

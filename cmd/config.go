package main

type Config struct {
    PlacesBaseURL string `envconfig:"PlacesBaseURL" default:"https://places.googleapis.com/v1/places:searchNearby"`
    API_Key string `envconfig:"API_KEY" required:"true"`
    Parameters []string `envconfig:"Parameters" default: "places.displayName,places.formattedAddress,places.types"`
}

func loadConfig() (Config, error) {
    var cfg Config
    if err := envconfig.Process("", &cfg); err != nil {
        return Config{}, fmt.Errorf("load config: %w", err) 
      }
      return cfg, nil
  }

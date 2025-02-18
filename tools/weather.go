package tools

import (
	"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
)

type WeatherTool struct {
	apiKey string
}

func NewWeatherTool() *WeatherTool {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")

	return &WeatherTool{

		apiKey: apiKey,
	}
}

func (wt *WeatherTool) Description() string {
	return "A tool to provide the current weather forecast"
}

func (wt *WeatherTool) Invoke() error {
	w, err := owm.NewCurrent("C", "EN", wt.apiKey)
	if err != nil {
		return fmt.Errorf("creating weather client: %w", err)
	}

	err = w.CurrentByName("London, UK")
	if err != nil {
		return fmt.Errorf("getting weather data: %w", err)
	}

	// TODO: Return instead
	fmt.Printf("Current weather in London: %v\n", w)
	return nil
}

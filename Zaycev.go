package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"encoding/json"
)

func main() {
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8080", nil)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Get the track ID from the request parameters
	trackID := r.URL.Query().Get("track_id")

	// Get the track URL using Zaycev.net API or any other method
	trackURL, err := getTrackURL(trackID)
	if err != nil {
		http.Error(w, "Failed to get track URL", http.StatusInternalServerError)
		return
	}

	// Download the track
	err = downloadTrack(trackURL, trackID)
	if err != nil {
		http.Error(w, "Failed to download track", http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprintf(w, "Track downloaded successfully")
}

func getTrackURL(trackID string) (string, error) {
	// Implement your logic to get the track URL from Zaycev.net API
	// Make the necessary API requests and extract the track URL from the response
	// Return the track URL or an error if something goes wrong
	// Here's an example implementation using the Zaycev.net search API:

	// Make a GET request to the search API
	searchURL := fmt.Sprintf("https://zaycev.net/api/search?type=tracks&q=%s", trackID)
	response, err := http.Get(searchURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Parse the response JSON to extract the track URL
	// Implement your JSON parsing logic based on the Zaycev.net API response structure
	// Here's an example assuming the response contains a "data" field with an array of tracks
	// and each track object has a "url" field containing the track URL
	// Adjust the parsing logic based on the actual response structure
	var searchResponse struct {
		Data []struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	err = json.NewDecoder(response.Body).Decode(&searchResponse)
	if err != nil {
		return "", err
	}

	// Check if any tracks were found
	if len(searchResponse.Data) > 0 {
		return searchResponse.Data[0].URL, nil
	}

	return "", fmt.Errorf("Track not found")
}

func downloadTrack(trackURL, trackID string) error {
	// Create a new file to save the downloaded track
	file, err := os.Create(trackID + ".mp3")
	if err != nil {
		return err
	}
	defer file.Close()

	// Send a GET request to the track URL
	response, err := http.Get(trackURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}





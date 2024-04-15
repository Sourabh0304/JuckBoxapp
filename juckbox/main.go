package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"juckbox/common"
	//"sync"
	//"time"
)

var store = common.DataStore{
	Albums:    []*common.MusicAlbum{},
	Musicians: []*common.Musician{},
}

// CreateOrUpdateAlbumHandler creates or updates a music album record.
func CreateOrUpdateAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var album common.MusicAlbum

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	log.Printf("Request Body: %s", string(body))

	// Use json.Unmarshal to decode the JSON data into the struct
	if err := json.Unmarshal(body, &album); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Validate album fields
	if len(album.Name) < 5 {
		http.Error(w, "Album name should be minimum 5 characters", http.StatusBadRequest)
		return
	}
	if album.ReleaseDate.IsZero() {
		http.Error(w, "Release date is required", http.StatusBadRequest)
		return
	}
	if album.Price < 100 || album.Price > 1000 {
		http.Error(w, "Price should be between 100 and 1000", http.StatusBadRequest)
		return
	}

	// Generate ID for the album
	album.ID = store.GetNextAlbumID()

	// Add/update album in the database
	store.Albums = append(store.Albums, &album)
	fmt.Fprintf(w, "Album created/updated successfully")
}

// CreateOrUpdateMusicianHandler creates or updates a musician record.
func CreateOrUpdateMusicianHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var musician common.Musician
	err := json.NewDecoder(r.Body).Decode(&musician)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate musician fields
	if len(musician.Name) < 3 {
		http.Error(w, "Musician name should be minimum 3 characters", http.StatusBadRequest)
		return
	}

	// Generate ID for the musician
	musician.ID = store.GetNextMusicianID()

	// Add/update musician in the database
	store.Musicians = append(store.Musicians, &musician)
	fmt.Fprintf(w, "Musician created/updated successfully")
}

// GetAlbumsByReleaseDateHandler retrieves the list of music albums sorted by release date.
func GetAlbumsByReleaseDateHandler(w http.ResponseWriter, r *http.Request) {
	// Sort albums by release date
	sort.Slice(store.Albums, func(i, j int) bool {
		return store.Albums[i].ReleaseDate.Before(store.Albums[j].ReleaseDate)
	})

	// Convert albums to JSON and send response
	jsonData, err := json.Marshal(store.Albums)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// GetAlbumsByMusicianHandler retrieves the list of music albums for a specified musician sorted by price.
func GetAlbumsByMusicianHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameter
	musicianID := r.URL.Query().Get("musician_id")
	if musicianID == "" {
		http.Error(w, "Missing musician ID", http.StatusBadRequest)
		return
	}

	// Find albums for the specified musician
	var albums []*common.MusicAlbum
	for _, album := range store.Albums {
		// Implement your logic to find albums by musician
		albums = append(albums, album)
	}

	// Sort albums by price
	sort.Slice(albums, func(i, j int) bool {
		return albums[i].Price < albums[j].Price
	})

	// Convert albums to JSON and send response
	jsonData, err := json.Marshal(albums)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// GetMusiciansByAlbumHandler retrieves the list of musicians for a specified music album sorted by musician's name.
func GetMusiciansByAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameter
	albumID := r.URL.Query().Get("album_id")
	if albumID == "" {
		http.Error(w, "Missing album ID", http.StatusBadRequest)
		return
	}

	// Find musicians for the specified album
	var musicians []*common.Musician
	for _, musician := range store.Musicians {
		// Implement your logic to find musicians by album
		musicians = append(musicians, musician)
	}

	// Sort musicians by name
	sort.Slice(musicians, func(i, j int) bool {
		return musicians[i].Name < musicians[j].Name
	})

	// Convert musicians to JSON and send response
	jsonData, err := json.Marshal(musicians)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	// Register API handlers
	http.HandleFunc("/albums", CreateOrUpdateAlbumHandler)
	http.HandleFunc("/musicians", CreateOrUpdateMusicianHandler)
	http.HandleFunc("/albums/sortedbyreleasedate", GetAlbumsByReleaseDateHandler)
	http.HandleFunc("/albums/sortedbypriceformusician", GetAlbumsByMusicianHandler)
	http.HandleFunc("/musicians/sortedbyalbum", GetMusiciansByAlbumHandler)

	// Start the server
	fmt.Println("Server started on port 8082")
	http.ListenAndServe(":8082", nil)
}


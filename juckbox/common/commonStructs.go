package common

import (

	 "sync"
	 "time"
)


// MusicAlbum represents a music album.
type MusicAlbum struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       string    `json:"genre"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}

// Musician represents a musician.
type Musician struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MusicianType string `json:"musician_type"`
}

// DataStore simulates a database.
type DataStore struct {
	Albums    []*MusicAlbum
	Musicians []*Musician
	mu        sync.Mutex
	albumID   int
	musicianID int
}



func (ds *DataStore) GetNextAlbumID() int {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.albumID++
	return ds.albumID
}

func (ds *DataStore) GetNextMusicianID() int {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.musicianID++
	return ds.musicianID
}

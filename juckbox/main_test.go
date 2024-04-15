package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrUpdateAlbumHandler(t *testing.T) {
	payload := []byte(`{
		"name": "Test Album",
		"release_date": "2022-04-11",
		"genre": "Pop",
		"price": 500,
		"description": "Test Description"
	}`)

	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrUpdateAlbumHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
}

func TestCreateOrUpdateMusicianHandler(t *testing.T) {
	payload := []byte(`{
		"name": "Test Musician",
		"musician_type": "Vocalist"
	}`)

	req, err := http.NewRequest("POST", "/musicians", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateOrUpdateMusicianHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
}

func TestGetAlbumsByReleaseDateHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/albums/sortedbyreleasedate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAlbumsByReleaseDateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
}

func TestGetAlbumsByMusicianHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/albums/sortedbypriceformusician?musician_id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAlbumsByMusicianHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
}

func TestGetMusiciansByAlbumHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/musicians/sortedbyalbum?album_id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMusiciansByAlbumHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


}

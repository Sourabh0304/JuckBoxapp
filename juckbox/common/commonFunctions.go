package common

import (
	"encoding/json"
	 "time"
)

func (a *MusicAlbum) UnmarshalJSON(data []byte) error {
    type Alias MusicAlbum
    aux := &struct {
        ReleaseDate string `json:"release_date"`
        *Alias
    }{
        Alias: (*Alias)(a),
    }
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    releaseDate, err := time.Parse("2006-01-02", aux.ReleaseDate)
    if err != nil {
        return err
    }
    a.ReleaseDate = releaseDate
    return nil
}

// Custom JSON marshaller for the ReleaseDate field
func (a *MusicAlbum) MarshalJSON() ([]byte, error) {
    type Alias MusicAlbum
    return json.Marshal(&struct {
        ReleaseDate string `json:"release_date"`
        *Alias
    }{
        ReleaseDate: a.ReleaseDate.Format("2006-01-02"),
        Alias:       (*Alias)(a),
    })
}

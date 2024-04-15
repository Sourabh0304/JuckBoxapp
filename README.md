# JuckBoxapp


Comiling the code
  cd juckbox
  go mod tidy
  go build -o appName JuckBoxApp.go
To run the binary:
  ./appName


For running the test use below commond
    go test
Note: It will start on port 8082

API:
1. Api to create/Update Music Album
   POST Method with below url:
   http://localhost:8082/albums
   Json Body:

   {
    "name": "album1",
    "release_date": "2024-04-12",
    "genre": "Pop",
    "price": 250,
    "description": "album 1 is test one"
}

2. Api to create/Update Musician:
   POST Method with below url:
   http://localhost:8082/musicians
   Json Body:

   {
    "name": "Badshah",
    "musician_type": "rapper"

  } 
 3. Api to get albums sorted by release date
     GET Method with below url:
     http://localhost:8082/albums/sortedbyreleasedate

 4. API to get albums sorted by price for musician with musician ID
    GET METHOD with below url
     http://localhost:8082/albums/sortedbypriceformusician?musician_id=1
 5. API to get musician sorted by album ID
     GET METHOD with below URL
    http://localhost:8082/musicians/sortedbyalbum?album_id=1

       

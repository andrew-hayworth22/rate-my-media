package media

import "time"

type DbMediaType struct {
	Id   int
	Name string
}

var MEDIA_TYPE_MOVIE = DbMediaType{
	Id:   1,
	Name: "Movie",
}
var MEDIA_TYPE_VIDEO_GAME = DbMediaType{
	Id:   2,
	Name: "Video Game",
}
var MEDIA_TYPE_BOOK = DbMediaType{
	Id:   3,
	Name: "Book",
}
var MEDIA_TYPE_TV_SHOW = DbMediaType{
	Id:   4,
	Name: "TV Show",
}

type DbMedia struct {
	Id          int
	MediaType   DbMediaType
	Name        string
	Description string
	ReleaseDate time.Time
}

type DbMovie struct {
	Media          DbMedia
	RuntimeMinutes int
}

type DbVideoGame struct {
	Media DbMedia
}

type DbBook struct {
	Media DbMedia
	Pages int
}

type DbTVShow struct {
	Media                 DbMedia
	EpisodeRuntimeMinutes int
}

type DbStoreMovieRequest struct {
	Name           string
	Description    string
	ReleaseDate    time.Time
	RuntimeMinutes int
}

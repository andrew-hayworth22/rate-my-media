package database

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
var MEDIA_TYPE_ALBUM = DbMediaType{
	Id:   5,
	Name: "Album",
}

type DbMedia struct {
	Id          int         `db:"id"`
	MediaType   DbMediaType `db:"-"`
	Name        string      `db:"name"`
	Description string      `db:"description"`
	ReleaseDate time.Time   `db:"release_date"`
}

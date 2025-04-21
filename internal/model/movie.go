package model

// this format of capital lettered names in struct
// is the standard for exported fields in Go to json
// small lettered names don't get exposed
// also after defining a propery and its type we specify that in json they will be reffered as
// the name we are giving now
// Title  string `json:"title"` means that Title property in go will be title when converted
// to json
type Movie struct {
	ID       string    `json:"id"`
	Imdb_id  string    `json:"imdb_id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type MovieRequest struct {
	Title    string    `json:"title"`
	Imdb_id  string    `json:"imdb_id"`
	Director *Director `json:"director"`
}

// slice or array of type Movie struct
var Movies = []Movie{
	{
		ID:      "52420926a8d403",
		Imdb_id: "tt1375666",
		Title:   "Inception",
		Director: &Director{ // we are passing address of struct cause we want to modify actual struct
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:      "3fd26c41ae733f",
		Imdb_id: "tt0133093",
		Title:   "The Matrix",
		Director: &Director{
			Firstname: "Lana",
			Lastname:  "Wachowski",
		},
	},
	{
		ID:      "9074185900698c",
		Imdb_id: "tt0468569",
		Title:   "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:      "e2fccb318304cf",
		Imdb_id: "tt1285016",
		Title:   "The Social Network",
		Director: &Director{
			Firstname: "David",
			Lastname:  "Fincher",
		},
	},
}

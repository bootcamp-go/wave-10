package internal

var Genres = []string{"Action", "Adventure", "Fighting", "Platform", "Puzzle", "Racing", "Role-Playing", "Shooter", "Simulation", "Sports", "Strategy"}

// VideoGame is a struct that represents a video game
type VideoGame struct {
	ID       int
	Name     string
	Platform string
	Year     int
	Genre    string
}
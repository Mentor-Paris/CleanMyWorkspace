package CleanMyWorkspace

import "math/rand"

// waste is a list of waste
var waste = []string{
	"Canette de soda",
	"Papier de brouillon",
	"Emballage de fast-food",
	"Paquet de confiserie",
	"Goblet",
	"Emballage de sandwich",
	"Mouchoir en papier",
	"Reste de nuggets",
	"os d'oscar",
	"git a la racine",
}

// randomWaste return a random waste
func randomWaste() *string {
	t := new(string)
	*t = waste[rand.Intn(len(waste))]
	return t
}

package peruvianslangs

import (
	"math/rand"
)

type Slang struct {
	slang   string
	meaning string
}

var listPeruvianSlangs = []Slang{
	{slang: "Asu mare", meaning: "Similar a wow o güau"},
	{slang: "Estar chihuan", meaning: "No tener dinero"},
	{slang: "Estar misio", meaning: "No tener dinero"},
	{slang: "Estar aguja", meaning: "No tener dinero"},
	{slang: "Qué palta!", meaning: "Qué vergüenza!"},
	{slang: "Qué roche!", meaning: "Qué vergüenza!"},
	{slang: "Piña, pues", meaning: "Mala suerte"},
	{slang: "Qué piña", meaning: "Qué mala suerte"},
	{slang: "Causa", meaning: "Amigo"},
}

func GetRandomPeruvianSlang() string {
	randomIndex := rand.Intn(len(listPeruvianSlangs))
	return listPeruvianSlangs[randomIndex].slang
}


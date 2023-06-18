package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	SIZE    int    = 5
	SYMBOLS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-_+"
)

// gera a semente para números aleatórios
func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateId() string {
	newId := func() string {
		id := make([]byte, SIZE, SIZE)
		for i := range id {
			id[i] = SYMBOLS[rand.Intn(len(SYMBOLS))]
		}
		return string(id)
	}
	for {
		if id := newId(); !repo.ExistId(id) {
			return id
		}
	}
}

// padroniza o uso do repositorio em memoria
type Repository interface {
	ExistId(id string) bool
	SearchById(id string) *Url
	SearchByUrl(url string) *Url
	Save(url Url) error
}

/*
variavel visivel apenas no pacote, sendo possivel
alterar sua configuracao de repositorio
apenas pela funcao Setting Repository
*/
var repo Repository

func SettingRepository(r Repository) {
	repo = r
}

// tipo url para armazenar o identificador curto
type Url struct {
	Id      string    `json:"id"`
	Created time.Time `json:"created"`
	Dest    string    `json:"dest"`
}

func SearchOrCreateAnewUrl(urlDest string) (u *Url, new bool, err error) {
	if u = repo.SearchByUrl(urlDest); u != nil {
		return u, false, nil
	}

	if _, err := url.ParseRequestURI(urlDest); err != nil {
		return nil, false, err
	}

	url := Url{Id: generateId(),
		Created: time.Now(),
		Dest:    urlDest}

	repo.Save(url)
	return &url, true, nil
}

func Search(id string) *Url {
	return repo.SearchById(id)
}

package handlers

import (
	"encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "pokedex/db"
    "pokedex/models"
)

func GetPokemons(w http.ResponseWriter, r *http.Request){
	rows, err := db.DB.Query("SELECT id, nome, tipo, tamanho, peso, descricao, rotas FROM pokemons")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pokemons []models.Pokemon
	for rows.Next() {
		var p models.Pokemon
		if err := rows.Scan(&p.ID, &p.Nome, &p.Tipo, &p.Tamanho, &p.Peso, &p.Descricao, &p.Rotas); 
		err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pokemons = append(pokemons, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func getPokemon(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err	:= strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Pokemon ID", http.StatusBadRequest)
		return
	}

	var.models.Pokemon
	err =db.DB.QueryRow("SELECT id, nome, tipo, tamanho, peso, descricao, rotas FROM pokemons WHERE id = $1", id).Scam
}
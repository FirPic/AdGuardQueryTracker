package world

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
    "net/http"
)

// RegisterHandlers enregistre les gestionnaires d'API pour la catégorie World.
func (w *World) RegisterHandlers() {
    if w.conf.HTTPRegister == nil {
        return
    }

    w.conf.HTTPRegister(http.MethodGet, "/control/world/ip", w.handleWorldRequestIpPerCountry)
    w.conf.HTTPRegister(http.MethodGet, "/control/world/count", w.handleWorldRequestCountPerCountry)
}

// handleWorldRequestCountryIP gère les requêtes pour obtenir des informations sur les IP par pays.
func (w *World) handleWorldRequestIpPerCountry(wr http.ResponseWriter, r *http.Request) {
    // Exemple de réponse pour les informations sur les IP par pays
    response := map[string]string{
        "message": "Information sur les IP par pays non encore implémentée",
    }
    wr.Header().Set("Content-Type", "application/json")
    wr.WriteHeader(http.StatusOK)
    json.NewEncoder(wr).Encode(response)
}

// handleWorldRequestCountryCount gère les requêtes pour obtenir le nombre de pays.
func (w *World) handleWorldRequestCountPerCountry(wr http.ResponseWriter, r *http.Request) {
    // Exemple de réponse pour le nombre de pays
    response := map[string]int{
        "country_count": 195, // Exemple fixe, mettre à jour avec une valeur dynamique si nécessaire
    }
    wr.Header().Set("Content-Type", "application/json")
    wr.WriteHeader(http.StatusOK)
    json.NewEncoder(wr).Encode(response)
}

// getQueryLog lit les logs des fichiers de requêtes et les affiche.
func getQueryLog() {
    // Liste des fichiers de log de requêtes
    files := []string{
        "data/querylog.json",
    }

    for _, filePath := range files {
        file, err := os.Open(filePath)
        if err != nil {
            log.Fatalf("Erreur lors de l'ouverture du fichier de log : %v", err)
        }
        defer file.Close()

        decoder := json.NewDecoder(file)
        for {
            var logEntry map[string]interface{}
            if err := decoder.Decode(&logEntry); err == io.EOF {
                break
            } else if err != nil {
                log.Fatalf("Erreur lors de la lecture du fichier de log : %v", err)
            }
            fmt.Println(logEntry)
        }
    }
}

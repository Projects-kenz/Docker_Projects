package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type Mapping struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

var (
	mappings []Mapping
	mu       sync.Mutex
	storeFile = "store.json"
)

func loadMappings() {
	file, err := os.ReadFile(storeFile)
	if err == nil {
		json.Unmarshal(file, &mappings)
	}
}

func saveMappings() {
	file, _ := json.MarshalIndent(mappings, "", "  ")
	_ = os.WriteFile(storeFile, file, 0644)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	for _, m := range mappings {
		if m.Short == short {
			http.Redirect(w, r, m.Long, http.StatusFound)
			return
		}
	}

	http.NotFound(w, r)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("✅ Running"))
}

func main() {
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "add":
			if len(os.Args) < 4 {
				fmt.Println("Usage: add <short> <long>")
				return
			}
			short, long := os.Args[2], os.Args[3]
			loadMappings()
			mappings = append(mappings, Mapping{Short: short, Long: long})
			saveMappings()
			fmt.Printf("✅ Added: %s → %s\n", short, long)
			return
		case "list":
			loadMappings()
			for _, m := range mappings {
				fmt.Printf("%s → %s\n", m.Short, m.Long)
			}
			return
		default:
			fmt.Println("Unknown command")
			return
		}
	}

	loadMappings()

	r := mux.NewRouter()
	r.HandleFunc("/health", handleHealth).Methods("GET")
	r.HandleFunc("/{short}", handleRedirect).Methods("GET")

	fmt.Println("🌐 Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

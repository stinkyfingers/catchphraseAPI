// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/stinkyfingers/catchphraseAPI/handlers"
// )

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		log.Fatal("$PORT must be set")
// 	}

// 	s, err := handlers.NewServer()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", s.HandleStatus)
// 	mux.HandleFunc("/all", s.Cors(s.HandleAll))
// 	mux.HandleFunc("/upload", s.Cors(s.HandleUpload))
// 	log.Print("Running on port ", port)
// 	log.Fatal(http.ListenAndServe(":"+port, mux))
// }

package main

import (
	"log"
	"net/http"
	"os"
	// "github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	// router.Run(":" + port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hey"))
	})
	http.ListenAndServe(":"+port, mux)
}

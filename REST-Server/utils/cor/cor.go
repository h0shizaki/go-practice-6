package cor

import (
	"log"
	"net/http"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.RequestURI)

		// if strings.Contains(r.RemoteAddr, "192.168.1.5") {
		// 	log.Println("Ping")
		// 	w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
		// } else {
		// 	w.Header().Set("Access-Control-Allow-Origin", "")
		// }

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})
}

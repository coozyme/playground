package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func (api *API) AllowOrigin(w http.ResponseWriter, req *http.Request) {
	// localhost:9000 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (api *API) AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// return unauthorized ketika token kosong
				w.WriteHeader(http.StatusUnauthorized)
				//encoder.Encode(map[string]string{"message": "unauthorized"})
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			//encoder.Encode(map[string]string{"message": "bad request"})
			return
		}
		// Task: Ambil value dari cookie token

		// TODO: answer here
		tknStr := c.Value
		// Task: Deklarasi variable claim

		// TODO: answer here
		claims := &Claims{}
		// Task: 1. parse JWT token ke dalam claim
		//       2. return unauthorized ketika signature invalid
		//       3. return bad request ketika field token tidak ada
		//       4. return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		// Task: Validasi
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// return unauthorized ketika signature invalid
				w.WriteHeader(http.StatusUnauthorized)
				//encoder.Encode(map[string]string{"message": "unauthorized"})
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			//encoder.Encode(map[string]string{"message": "bad request"})
			return
		}
		// Task: Validasi token
		if !tkn.Valid {
			// return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(map[string]string{"message": "unauthorized"})
			return
		}
		ctx := context.WithValue(r.Context(), "props", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
		//w.Write([]byte("Welcome to Cashier App"))

		//return next.ServeHTTP(w, r) // TODO: replace this
	})
}

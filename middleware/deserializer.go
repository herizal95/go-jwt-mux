package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/utils"
)

func Deserializer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aksesToken := r.Header.Get("Authorization")

		// Cek apakah header dimulai dengan "Bearer "
		if strings.HasPrefix(aksesToken, "Bearer ") {
			// Hapus prefix "Bearer " dan spasi di awal
			aksesToken = strings.TrimPrefix(aksesToken, "Bearer ")
		} else {
			// Jika header tidak dimulai dengan "Bearer ", kembalikan respons error
			helper.ResponseJson(w, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		if aksesToken == "" {
			helper.ResponseJson(w, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		user, err := utils.ValidateToken(aksesToken)
		if err != nil {
			helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		// Tambahkan informasi user ke dalam context
		ctx := context.WithValue(r.Context(), "userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

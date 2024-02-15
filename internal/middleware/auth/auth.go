package auth

import (
	"net/http"
	"strings"

	"github.com/dimitrijed93/load-balancer/config"
	"github.com/dimitrijed93/load-balancer/util"
	"github.com/rs/zerolog/log"
)

type Auth struct {
	handler http.Handler
}

func NewAuth(handler http.Handler) *Auth {
	return &Auth{handler: handler}
}

func (t Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authConfig := config.NewAuthConfig(r.URL.Path)
	if authConfig.Enabled {
		var allow bool = true
		log.Info().Msgf("Auth >> config enabled for path %s checking token", r.URL.Path)
		authHeader := r.Header.Get(util.HEADER_NAME_AUTHORIZATION)
		if authHeader == util.EMPTY_STRING {
			log.Info().Msgf("Auth >> auth header empty for path %s ", r.URL.Path)
			allow = false
		} else {
			splitToken := strings.Split(authHeader, util.PREFIX_BEARER_TOKEN)
			token := splitToken[1]
			if token == util.EMPTY_STRING {
				log.Info().Msgf("Auth >> auth token empty for path %s ", r.URL.Path)
				allow = false
			}
		}

		if !allow {
			log.Info().Msg("Unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Auth >> Unauthorized"))
			return
		}
	}
	t.handler.ServeHTTP(w, r)
}

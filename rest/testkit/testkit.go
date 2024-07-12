package testkit

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
	"github.com/ukasyah-dev/common/rest/server"
)

func New(s *server.Server) *apitest.APITest {
	return apitest.New().HandlerFunc(FiberToHandlerFunc(s.FiberApp))
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r, 10000)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

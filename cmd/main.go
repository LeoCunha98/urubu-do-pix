package main

import (
	"net/http"

	"github.com/LeoCunha98/urubu-do-pix/internal/domain/entity"
)

func main() {
	http.ListenAndServe(":8080", nil)
}
func Start(w http.ResponseWriter, r *http.Request) {
	user := entity.NewUser("Leo")
	println(user.Name)
}

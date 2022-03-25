package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/facutk/golaburo/utils"
)

func GetRedisHandler(w http.ResponseWriter, r *http.Request) {
	nn, _ := utils.Redis.Get("nn").Result()
	fmt.Fprint(w, nn)
}

func GetRedisUpdateHandler(w http.ResponseWriter, r *http.Request) {
	utils.Redis.Set("nn", strconv.FormatFloat(rand.Float64(), 'E', -1, 32), 0)
	nn, _ := utils.Redis.Get("nn").Result()
	fmt.Fprint(w, nn)
}

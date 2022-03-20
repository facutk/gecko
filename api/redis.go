package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/facutk/golaburo/utils"
)

func GetRedisHandler(w http.ResponseWriter, r *http.Request) {
	name := utils.Redis.Get("nn")
	fmt.Fprint(w, name)
}

func GetRedisUpdateHandler(w http.ResponseWriter, r *http.Request) {
	utils.Redis.Set("nn", strconv.FormatFloat(rand.Float64(), 'E', -1, 32), 0)
	name := utils.Redis.Get("nn")
	fmt.Fprint(w, name)
}

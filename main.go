package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// --- Une route GET “/” affichant l’heure qu’il est
func addZeroIfNecessary(number int) string {
	if number <= 9 {
		return "0" + fmt.Sprintf("%d", number)
	}
	return fmt.Sprintf("%d", number)
}

func actualtimeHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		day := time.Now()
		fmt.Fprintf(w, "%sh%s", addZeroIfNecessary(day.Hour()), addZeroIfNecessary(day.Minute()))
		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// --- Une route GET “/dice” affichant le résultat d’un dé à 1000 faces (D1000)
func diceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())
		fmt.Fprintf(w, "%d", rand.Intn(1000)+1)
		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

/* --- Une route GET “/dices” affichant quinze dés aux nombres de faces  */
func dicesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())

		if req.URL.Query().Get("type") != "" {
			if req.URL.Query().Get("type")[0] == 'd' {

				faces, _ := strconv.Atoi(req.URL.Query().Get("type")[1:])

				for i := 0; i < 15; i++ {
					fmt.Fprintf(w, "%d ", rand.Intn(faces)+1)
				}
				return

			} else {
				fmt.Fprintf(w, "Bad request")
				return
			}
		} else {

			dices := []int{2, 4, 6, 8, 10, 12, 20, 100}

			for i := 0; i < 15; i++ {
				faces := dices[rand.Intn(len(dices))]
				fmt.Fprintf(w, "%d ", rand.Intn(faces)+1)
			}
			return
		}

	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// --- Une route POST “/randomize-words”
func randomizeWordsHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		rand.Seed(time.Now().UnixNano())

		sentence := req.FormValue("words")

		words := strings.Split(sentence, " ")

		rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

		sentence = strings.Join(words, " ")

		fmt.Fprintf(w, "%s", sentence)

		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

// --- Une route POST “/semi-capitalize-sentence”
func semiCapitalizeSentenceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		sentence := req.FormValue("sentence")

		for i := 0; i < len(sentence); i++ {
			fmt.Println(sentence[i])
			if i%2 == 0 {
				sentence = sentence[:i] + strings.ToUpper(string(sentence[i])) + sentence[i+1:]
			}

		}

		fmt.Fprintf(w, "%s", sentence)

		return
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}

}

func main() {
	http.HandleFunc("/", actualtimeHandler)
	http.HandleFunc("/dice", diceHandler)
	http.HandleFunc("/dices?type=", dicesHandler)
	http.HandleFunc("/dices", dicesHandler)
	http.HandleFunc("/randomize-words", randomizeWordsHandler)
	http.HandleFunc("/semi-capitalize-sentence", semiCapitalizeSentenceHandler)
	http.ListenAndServe(":4567", nil)
}

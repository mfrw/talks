package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	p = flag.Bool("pprof", false, "Enable Profiling")
	t = flag.Bool("trace", false, "Enable Tracing")
)

/*
Then use the pprof tool to look at the heap profile:

go tool pprof http://localhost:8080/debug/pprof/heap

Or to look at a 30-second CPU profile:

go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30

Or to look at the goroutine blocking profile, after calling
runtime.SetBlockProfileRate in your program:

go tool pprof http://localhost:8080/debug/pprof/block

Or to collect a 5-second execution trace:

wget http://localhost:8080/debug/pprof/trace?seconds=5

Or to look at the holders of contended mutexes, after calling
runtime.SetMutexProfileFraction in your program:

go tool pprof http://localhost:8080/debug/pprof/mutex

To view all available profiles, open http://localhost:8080/debug/pprof/ in your
browser.

For a study of the facility in action, visit

https://blog.golang.org/2011/06/profiling-go-programs.html
*/

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}

func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		defer trackTime(time.Now(), "TIME")
		log.Printf("Logged connection from %s to %s\n", r.RemoteAddr, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if *p {
		log.Println("Profiling Enabled")
		pf, err := os.Create("pprof.out")
		if err != nil {
			log.Fatal("Could not create pprof file")
		}
		defer pf.Close()
		pprof.StartCPUProfile(pf)
		defer pprof.StopCPUProfile()
	}

	if *t {
		log.Println("Tracing Enabled")
		tf, err := os.Create("trace.out")
		if err != nil {
			log.Fatal("Could not create trace file")
		}
		defer tf.Close()
		trace.Start(tf)
		defer trace.Stop()
	}
	// START OMIT
	router := mux.NewRouter()
	router.Use(simpleMw)
	router.Use(handlers.CompressHandler)

	router.HandleFunc("/", IndexPage).Methods("GET")
	router.HandleFunc("/chart", chart).Methods("GET")
	router.HandleFunc("/data", data).Methods("GET")
	router.HandleFunc("/graph", graph).Methods("GET")
	router.HandleFunc("/style", style).Methods("GET")

	fmt.Println("Starting Server on : [localhost:8080]")
	log.Fatal(http.ListenAndServe(":8080", router))
	// END OMIT
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "final.html")
}

func chart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, "js/chart.js")
}
func style(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "css/style.css")
}

func data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{")
	fmt.Fprintf(w, "\"delayedPrice\":%f, \"delayedPriceTime\":%d", rand.Float32()*50.0, time.Now().Second())
	fmt.Fprintf(w, "}")
}

func graph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, "js/plotter.js")
}

func data_web(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://api.iextrading.com/1.0/stock/aapl/delayed-quote")
	if err != nil {
		log.Println("Could not get latest data", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

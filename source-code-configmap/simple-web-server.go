package main

import (
	"fmt"
	"net/http"
	"os"
)

type configurationListHandler struct {
	version           string
	environment       string
	environment_type  string
	region            string
	paypal_url        string
	db_user           string
	db_password       string
	gpu_enabled       string
	ui_theme          string
	cache_size        string
	page_limit        string
	sorting           string
	number_of_buckets string
}

func (h *configurationListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>I am a GO application running inside Kubernetes.<h2> <h3>My properties are:</h3><ul>")
	fmt.Fprintf(w, "<li>version: "+h.version+"</li>")
	fmt.Fprintf(w, "<li>environment: "+h.environment+"</li>")
	fmt.Fprintf(w, "<li>environment_type: "+h.environment_type+"</li>")
	fmt.Fprintf(w, "<li>region: "+h.region+"</li>")
	fmt.Fprintf(w, "<li>paypal_url: "+h.paypal_url+"</li>")
	fmt.Fprintf(w, "<li>db_user: "+h.db_user+"</li>")
	fmt.Fprintf(w, "<li>db_password: "+h.db_password+"</li>")
	fmt.Fprintf(w, "<li>gpu_enabled: "+h.gpu_enabled+"</li>")
	fmt.Fprintf(w, "<li>ui_theme: "+h.ui_theme+"</li>")
	fmt.Fprintf(w, "<li>cache_size: "+h.cache_size+"</li>")
	fmt.Fprintf(w, "<li>page_limit: "+h.page_limit+"</li>")
	fmt.Fprintf(w, "<li>sorting: "+h.sorting+"</li>")
	fmt.Fprintf(w, "<li>number_of_buckets: "+h.number_of_buckets+"</li>")
	fmt.Fprintf(w, "</ol>")

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")

}

func main() {

	clh := configurationListHandler{}
	clh.version = "3.0"
	clh.environment = os.Getenv("ENV")
	clh.environment_type = os.Getenv("ENV_TYPE")
	clh.region = os.Getenv("REGION")
	clh.paypal_url = os.Getenv("PAYPAL_URL")
	clh.db_user = os.Getenv("DB_USER")
	clh.db_password = os.Getenv("DB_PASSWORD")
	clh.gpu_enabled = os.Getenv("GPU_ENABLED")
	clh.ui_theme = os.Getenv("UI_THEME")
	clh.cache_size = os.Getenv("CACHE_SIZE")
	clh.page_limit = os.Getenv("PAGE_LIMIT")
	clh.sorting = os.Getenv("SORTING")
	clh.number_of_buckets = os.Getenv("N_BUCKETS")

	fmt.Println("Simple web server is starting on port 8080...")
	http.Handle("/", &clh)
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server at port 8080: %v", err)
	}
}

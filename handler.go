package urlshort

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	mux := http.NewServeMux()

	for k, v := range pathsToUrls {
		mux.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, v, http.StatusFound)
		})
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fallback.ServeHTTP(w, r)
	})

	return mux.ServeHTTP
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data := make([]map[string]string, 0)

	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		log.Fatal(err)
		return fallback.ServeHTTP, err
	}

	pathsToUrl := make(map[string]string)

	for _, v := range data {
		pathsToUrl[v["path"]] = v["url"]
	}

	return MapHandler(pathsToUrl, fallback), nil
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data := make([]map[string]string, 0)

	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Fatal(err)
		return fallback.ServeHTTP, err
	}

	pathsToUrl := make(map[string]string)

	for _, v := range data {
		pathsToUrl[v["path"]] = v["url"]
	}

	return MapHandler(pathsToUrl, fallback), nil
}

func DBHandler(fallback http.Handler) (http.HandlerFunc, error) {
	var db *sql.DB

	cfg := mysql.NewConfig()
	cfg.User = "user"
	cfg.Passwd = "cet.123$"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "urlshort"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	rows, err := db.Query("SELECT path, url FROM redirect")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pathsToUrl := make(map[string]string)

	for rows.Next() {

		var path string
		var url string

		err := rows.Scan(&path, &url)
		if err != nil {
			log.Fatal(err)
		}

		pathsToUrl[path] = url
		pathsToUrl[url] = url
	}

	return MapHandler(pathsToUrl, fallback), nil
}

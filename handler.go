package urlshort

import (
	"bytes"
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

	err := yaml.NewDecoder(bytes.NewReader(yml)).Decode(&data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = yaml.Unmarshal(yml, &data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pathsToUrl := make(map[string]string)

	for _, v := range data {
		pathsToUrl[v["path"]] = v["url"]
	}

	return MapHandler(pathsToUrl, fallback), nil
}

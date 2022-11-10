package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	wikiapi "github.com/mehoggan/simple-wiki-web-app-go/endpoints"
	wikitypes "github.com/mehoggan/simple-wiki-web-app-go/types"
	wikiutil "github.com/mehoggan/simple-wiki-web-app-go/util"
)

func generateWikiPage(rootDir string, page *wikitypes.Page) error {
	return wikiutil.Save(page, rootDir)
}

func generateConfigFile(rootDir string) (string, error) {
	configString := string("server:\n")
	configString += "  doc_root: \""
	configString += rootDir
	configString += "\""
	settingsFile := path.Join(rootDir, "settings.yaml")
	log.Printf("Saving \n%v \n\nto %v...", string(configString), settingsFile)
	err := os.WriteFile(settingsFile, []byte(configString), 0644)
	return settingsFile, err
}

func setupForWikiService() (string, string) {
	rootDir, err := createRootDir()
	if err != nil {
		log.Fatalf("Failed to create root directory for http server with %s.", err)
	}

	settingsFile, err := generateConfigFile(rootDir)
	if err != nil {
		log.Fatalf("Failed to create settings file with error %s.", err)
	}

	page := &wikitypes.Page{
		Title: "ABC",
		Body:  []byte("This is a test page.")}
	err = generateWikiPage(rootDir, page)
	if err != nil {
		log.Fatalf("Failed to create temporary wiki page.")
	}

	return rootDir, settingsFile
}

func createRootDir() (string, error) {
	return ioutil.TempDir(os.TempDir(), "go-web-services")
}

func main() {
	rootDir, settingsFile := setupForWikiService()
	wikiwebRoutes := wikiapi.InitializeEndpoints(settingsFile)
	http.HandleFunc("/view/",
		wikiwebRoutes.MakeHandler(wikiwebRoutes.ViewHandler))
	http.HandleFunc("/edit/",
		wikiwebRoutes.MakeHandler(wikiwebRoutes.EditHandler))
	http.HandleFunc("/save/",
		wikiwebRoutes.MakeHandler(wikiwebRoutes.SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer os.RemoveAll(rootDir)
}

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	wd "golang.org/x/net/webdav"
)

var rootDir string = os.Getenv("ENV_ROOT_DIR")

var wdDir wd.Dir = wd.Dir(rootDir)

var wdFs wd.FileSystem = wdDir

var wdHandle *wd.Handler = &wd.Handler{
	Prefix:     "",
	FileSystem: wdFs,
	LockSystem: wd.NewMemLS(),
	Logger: func(_ *http.Request, e error) {
		if nil != e {
			log.Printf("%v\n", e)
		}
	},
}

var htHandle http.Handler = wdHandle

var addr string = os.Getenv("ENV_ADDR_PORT")

func main() {
	http.Handle("/", htHandle)

	var s http.Server
	s.Addr = addr
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	e := s.ListenAndServe()
	if nil != e {
		log.Printf("%v\n", e)
	}
}

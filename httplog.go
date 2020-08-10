package plugin_httplog

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"

	"github.com/kjk/betterguid"
)

// Config holds the plugin configuration.
type Config struct {
	Request      bool `json:"request,omitempty"`
	RequestBody  bool `json:"requestbody,omitempty"`
	Response     bool `json:"response,omitempty"`
	ResponseBody bool `json:"responsebody,omitempty"`
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// HTTPLog is a httpLog plugin.
type HTTPLog struct {
	Name       string
	Next       http.Handler
	Config     *Config
	LogHandler *log.Logger
}

// New creates and returns a plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HTTPLog{
		Name:   name,
		Next:   next,
		Config: config,
	}, nil
}

func (h *HTTPLog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	id := betterguid.New()
	h.LogHandler = log.New(os.Stdout, "[HTTPLOG"+id+"] ", 0)
	h.logger().ServeHTTP(rw, req)
}

func (h *HTTPLog) printDump(dump []byte, msg string) {
	h.LogHandler.Printf("********* %s *********", msg)
	reader := bytes.NewReader(dump)
	bufReader := bufio.NewReader(reader)
	for {
		line, _, err := bufReader.ReadLine()
		h.LogHandler.Println(string(line))
		if err != nil {
			break
		}
	}
}

func (h *HTTPLog) logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, h.Config.RequestBody)
		if err != nil {
			h.LogHandler.Println(err)
		}
		h.printDump(requestDump, "REQUEST")

		rec := httptest.NewRecorder()
		h.Next.ServeHTTP(rec, r)

		responseDump, err := httputil.DumpResponse(rec.Result(), h.Config.ResponseBody)
		if err != nil {
			h.LogHandler.Println(err)
		}
		h.printDump(responseDump, "RESPONSE")

		for k, vv := range rec.Header() {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}

		data := rec.Body.Bytes()

		w.WriteHeader(rec.Code)
		w.Write(data)
	})
}

package jvs

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"io"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"sync"
)

const (
	listingEndpoint = "https://jbovlaste.lojban.org/export/xml.html"
	exportEndpoint  = "https://jbovlaste.lojban.org/export/xml-export.html"
)

var (
	langRegexp = regexp.MustCompile(`<a href="xml-export\.html\?lang=([a-z]+)\">.*?</a>`)
)

func FetchAll(cookieLine string, out IndexType, ok chan<- struct{}) error {
	resp, err := http.Get(listingEndpoint)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	matches := langRegexp.FindAllSubmatch(all, -1)
	keys := make([]string, 0, len(matches))
	for _, v := range matches {
		keys = append(keys, string(v[1]))
	}

  close(ok)

	var (
		wg  = sync.WaitGroup{}
		cl  = &http.Client{}
		sem = semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
		ctx = context.Background()

    mut = sync.Mutex{}
    tot = 0
	)

	for _, k := range keys {
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}
		wg.Add(1)
		k := k
		go func() {
			defer func() {
				wg.Done()
				sem.Release(1)
			}()
			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?lang=%s", exportEndpoint, k), nil)
			if err != nil {
				log.Print(err)
				return
			}
			req.Header.Add("Cookie", cookieLine)
			res, err := cl.Do(req)
			if err != nil {
				log.Print(err)
				return
			}
			all, err := io.ReadAll(res.Body)
			if err != nil {
				log.Print(err)
				return
			}
			if bytes.Contains(all, []byte(`<script src="https://www.google.com/recaptcha/api.js" async defer>`)) {
				log.Println(fmt.Errorf("got captcha page"))
				return
			}
			dic, err := Parse(bytes.NewReader(all))
			if err != nil {
				log.Print(err)
				return
			}

			out.Mutex.Lock()
			out.Index[k] = dic
      l := len(out.Index[k])
      if l > 0 {
        log.Printf("success on dic %s – %d entries loaded", k, l)
      }
      out.Mutex.Unlock()

      mut.Lock()
      defer mut.Unlock()
      tot += l
		}()
	}

  go func(){
    wg.Wait()
    log.Printf("loaded %d dictionaries with %d entries total", len(out.Index), tot)
  }()

	return nil
}

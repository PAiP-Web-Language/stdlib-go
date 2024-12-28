package test

import (
	"testing"
	"github.com/levigross/grequests"
	"github.com/PAiP-Web-Language/stdlib-go/client/http"
)

type HttpClientTestResponse struct {
	*grequests.Response
}

func HttpClientGet(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Get(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func HttpClientHead(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Head(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func HttpClientPost(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Post(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func HttpClientPut(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Put(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func HttpClientDelete(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Delete(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func HttpClientPatch(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Patch(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func HttpClientOptions(t testing.TB, url string, ro *grequests.RequestOptions) HttpClientTestResponse {
	t.Helper()
	r := http.Options(url, ro)
	Assert(t, r.IsOk())
	return HttpClientTestResponse{r.Unwrap()}
}

func (r HttpClientTestResponse) ExpectStatus(t testing.TB, expected int) HttpClientTestResponse {
	t.Helper()
	Assert(t, r.StatusCode == expected)
	return r
}

func (r HttpClientTestResponse) ExpectContentType(t testing.TB, expected string) HttpClientTestResponse {
	t.Helper()
	Assert(t, r.Header.Get("Content-Type") == expected)
	return r
}

func (r HttpClientTestResponse) ExpectHeader(t testing.TB, key string, expected string) HttpClientTestResponse {
	t.Helper()
	Assert(t, r.Header.Get(key) == expected)
	return r
}

func (r HttpClientTestResponse) ExpectBody(t testing.TB, expected string) HttpClientTestResponse {
	t.Helper()
	Assert(t, r.String() == expected)
	return r
}

package http

import (
	"github.com/levigross/grequests"
	. "github.com/PAiP-Web-Language/stdlib-go/base"
)

func Get(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Get(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func Post(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Post(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func Delete(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Delete(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func Head(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Head(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func Put(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Put(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func Patch(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Patch(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func Options(url string, ro *grequests.RequestOptions) Result[*grequests.Response] {
	r, err := grequests.Options(url, ro)
	if err != nil {
		return Err[*grequests.Response](err)
	}
	_ = r.String() // Cache response
	return Ok(r)
}

func ClearInternalBuffer(r *grequests.Response) {
	r.ClearInternalBuffer()
}

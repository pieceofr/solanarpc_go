package solanarpc

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
)

func RandomID() uint64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(MAXID))
	if err != nil {
		return 1
	}
	return nBig.Uint64()
}

// A better way to close http response body for reusing connection efficiency
// https://github.com/google/go-github/pull/317
func CloseRespBody(resp *http.Response) {
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
}

package browser

import "testing"

func TestBrowse(t *testing.T) {
	browse := History{}
	browse.Visit("https://1")
	browse.Visit("https://2")
	browse.Visit("https://3")

	t.Log(browse.Back())
	t.Log(browse.Back())
	t.Log(browse.Back())
	t.Log(browse.Back())
	t.Log(browse.Forward())
	t.Log(browse.Forward())
	t.Log(browse.Forward())
	t.Log(browse.Forward())
	t.Log(browse.Forward())
}

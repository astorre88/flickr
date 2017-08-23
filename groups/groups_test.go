package groups

import (
  "testing"

  "github.com/astorre88/flickr"
  flickErr "gopkg.in/masci/flickr.v2/error"
)

func TestAddPhoto(t *testing.T) {
  fclient := flickr.GetTestClient()
  server, client := flickr.FlickrMock(200, `<rsp stat="ok"></rsp>`, "text/xml")
  defer server.Close()
  fclient.HTTPClient = client

  _, err := AddPhoto(fclient, "123456", "123")
  flickr.Expect(t, err, nil)

  server, client = flickr.FlickrMock(200, `<rsp stat="fail"></rsp>`, "text/xml")
  defer server.Close()
  fclient.HTTPClient = client

  resp, err := AddPhoto(fclient, "123456", "123")
  _, ok := err.(*flickErr.Error)
  flickr.Expect(t, ok, true)
  flickr.Expect(t, resp.HasErrors(), true)

  // check params, reset Flickr client to dismiss mocked responses
  fclient = flickr.GetTestClient()
  AddPhoto(fclient, "123456", "123")
  params := []string{"group_id", "photo_id"}
  flickr.AssertParamsInBody(t, fclient, params)
}

func TestSearch(t *testing.T) {
  fclient := flickr.GetTestClient()
  server, client := flickr.FlickrMock(200, `<rsp stat="ok"></rsp>`, "text/xml")
  defer server.Close()
  fclient.HTTPClient = client

  _, err := Search(fclient, "query")
  flickr.Expect(t, err, nil)

  server, client = flickr.FlickrMock(200, `<rsp stat="fail"></rsp>`, "text/xml")
  defer server.Close()
  fclient.HTTPClient = client

  resp, err := Search(fclient, "query")
  _, ok := err.(*flickErr.Error)
  flickr.Expect(t, ok, true)
  flickr.Expect(t, resp.HasErrors(), true)

  // check params, reset Flickr client to dismiss mocked responses
  fclient = flickr.GetTestClient()
  Search(fclient, "query")
  params := []string{"text"}
  flickr.AssertParamsInBody(t, fclient, params)
}

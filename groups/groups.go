// Package implementing methods: flickr.groups.*
package groups

import (
  "github.com/astorre88/flickr"
)

// Add a photo to a group
// This method requires authentication with 'write' permission.
func AddPhoto(client *flickr.FlickrClient, groupId string, photoId string) (*flickr.BasicResponse, error) {
  client.Init()
  client.HTTPVerb = "POST"
  client.Args.Set("method", "flickr.groups.pools.add")
  client.Args.Set("group_id", groupId)
  client.Args.Set("photo_id", photoId)

  client.OAuthSign()

  response := &flickr.BasicResponse{}
  err := flickr.DoPost(client, response)
  return response, err
}

// Search a group by a query
func Search(client *flickr.FlickrClient, text string) (*flickr.BasicResponse, error) {
  client.Init()
  client.HTTPVerb = "GET"
  client.Args.Set("method", "flickr.groups.search")
  client.Args.Set("text", text)

  client.OAuthSign()

  response := &flickr.BasicResponse{}
  err := flickr.DoGet(client, response)
  return response, err
}

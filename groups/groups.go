package groups

import (
  "gopkg.in/masci/flickr.v2"
)

// Add a photo to a group
// This method requires authentication with 'write' permission.
func AddPhoto(client *flickr.FlickrClient, groupId, photoId string) (*flickr.BasicResponse, error) {
  client.Init()
  client.HTTPVerb = "POST"
  client.Args.Set("method", "flickr.groups.pools.add")
  client.Args.Set("group_id", groupId)
  client.Args.Set("photo_id", photoId)

  client.OAuthSign()

  response := &flickr.Basic.BasicResponse{}
  err := flickr.DoPost(client, response)
  return response, err
}

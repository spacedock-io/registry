package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestTagCreate(t *testing.T) {
  err := SpacedockFooLatestTag.Create()
  assert.Nil(t, err, "Error should be `nil`")
}

func TestGetTags(t *testing.T) {
  tags, err := GetTags(SpacedockFooLatestTag.Namespace, SpacedockFooLatestTag.Repo)
  assert.Nil(t, err, "Error should be `nil`")
  assert.Equal(t, len(tags), 1, "One tag should exist")
  assert.Equal(t, tags[0].Tag, SpacedockFooLatestTag.Tag, "Tag should be correct")
  assert.Equal(t, tags[0].Repo, SpacedockFooLatestTag.Repo, "Repo should be correct")
  assert.Equal(t, tags[0].Namespace, SpacedockFooLatestTag.Namespace, "Namespace should be correct")
}

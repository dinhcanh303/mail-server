package meili

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	Host   = "http://localhost:7700"
	APIKey = "ms"
)

func TestMeiliSearchAddDocuments(t *testing.T) {
	meiliSearch := connectMeiliSearch()
	documents := []map[string]interface{}{
		{"uuid": "abdf-123394j", "name": "Foden Ngo", "genres": []string{"Drama"}},
		{"uuid": "abdf-123395j", "name": "Fortune Truong", "genres": []string{"Drama"}},
		{"uuid": "abdf-123396j", "title": "Kaka", "genres": []string{"Drama"}},
		{"uuid": "abdf-123398j", "title": "Messi", "genres": []string{"Drama"}},
	}
	task, err := meiliSearch.AddDocuments("abc", documents)
	require.NoError(t, err)
	require.NotEmpty(t, task)
}
func TestCreateIndex(t *testing.T) {
	meiliSearch := connectMeiliSearch()
	task, err := meiliSearch.CreateIndex("test1", "uuid")
	require.NoError(t, err)
	require.NotEmpty(t, task)
}

func TestDeleteIndex(t *testing.T) {
	meiliSearch := connectMeiliSearch()
	task, err := meiliSearch.DeleteIndex("groups")
	require.NoError(t, err)
	require.NotEmpty(t, task)
}
func TestMeiliSearch(t *testing.T) {
	meiliSearch := connectMeiliSearch()
	task, err := meiliSearch.Search("test", "Car", nil)
	require.NoError(t, err)
	require.NotEmpty(t, task.Hits)
}
func TestMeiliSearchDeleteDocument(t *testing.T) {
	meiliSearch := connectMeiliSearch()
	task, err := meiliSearch.DeleteDocument("test", "7")
	require.NoError(t, err)
	require.NotEmpty(t, task)
}
func TestMeiliSearchDeleteAllDocuments(t *testing.T) {
	meiliSearch := connectMeiliSearch()
	task, err := meiliSearch.DeleteAllDocuments("user-group")
	require.NoError(t, err)
	require.NotEmpty(t, task)
}

func connectMeiliSearch() MeiliSearch {
	conn := MeiliSearchConn{
		Host:   Host,
		APIKey: APIKey,
	}
	meiliSearch := NewMeiliSearch(conn)
	return meiliSearch
}

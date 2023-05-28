package helpers

import (
	"encoding/json"
	"github.com/aldarisbm/memory/pkg/types"
	"os"
	"time"
)

// RetrieveDocumentsFromJsonFile retrieves documents from a json file
// that is in the format of the OpenAIConversation struct (see struct for more info)
// and returns them as a slice of Documents
func RetrieveDocumentsFromJsonFile(path string) []*types.Document {
	var conversations []*OpenAIConversation
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &conversations)
	if err != nil {
		panic(err)
	}

	var documents []*types.Document
	for _, item := range conversations {
		for _, v := range item.Mapping {
			if v.Message == nil {
				continue
			}
			if v.Message.Content.Parts[0] == "" {
				continue
			}
			documents = append(documents, &types.Document{
				ID:         v.Message.ID,
				User:       v.Message.Author.Role,
				Text:       v.Message.Content.Parts[0],
				CreatedAt:  time.Unix(int64(v.Message.CreateTime), 0),
				LastReadAt: time.Unix(int64(v.Message.CreateTime), 0),
				Metadata: map[string]any{
					"conversation_title": item.Title,
				},
			})
		}
	}
	return documents
}

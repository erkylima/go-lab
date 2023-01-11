package route

import (
	"encoding/json"
	"net/http"

	"github.com/erkylima/golab/internal/entity"
	"github.com/erkylima/golab/internal/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts"}`))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func AddPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}
	result, err := json.Marshal(post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the post"}`))
		return
	}
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

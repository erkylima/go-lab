package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/erkylima/golab/internal/entity"
	"google.golang.org/api/option"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId = "golab-af2f4"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("assets/golab-af2f4-firebase-adminsdk-87lbh-50001a2dee.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal("firebase.NewApp: ", err)
	}
	client, err := app.Firestore(ctx)
	defer client.Close()
	if err != nil {
		log.Fatal("app.Firestore: ", err)
	}
	client.Collection("posts").Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("assets/golab-af2f4-firebase-adminsdk-87lbh-50001a2dee.json")
	config := &firebase.Config{
		AuthOverride:     &map[string]interface{}{},
		DatabaseURL:      "",
		ProjectID:        projectId,
		ServiceAccountID: "",
		StorageBucket:    "",
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal("firebase.NewApp: ", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("app.Firestore: ", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iterator := client.Collection("posts").Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		log.Printf("doc.Data: %v", doc.Data())

		/*post := entity.Post{
			ID:    doc.Data()["ID"].(int),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}*/
		//posts = append(posts, post)
	}
	return posts, nil
}

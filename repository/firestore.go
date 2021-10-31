package repository

import (
	"context"
	"go-rest-api/entity"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	ProjectId      string = "go-rest-api-16128"
	CollectionName string = "posts"
)

type Repository struct{}

func NewFirestoreRepository() PostRepository {
	return &Repository{}
}

func (r *Repository) Save(post *entity.Post) (*entity.Post, error) {
	context := context.Background()
	var opt = option.WithCredentialsFile("/home/harv3n/firebase/firestore-key.json")
	config := &firebase.Config{ProjectID: ProjectId}
	app, err := firebase.NewApp(context, config, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase App: %v", err)
	}

	client, err := app.Firestore(context)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(CollectionName).Add(context, map[string]interface{}{
		"id":    post.Id,
		"title": post.Title,
		"text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (r *Repository) FindAll() ([]entity.Post, error) {
	context := context.Background()

	var opt = option.WithCredentialsFile("/home/harv3n/firebase/firestore-key.json")
	config := &firebase.Config{ProjectID: ProjectId}
	app, err := firebase.NewApp(context, config, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase App: %v", err)
	}

	client, err := app.Firestore(context)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(CollectionName).Documents(context)
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf(" Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			Id:    doc.Data()["id"].(int64),
			Title: doc.Data()["title"].(string),
			Text:  doc.Data()["text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}

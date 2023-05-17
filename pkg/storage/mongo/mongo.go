package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store Хранилище данных.
type Store struct {
	client mongo.Client
	coll   mongo.Collection
}

// New Конструктор объекта хранилища.
func New(mongoUri string) (*Store, error) {
	opts := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	s := Store{
		client: *client,
		coll:   *client.Database("News").Collection("posts"),
	}

	return &s, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	var posts []storage.Post

	cur, err := s.coll.Find(context.Background(), bson.M{})
	defer cur.Close(context.Background())

	if err != nil {
		return nil, err
	}

	err = cur.All(context.Background(), &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Store) AddPost(p storage.Post) error {

	_, err := s.coll.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}
func (s *Store) UpdatePost(p storage.Post) error {

	_, err := s.coll.UpdateOne(context.Background(), p.ID, p)
	if err != nil {
		return err
	}

	return nil
}
func (s *Store) DeletePost(p storage.Post) error {

	_, err := s.coll.DeleteOne(context.Background(), p.ID)
	if err != nil {
		return err
	}

	return nil
}

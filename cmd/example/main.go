package main

import (
	"context"
	"encoding/json"

	"github.com/mongorpc/mongorpc/lib/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	c := client.NewClient("localhost:8080")
	conn, err := c.Connect(
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()

	// Initilize database
	db := c.Database("sample_mflix")

	// Get Document By ID
	// doc, err := db.Collection("movies").Document("573a1390f29313caabcd4135").Get(context.TODO())
	// if err != nil {
	// 	logrus.Errorln(err)
	// }
	// logrus.Infoln(doc)

	// err = db.Collection("movies").CreateIndex(context.TODO(), client.Index{
	// 	Name: "title_idx",
	// 	Keys: []*client.IndexKey{
	// 		{
	// 			Field:     "title",
	// 			Direction: client.IndexDirection_ASCENDING,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	logrus.Errorln(err)
	// }

	// docs, err := db.Collection("movies").Documents().Limit(10).Skip(10).Sort("name", client.ASCENDING).Get(context.TODO())
	// // docs, err := db.Collection("movies").Documents().EqualTo("title", "Gertie the Dinosaur").EqualTo("year", 1914).Get(context.TODO())
	// // docs, err := db.Collection("movies").Documents().Search("Batman").Get(context.TODO())
	// if err != nil {
	// 	logrus.Errorln(err)
	// }
	// logrus.Infoln(docs)

	// listen for changes
	listner, err := db.Collection("movies").Documents().EqualTo("title", "Gertie the Dinosaur").Listen(context.TODO(), nil)
	if err != nil {
		logrus.Errorln(err)
	}
	defer close(listner)

	// go func() {
	for doc := range listner {
		jsonString, err := json.Marshal(doc)
		if err != nil {
			logrus.Errorln(err)
		}
		logrus.Infoln(string(jsonString))
	}
	// }()
}

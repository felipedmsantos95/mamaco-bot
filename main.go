package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env File not found")
	}

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("TOKEN_KEY"), os.Getenv("TOKEN_SECRET_KEY"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Client do twitter
	client := twitter.NewClient(httpClient)

	// Publicar um tweet
	tweet, _, err := client.Statuses.Update("Beto Otangorando tá on, pae!", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(tweet.Text)

}

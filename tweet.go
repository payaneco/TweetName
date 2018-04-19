package main

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
)

//下記をお借りして実装しました
//http://qiita.com/hnaohiro/items/df6bcc4fcff3633b4c62
type Twitter struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

type OAuth struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Twitter {
	twitter := new(Twitter)
	twitter.consumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	twitter.accessToken = &oauth.AccessToken{accessToken, accessTokenSecret, nil}
	return twitter
}

func (t *Twitter) Post(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Post(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}

func Tweet(text string) {
	oa := GetOAuth()
	twitter := NewTwitter(oa.ConsumerKey, oa.ConsumerSecret, oa.AccessToken, oa.AccessTokenSecret)

	res, err := twitter.Post(
		"https://api.twitter.com/1.1/statuses/update.json", // Resource URL
		map[string]string{"status": text}) // Parameters
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func Rename(text string) {
	oa := GetOAuth()
	twitter := NewTwitter(oa.ConsumerKey, oa.ConsumerSecret, oa.AccessToken, oa.AccessTokenSecret)

	res, err := twitter.Post(
		"https://api.twitter.com/1.1/account/update_profile.json", // Resource URL
		map[string]string{"name": text}) // Parameters
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func GetOAuth() OAuth {
	bytes, err := ioutil.ReadFile("oauth.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var oauth OAuth
	if err := json.Unmarshal(bytes, &oauth); err != nil {
		log.Fatal(err)
	}
	return oauth

}

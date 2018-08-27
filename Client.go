package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Account struct {
	SellerId      string
	AccessKey     string
	SecretKey     string
	MarketplaceId string
	AuthToken     string
}

type Client struct {
	Account
	Host    string
}

type Request map[string]interface{}
type Parameters map[string]interface{}

func NewClient(account Account) *Client {
	host := GetMarketplaceEndpoint(account.MarketplaceId)
	return &Client{Account: account, Host: host}
}

func (c Client) genSignAndFetch(Action string, ActionPath string, Parameters map[string]string) (string, error) {
	genUrl, err := c.GenerateAmazonUrl(Action, ActionPath, Parameters)
	if err != nil {
		return "", err
	}

	SetTimestamp(genUrl)

	signedurl, err := c.SignAmazonUrl(genUrl)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(signedurl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c Client) GenerateAmazonUrl(Action string, ActionPath string, Parameters map[string]string) (finalUrl *url.URL, err error) {
	result, err := url.Parse(c.Host)
	if err != nil {
		return nil, err
	}

	result.Host = c.Host
	result.Scheme = "https"
	result.Path = ActionPath

	values := url.Values{}
	values.Add("Action", Action)

	if c.AuthToken != "" {
		values.Add("MWSAuthToken", c.AuthToken)
	}

	values.Add("AWSAccessKeyId", c.AccessKey)
	values.Add("SellerId", c.SellerId)
	values.Add("SignatureVersion", "2")
	values.Add("SignatureMethod", "HmacSHA256")
	values.Add("Version", "2011-10-01")

	for k, v := range Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	result.RawQuery = params

	return result, nil
}

func SetTimestamp(origUrl *url.URL) (err error) {
	values, err := url.ParseQuery(origUrl.RawQuery)
	if err != nil {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origUrl.RawQuery = values.Encode()

	return nil
}

func (c Client) SignAmazonUrl(origUrl *url.URL) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(c.SecretKey))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}

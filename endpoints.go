package larry

// LookupTweetsByIDs returns a variety of information about the Tweets specified by the requested ID or list of IDs.
func (c *Client) LookupTweetsByIDs(ids []string, opts ...RequestOption) (*Tweets, error) {
	endpointUrl := "https://api.twitter.com/2/tweets"
	opts = append(opts, WithIDs(ids...))
	req, err := NewRequest(endpointUrl, opts...)
	if err != nil {
		return nil, err
	}
	tweets := new(Tweets)
	err = c.MakeRequest(req, tweets)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

// LookupTweetByID returns a variety of information about a single Tweet specified by the requested ID.
func (c *Client) LookupTweetByID(id string, opts ...RequestOption) (*Tweet, error) {
	endpointUrl := "https://api.twitter.com/2/tweets/" + id
	req, err := NewRequest(endpointUrl, opts...)
	if err != nil {
		return nil, err
	}
	tweet := new(Tweet)
	err = c.MakeRequest(req, tweet)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

// SearchRecentTweets The recent search endpoint returns Tweets from the last 7 days that match a search query.
func (c *Client) SearchRecentTweets(query string, opts ...RequestOption) (*Tweets, error) {
	endpointUrl := "https://api.twitter.com/2/tweets/search/recent"
	opts = append(opts, WithQuery(query))
	req, err := NewRequest(endpointUrl, opts...)
	if err != nil {
		return nil, err
	}
	tweets := new(Tweets)
	err = c.MakeRequest(req, tweets)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

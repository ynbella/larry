package larry

type Tweets struct {
	Tweets   []Tweet         `json:"data"`
	Metadata *TweetsMetadata `json:"meta"`
}

type Domain struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Entity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Entities struct {
	Annotations []Annotation `json:"annotations"`
	Urls        []Url        `json:"urls"`
	Hashtags    []Hashtag    `json:"hashtags"`
	Mentions    []Mention    `json:"mentions"`
	Cashtags    []Cashtag    `json:"cashtags"`
}

type Annotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type Url struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
	UnwoundUrl  string `json:"unwound_url"`
}

type Hashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type Mention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
}

type Cashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type Withheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type Includes struct {
	Tweets []Tweet `json:"tweets"`
	Users  []User  `json:"users"`
	Places []Place `json:"places"`
	Media  []Media `json:"media"`
	Polls  []Poll  `json:"polls"`
}

type TweetsMetadata struct {
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

// region Tweet Object

type Tweet struct {
	Id                 string                   `json:"id"`
	Text               string                   `json:"text"`
	Attachments        *TweetAttachments        `json:"attachments"`
	AuthorId           string                   `json:"author_id"`
	ContextAnnotations []TweetContextAnnotation `json:"context_annotations"`
	ConversationId     string                   `json:"conversation_id"`
	CreatedAt          string                   `json:"created_at"`
	Entities           *Entities                `json:"entities"`
	Geo                *TweetGeo                `json:"geo"`
	InReplyToUserId    string                   `json:"in_reply_to_user_id"`
	Lang               string                   `json:"lang"`
	NonPublicMetrics   *TweetNonPublicMetrics   `json:"non_public_metrics"`
	OrganicMetrics     *TweetOrganicMetrics     `json:"organic_metrics"`
	PossiblySensitive  bool                     `json:"possibly_sensitive"`
	PromotedMetrics    *TweetPromotedMetrics    `json:"promoted_metrics"`
	PublicMetrics      *TweetPublicMetrics      `json:"public_metrics"`
	ReferencedTweets   []ReferencedTweet        `json:"referenced_tweets"`
	Source             string                   `json:"source"`
	Withheld           *Withheld                `json:"withheld"`

	Includes *Includes `json:"includes"`
	Errors   []string  `json:"errors"` // TODO Implement errors
}

type Coordinates struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

type ReferencedTweet struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type TweetAttachments struct {
	MediaKeys []string `json:"media_keys"`
	PollIDs   []string `json:"poll_ids"`
}

type TweetContextAnnotation struct {
	Domain Domain `json:"domain"`
	Entity Entity `json:"entity"`
}

type TweetGeo struct {
	Coordinates Coordinates `json:"coordinates"`
	PlaceID     string      `json:"place_id"`
}

type TweetNonPublicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	UrlLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

type TweetOrganicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	UrlLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	RetweetCount      int `json:"retweet_count"`
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
}

type TweetPromotedMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	UrlLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	RetweetCount      int `json:"retweet_count"`
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
}

type TweetPublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

// endregion

// region User Object

// User https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/user
type User struct {
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	UserName        string            `json:"username"`
	CreatedAt       string            `json:"created_at"`
	Description     string            `json:"description"`
	Entities        Entities          `json:"entities"`
	Location        string            `json:"location"`
	PinnedTweetId   string            `json:"pinned_tweet_id"`
	ProfileImageUrl string            `json:"profile_image_url"`
	Protected       bool              `json:"protected"`
	PublicMetrics   UserPublicMetrics `json:"public_metrics"`
	URL             string            `json:"url"`
	Verified        bool              `json:"verified"`
	WithHeld        Withheld          `json:"withheld"`
}

type UserPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetsCount    int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

// endregion

// region Media Object

// Media https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/media
type Media struct {
	MediaKey         string                `json:"media_key"`
	Type             string                `json:"type"`
	DurationMS       int                   `json:"duration_ms"`
	Height           int                   `json:"height"`
	NonPublicMetrics MediaNonPublicMetrics `json:"non_public_metrics"`
	OrganicMetrics   MediaOrganicMetrics   `json:"organic_metrics"`
	PreviewImageUrl  string                `json:"preview_image_url"`
	PromotedMetrics  MediaPromotedMetrics  `json:"promoted_metrics"`
	PublicMetrics    MediaPublicMetrics    `json:"public_metrics"`
	Width            int                   `json:"width"`
}

type MediaNonPublicMetrics struct {
	Playback0   int `json:"playback_0_count"`
	Playback100 int `json:"playback_100_count"`
	Playback25  int `json:"playback_25_count"`
	Playback50  int `json:"playback_50_count"`
	Playback75  int `json:"playback_75_count"`
}

type MediaOrganicMetrics struct {
	Playback0   int `json:"playback_0_count"`
	Playback100 int `json:"playback_100_count"`
	Playback25  int `json:"playback_25_count"`
	Playback50  int `json:"playback_50_count"`
	Playback75  int `json:"playback_75_count"`
	ViewCount   int `json:"view_count"`
}

type MediaPromotedMetrics struct {
	Playback0   int `json:"playback_0_count"`
	Playback100 int `json:"playback_100_count"`
	Playback25  int `json:"playback_25_count"`
	Playback50  int `json:"playback_50_count"`
	Playback75  int `json:"playback_75_count"`
	ViewCount   int `json:"view_count"`
}

type MediaPublicMetrics struct {
	ViewCount int `json:"view_count"`
}

// endregion

// region Poll Object

// Poll https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/poll
type Poll struct {
	Id              string        `json:"id"`
	Options         []PollOptions `json:"options"`
	DurationMinutes int           `json:"duration_minutes"`
	EndDateTime     string        `json:"end_datetime"`
	VotingStatus    string        `json:"voting_status"`
}

type PollOptions struct {
	Position int    `json:"position"`
	Label    string `json:"label"`
	Votes    int    `json:"votes"`
}

// endregion

// region Place Object

// Place https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/place
type Place struct {
	FullName        string   `json:"full_name"`
	Id              string   `json:"id"`
	ContainedWithin []string `json:"contained_within"`
	Country         string   `json:"country"`
	CountryCode     string   `json:"country_code"`
	Geo             string   `json:"geo"` // TODO Need GeoJSON library
	Name            string   `json:"name"`
	PlaceType       string   `json:"place_type"'`
}

type GeoJson struct {
	Type       string                 `json:"type"`
	BBox       []float64              `json:"bbox,omitempty"`
	Properties map[string]interface{} `json:"properties"`
}

// endregion

package larry

import (
	"net/url"
	"strconv"
	"time"
)

type Request url.URL

type RequestOption func(*Request)

func WithQueryParams(key string, values ...string) RequestOption {
	return func(req *Request) {
		params := new(url.Values)
		for _, value := range values {
			params.Add(key, value)
		}
		req.RawQuery = params.Encode()
	}
}

func NewRequest(endpoint string, opts ...RequestOption) (*Request, error) {
	req, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Loop through each option
	for _, opt := range opts {
		opt((*Request)(req))
	}

	return (*Request)(req), nil
}

type Expansion string

const (
	AttachmentsPollIds         Expansion = "attachments.poll_ids"
	AttachmentsMediaKeys       Expansion = "attachments.media_keys"
	AuthorId                   Expansion = "author_id"
	EntitiesMentionsUsername   Expansion = "entities.mentions.username"
	GeoPlaceId                 Expansion = "geo.place_id"
	InReplyToUserId            Expansion = "in_reply_to_user_id"
	ReferencedTweetsId         Expansion = "referenced_tweets.id"
	ReferencedTweetsIdAuthorId Expansion = "referenced_tweets.id.author_id"
)

func WithExpansions(expansions ...Expansion) RequestOption {
	values := make([]string, len(expansions))
	for i, v := range expansions {
		values[i] = string(v)
	}
	return WithQueryParams("expansions", values...)
}

func WithIDs(ids ...string) RequestOption {
	return WithQueryParams("ids", ids...)
}

type MediaField string

const (
	DurationMs       MediaField = "duration_ms"
	Height           MediaField = "height"
	MediaKey         MediaField = "media_key"
	PreviewImageUrl  MediaField = "preview_image_url"
	Type             MediaField = "type"
	MediaUrl         MediaField = "url"
	Width            MediaField = "width"
	PublicMetrics    MediaField = "public_metrics"
	NonPublicMetrics MediaField = "non_public_metrics"
	OrganicMetrics   MediaField = "organic_metrics"
	PromotedMetrics  MediaField = "promoted_metrics"
)

func WithMediaFields(fields ...MediaField) RequestOption {
	values := make([]string, len(fields))
	for i, v := range fields {
		values[i] = string(v)
	}
	return WithQueryParams("media.fields", values...)
}

type PlaceField string

const (
	ContainedWithin PlaceField = "contained_within"
	Country         PlaceField = "country"
	CountryCode     PlaceField = "country_code"
	FullName        PlaceField = "full_name"
	PlaceGeo        PlaceField = "geo"
	PlaceId         PlaceField = "id"
	PlaceName       PlaceField = "name"
	PlaceType       PlaceField = "place_type"
)

func WithPlaceFields(fields ...PlaceField) RequestOption {
	values := make([]string, len(fields))
	for i, v := range fields {
		values[i] = string(v)
	}
	return WithQueryParams("place.fields", values...)
}

type PollField string

const (
	DurationMinutes PollField = "duration_minutes"
	EndDatetime     PollField = "end_datetime"
	PollId          PollField = "id"
	Options         PollField = "options"
	VotingStatus    PollField = "voting_status"
)

func WithPollFields(fields ...PollField) RequestOption {
	values := make([]string, len(fields))
	for i, v := range fields {
		values[i] = string(v)
	}
	return WithQueryParams("poll.fields", values...)
}

type TweetField string

const (
	Attachments                  TweetField = "attachments"
	TweetAuthorId                TweetField = "author_id"
	ContextAnnotations           TweetField = "context_annotations"
	ConversationId               TweetField = "conversation_id"
	TweetCreatedAt               TweetField = "created_at"
	TweetEntities                TweetField = "entities"
	Geo                          TweetField = "geo"
	TweetId                      TweetField = "id"
	TweetInReplyToUserId         TweetField = "in_reply_to_user_id"
	Lang                         TweetField = "lang"
	IncludeTweetNonPublicMetrics TweetField = "non_public_metrics"
	IncludeTweetPublicMetrics    TweetField = "public_metrics"
	IncludeTweetOrganicMetrics   TweetField = "organic_metrics"
	IncludeTweetPromotedMetrics  TweetField = "promoted_metrics"
	PossiblySensitive            TweetField = "possibly_sensitive"
	ReferencedTweets             TweetField = "referenced_tweets"
	ReplySettings                TweetField = "reply_settings"
	Source                       TweetField = "source"
	Text                         TweetField = "text"
	TweetWithheld                TweetField = "withheld"
)

func WithTweetFields(fields ...TweetField) RequestOption {
	values := make([]string, len(fields))
	for i, v := range fields {
		values[i] = string(v)
	}
	return WithQueryParams("tweet.fields", values...)
}

type UserField string

const (
	CreatedAt                UserField = "created_at"
	Description              UserField = "description"
	IncludeEntities          UserField = "entities"
	Id                       UserField = "id"
	Location                 UserField = "location"
	Name                     UserField = "name"
	PinnedTweetId            UserField = "pinned_tweet_id"
	ProfileImageUrl          UserField = "profile_image_url"
	Protected                UserField = "protected"
	IncludeUserPublicMetrics UserField = "public_metrics"
	UserUrl                  UserField = "url"
	Username                 UserField = "username"
	Verified                 UserField = "verified"
	UserWithheld             UserField = "withheld"
)

func WithUserFields(fields ...TweetField) RequestOption {
	values := make([]string, len(fields))
	for i, v := range fields {
		values[i] = string(v)
	}
	return WithQueryParams("user.fields", values...)
}

func WithEndTime(endTime time.Time) RequestOption {
	return WithQueryParams("end_time", endTime.Format(time.RFC3339))
}

func WithMaxResults(max int) RequestOption {
	return WithQueryParams("max_results", strconv.FormatInt(int64(max), 10))
}

func WithNextToken(token string) RequestOption {
	return WithQueryParams("next_token", token)
}

func WithQuery(query string) RequestOption {
	return WithQueryParams("query", query)
}

func WithSinceID(since string) RequestOption {
	return WithQueryParams("since_id", since)
}

func WithStartTime(startTime time.Time) RequestOption {
	return WithQueryParams("start_time", startTime.Format(time.RFC3339))
}

func WithUntilID(until string) RequestOption {
	return WithQueryParams("until_id", until)
}

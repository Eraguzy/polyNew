package messenger

type GetArgs struct {
	Param URLParam
	Value string
}

type URL string

const (
	URLGammaAPI       URL = "https://gamma-api.polymarket.com"
	URLTelegramBotAPI URL = "https://api.telegram.org"
)

type URLPath string

const (
	// polymarket
	PathEvents URLPath = "/events"
	PathSports URLPath = "/sports"
	PathTags   URLPath = "/tags"
	// telegram
	PathSendMessage URLPath = "/sendMessage"
)

// polymarket tag by slug specific
func PathTagsSlug(slug string) URLPath {
	return URLPath("/tags/slug/" + slug)
}
func TelegramBotPathBuilder(botToken string, path URLPath) URLPath {
	return URLPath("/bot" + botToken + string(path))
}

// url parameters
type URLParam string

const (
	// polymarket
	URLParamActive URLParam = "active"
	URLParamClosed URLParam = "closed"
	URLParamLimit  URLParam = "limit"
	URLParamOffset URLParam = "offset"
	URLParamTagID  URLParam = "tag_id"
	URLParamSlug   URLParam = "slug"
	// telegram
	URLParamChatID URLParam = "chat_id"
	URLParamText   URLParam = "text"
)

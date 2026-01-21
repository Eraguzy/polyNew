package messenger

type GetArgs struct {
	Param URLParam
	Value string
}

type URL string

const (
	URLGammaAPI URL = "https://gamma-api.messenger.com"
)

type URLPath string

const (
	PathEvents URLPath = "/events"
	PathSports URLPath = "/sports"
	PathTags   URLPath = "/tags"
)

func PathTagsSlug(slug string) URLPath {
	return URLPath("/tags/slug/" + slug)
}

// url parameters
type URLParam string

const (
	URLParamActive URLParam = "active"
	URLParamClosed URLParam = "closed"
	URLParamLimit  URLParam = "limit"
	URLParamOffset URLParam = "offset"
	URLParamTagID  URLParam = "tag_id"
	URLParamSlug   URLParam = "slug"
)

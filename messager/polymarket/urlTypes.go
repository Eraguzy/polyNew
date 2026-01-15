package polymarket

type GetArgs struct {
	Param URLParam
	Value string
}

type URL string

const (
	URLGammaAPI URL = "https://gamma-api.polymarket.com"
)

type URLPath string

const (
	PathEvents URLPath = "/events"
	PathSports URLPath = "/sports"
	PathTags   URLPath = "/tags"
)

// url parameters
type URLParam string

const (
	URLParamActive URLParam = "active"
	URLParamClosed URLParam = "closed"
	URLParamLimit  URLParam = "limit"
	URLParamOffset URLParam = "offset"
)

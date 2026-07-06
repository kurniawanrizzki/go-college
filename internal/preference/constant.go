package preference

type contextKey string

const (
	AppVersion string = "1.0.0"

	CONTEXT_KEY_LOG_TRACE_ID contextKey = "trace_id"
	CONTEXT_KEY_LOG_SPAN_ID  contextKey = "span_id"

	FormatError  string = "please check the format with your input"
	CommandError string = "the command of first number should > 0"

	EVENT      string = "event"
	METHOD     string = "method"
	URL        string = "url"
	ADDR       string = "addr"
	STATUS     string = "status_code"
	LATENCY    string = "latency"
	USER_AGENT string = "user_agent"

	APP_LANG string = `x-app-lang`

	LANG_EN string = `en`
	LANG_ID string = `id`

	HeaderAccessControlAllowOrigin  string = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowHeaders string = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowMethods string = "Access-Control-Allow-Methods"
	HeaderXFrameOptions             string = "X-Frame-Options"
	HeaderContentSecurityPolicy     string = "Content-Security-Policy"
	HeaderXXSSProtection            string = "X-XSS-Protection"
	HeaderStrictTransportSecurity   string = "Strict-Transport-Security"
	HeaderReferrerPolicy            string = "Referrer-Policy"
	HeaderXContentTypeOptions       string = "X-Content-Type-Options"
	HeaderPermissionsPolicy         string = "Permissions-Policy"

	// Allowed HTTP Methods
	AllowedMethods string = "GET, POST, PUT, DELETE"

	// CORS Security Values
	CSPValue               string = "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';"
	PermissionsPolicyValue string = "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()"

	HeaderContentType string = "Content-Type"
	ContentTypeJSON   string = "application/json"

	//API Routes
	RouteCreateCollege string = "/college/create"
	RouteGetColleges string = "/college/all"
)

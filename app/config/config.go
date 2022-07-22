package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"os"
	"time"
)

func Init() {
	// Set default configurations
	setDefaults()

	// Select the .env file
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	// Automatically refresh environment variables
	viper.AutomaticEnv()
}

var defaultErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Return HTTP response
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(code).SendString(err.Error())
}

func GetFiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:                   viper.GetBool("FIBER_PREFORK"),
		ServerHeader:              viper.GetString("FIBER_SERVERHEADER"),
		StrictRouting:             viper.GetBool("FIBER_STRICTROUTING"),
		CaseSensitive:             viper.GetBool("FIBER_CASESENSITIVE"),
		Immutable:                 viper.GetBool("FIBER_IMMUTABLE"),
		UnescapePath:              viper.GetBool("FIBER_UNESCAPEPATH"),
		ETag:                      viper.GetBool("FIBER_ETAG"),
		BodyLimit:                 viper.GetInt("FIBER_BODYLIMIT"),
		Concurrency:               viper.GetInt("FIBER_CONCURRENCY"),
		ReadTimeout:               viper.GetDuration("FIBER_READTIMEOUT"),
		WriteTimeout:              viper.GetDuration("FIBER_WRITETIMEOUT"),
		IdleTimeout:               viper.GetDuration("FIBER_IDLETIMEOUT"),
		ReadBufferSize:            viper.GetInt("FIBER_READBUFFERSIZE"),
		WriteBufferSize:           viper.GetInt("FIBER_WRITEBUFFERSIZE"),
		CompressedFileSuffix:      viper.GetString("FIBER_COMPRESSEDFILESUFFIX"),
		ProxyHeader:               viper.GetString("FIBER_PROXYHEADER"),
		GETOnly:                   viper.GetBool("FIBER_GETONLY"),
		ErrorHandler:              defaultErrorHandler,
		DisableKeepalive:          viper.GetBool("FIBER_DISABLEKEEPALIVE"),
		DisableDefaultDate:        viper.GetBool("FIBER_DISABLEDEFAULTDATE"),
		DisableDefaultContentType: viper.GetBool("FIBER_DISABLEDEFAULTCONTENTTYPE"),
		DisableHeaderNormalizing:  viper.GetBool("FIBER_DISABLEHEADERNORMALIZING"),
		DisableStartupMessage:     viper.GetBool("FIBER_DISABLESTARTUPMESSAGE"),
		ReduceMemoryUsage:         viper.GetBool("FIBER_REDUCEMEMORYUSAGE"),
	}
}

func setDefaults() {
	// Set default App configuration
	viper.SetDefault("APP_ADDR", ":8080")
	viper.SetDefault("APP_ENV", "local")

	// Set default database configuration
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USERNAME", "admin")
	viper.SetDefault("DB_PASSWORD", "masterkey")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_DATABASE", "db")

	//// Set default session configuration
	//viper.SetDefault("SESSION_PROVIDER", "mysql")
	//viper.SetDefault("SESSION_KEYPREFIX", "session")
	//viper.SetDefault("SESSION_HOST", "localhost")
	//viper.SetDefault("SESSION_PORT", 3306)
	//viper.SetDefault("SESSION_USERNAME", "fiber")
	//viper.SetDefault("SESSION_PASSWORD", "secret")
	//viper.SetDefault("SESSION_DATABASE", "boilerplate")
	//viper.SetDefault("SESSION_TABLENAME", "sessions")
	//viper.SetDefault("SESSION_LOOKUP", "cookie:session_id")
	//viper.SetDefault("SESSION_DOMAIN", "")
	//viper.SetDefault("SESSION_SAMESITE", "Lax")
	//viper.SetDefault("SESSION_EXPIRATION", "12h")
	//viper.SetDefault("SESSION_SECURE", false)
	//viper.SetDefault("SESSION_GCINTERVAL", "1m")

	// Set default Fiber configuration
	viper.SetDefault("FIBER_PREFORK", false)
	viper.SetDefault("FIBER_SERVERHEADER", "")
	viper.SetDefault("FIBER_STRICTROUTING", false)
	viper.SetDefault("FIBER_CASESENSITIVE", false)
	viper.SetDefault("FIBER_IMMUTABLE", false)
	viper.SetDefault("FIBER_UNESCAPEPATH", false)
	viper.SetDefault("FIBER_ETAG", false)
	viper.SetDefault("FIBER_BODYLIMIT", 4194304)
	viper.SetDefault("FIBER_CONCURRENCY", 262144)
	viper.SetDefault("FIBER_VIEWS", nil)
	viper.SetDefault("FIBER_VIEWS_DIRECTORY", "resources/views")
	viper.SetDefault("FIBER_VIEWS_RELOAD", false)
	viper.SetDefault("FIBER_VIEWS_DEBUG", false)
	viper.SetDefault("FIBER_VIEWS_LAYOUT", "")
	viper.SetDefault("FIBER_VIEWS_DELIMS_L", "{{")
	viper.SetDefault("FIBER_VIEWS_DELIMS_R", "}}")
	viper.SetDefault("FIBER_READTIMEOUT", 0)
	viper.SetDefault("FIBER_WRITETIMEOUT", 0)
	viper.SetDefault("FIBER_IDLETIMEOUT", 0)
	viper.SetDefault("FIBER_READBUFFERSIZE", 4096)
	viper.SetDefault("FIBER_WRITEBUFFERSIZE", 4096)
	viper.SetDefault("FIBER_COMPRESSEDFILESUFFIX", ".fiber.gz")
	viper.SetDefault("FIBER_PROXYHEADER", "")
	viper.SetDefault("FIBER_GETONLY", false)
	viper.SetDefault("FIBER_DISABLEKEEPALIVE", false)
	viper.SetDefault("FIBER_DISABLEDEFAULTDATE", false)
	viper.SetDefault("FIBER_DISABLEDEFAULTCONTENTTYPE", false)
	viper.SetDefault("FIBER_DISABLEHEADERNORMALIZING", false)
	viper.SetDefault("FIBER_DISABLESTARTUPMESSAGE", false)
	viper.SetDefault("FIBER_REDUCEMEMORYUSAGE", false)

	// Set default Custom Access Logger middleware configuration
	viper.SetDefault("MW_ACCESS_LOGGER_ENABLED", false)
	viper.SetDefault("MW_ACCESS_LOGGER_TYPE", "console")
	viper.SetDefault("MW_ACCESS_LOGGER_FILENAME", "access.log")
	viper.SetDefault("MW_ACCESS_LOGGER_MAXSIZE", 500)
	viper.SetDefault("MW_ACCESS_LOGGER_MAXAGE", 28)
	viper.SetDefault("MW_ACCESS_LOGGER_MAXBACKUPS", 3)
	viper.SetDefault("MW_ACCESS_LOGGER_LOCALTIME", false)
	viper.SetDefault("MW_ACCESS_LOGGER_COMPRESS", false)

	// Set default Force HTTPS middleware configuration
	viper.SetDefault("MW_FORCE_HTTPS_ENABLED", false)

	// Set default Force trailing slash middleware configuration
	viper.SetDefault("MW_FORCE_TRAILING_SLASH_ENABLED", false)

	// Set default HSTS middleware configuration
	viper.SetDefault("MW_HSTS_ENABLED", false)
	viper.SetDefault("MW_HSTS_MAXAGE", 31536000)
	viper.SetDefault("MW_HSTS_INCLUDESUBDOMAINS", true)
	viper.SetDefault("MW_HSTS_PRELOAD", false)

	// Set default Suppress WWW middleware configuration
	viper.SetDefault("MW_SUPPRESS_WWW_ENABLED", true)

	// Set default Fiber Cache middleware configuration
	viper.SetDefault("MW_FIBER_CACHE_ENABLED", false)
	viper.SetDefault("MW_FIBER_CACHE_EXPIRATION", "1m")
	viper.SetDefault("MW_FIBER_CACHE_CACHECONTROL", false)

	// Set default Fiber Compress middleware configuration
	viper.SetDefault("MW_FIBER_COMPRESS_ENABLED", false)
	viper.SetDefault("MW_FIBER_COMPRESS_LEVEL", 0)

	// Set default Fiber CORS middleware configuration
	viper.SetDefault("MW_FIBER_CORS_ENABLED", false)
	viper.SetDefault("MW_FIBER_CORS_ALLOWORIGINS", "*")
	viper.SetDefault("MW_FIBER_CORS_ALLOWMETHODS", "GET,POST,HEAD,PUT,DELETE,PATCH")
	viper.SetDefault("MW_FIBER_CORS_ALLOWHEADERS", "")
	viper.SetDefault("MW_FIBER_CORS_ALLOWCREDENTIALS", false)
	viper.SetDefault("MW_FIBER_CORS_EXPOSEHEADERS", "")
	viper.SetDefault("MW_FIBER_CORS_MAXAGE", 0)

	// Set default Fiber CSRF middleware configuration
	viper.SetDefault("MW_FIBER_CSRF_ENABLED", false)
	viper.SetDefault("MW_FIBER_CSRF_TOKENLOOKUP", "header:X-CSRF-Token")
	viper.SetDefault("MW_FIBER_CSRF_COOKIE_NAME", "_csrf")
	viper.SetDefault("MW_FIBER_CSRF_COOKIE_SAMESITE", "Strict")
	viper.SetDefault("MW_FIBER_CSRF_COOKIE_EXPIRES", "24h")
	viper.SetDefault("MW_FIBER_CSRF_CONTEXTKEY", "csrf")

	// Set default Fiber ETag middleware configuration
	viper.SetDefault("MW_FIBER_ETAG_ENABLED", false)
	viper.SetDefault("MW_FIBER_ETAG_WEAK", false)

	// Set default Fiber Expvar middleware configuration
	viper.SetDefault("MW_FIBER_EXPVAR_ENABLED", false)

	// Set default Fiber Favicon middleware configuration
	viper.SetDefault("MW_FIBER_FAVICON_ENABLED", false)
	viper.SetDefault("MW_FIBER_FAVICON_FILE", "")
	viper.SetDefault("MW_FIBER_FAVICON_CACHECONTROL", "public, max-age=31536000")

	// Set default Fiber Limiter middleware configuration
	viper.SetDefault("MW_FIBER_LIMITER_ENABLED", true)
	viper.SetDefault("MW_FIBER_LIMITER_MAX", 5)
	viper.SetDefault("MW_FIBER_LIMITER_EXPIRATION", "1m")

	// Set default Fiber Monitor middleware configuration
	viper.SetDefault("MW_FIBER_MONITOR_ENABLED", false)

	// Set default Fiber Pprof middleware configuration
	viper.SetDefault("MW_FIBER_PPROF_ENABLED", false)

	// Set default Fiber Recover middleware configuration
	viper.SetDefault("MW_FIBER_RECOVER_ENABLED", true)

	// Set default Fiber RequestID middleware configuration
	viper.SetDefault("MW_FIBER_REQUESTID_ENABLED", false)
	viper.SetDefault("MW_FIBER_REQUESTID_HEADER", "X-Request-ID")
	viper.SetDefault("MW_FIBER_REQUESTID_CONTEXTKEY", "requestid")

	// Set default Fiber Logger middleware configuration
	viper.SetDefault("MW_FIBER_LOGGER_ENABLED", true)
	viper.SetDefault("MW_FIBER_LOGGER_FORMAT", "${pid} ${locals:requestid} ${status} - ${method} ${path}\n")
	viper.SetDefault("MW_FIBER_LOGGER_TIMEFORMAT", "15:04:05")
	viper.SetDefault("MW_FIBER_LOGGER_TIMEINTERVAL", 500*time.Millisecond)
	viper.SetDefault("MW_FIBER_LOGGER_TIMEZONE", "Europe/Moscow")

	// Set  Fiber Helmet middleware configuration
	viper.SetDefault("MW_FIBER_HELMET_ENABLED", false)
	viper.SetDefault("MW_FIBER_HELMET_XSS_PROTECTION", "1; mode=block")
	viper.SetDefault("MW_FIBER_HELMET_CONTENT_TYPE_NOSNIFF", "nosniff")
	viper.SetDefault("MW_FIBER_HELMET_X_FRAMEOPTIONS", "SAMEORIGIN")
	viper.SetDefault("MW_FIBER_HELMET_HSTS_MAX_AGE", 0)
	viper.SetDefault("MW_FIBER_HELMET_HSTS_EXCLUDE_SUBDOMAINS", false)
	viper.SetDefault("MW_FIBER_HELMET_CONTENT_SECURITY_POLICY", "")
	viper.SetDefault("MW_FIBER_HELMET_CSP_REPORT_ONLY", false)
	viper.SetDefault("MW_FIBER_HELMET_HSTS_PRELOAD_ENABLED", false)
	viper.SetDefault("MW_FIBER_HELMET_REFERRER_POLICY", "")
	viper.SetDefault("MW_FIBER_HELMET_PERMISSION_POLICY", "")

	// Set Fiber Prometheus middleware configurations
	viper.SetDefault("MW_FIBER_PROMETHEUS_ENABLED", false)
	viper.SetDefault("MW_FIBER_PROMETHEUS_SERVICE_NAME", "my-service")

}

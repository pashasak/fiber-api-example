package middleware

import (
	"fiber-api-example/app/server/middleware/fiberprometheus"
	l "fiber-api-example/app/utils/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/expvar"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/spf13/viper"
)

func RegisterMiddlewares(app *fiber.App) {

	// Middleware - Recover
	if viper.GetBool("MW_FIBER_RECOVER_ENABLED") {
		app.Use(recover.New())
	}

	// Middleware - Custom Access Logger based on zap
	if viper.GetBool("MW_ACCESS_LOGGER_ENABLED") {
		app.Use(AccessLogger(&AccessLoggerConfig{
			Type:        viper.GetString("MW_ACCESS_LOGGER_TYPE"),
			Environment: viper.GetString("APP_ENV"),
			Filename:    viper.GetString("MW_ACCESS_LOGGER_FILENAME"),
			MaxSize:     viper.GetInt("MW_ACCESS_LOGGER_MAXSIZE"),
			MaxAge:      viper.GetInt("MW_ACCESS_LOGGER_MAXAGE"),
			MaxBackups:  viper.GetInt("MW_ACCESS_LOGGER_MAXBACKUPS"),
			LocalTime:   viper.GetBool("MW_ACCESS_LOGGER_LOCALTIME"),
			Compress:    viper.GetBool("MW_ACCESS_LOGGER_COMPRESS"),
		}))
	}

	// Middleware - Force HTTPS
	if viper.GetBool("MW_FORCE_HTTPS_ENABLED") {
		app.Use(ForceHTTPS())
	}

	// Middleware - Force trailing slash
	if viper.GetBool("MW_FORCE_TRAILING_SLASH_ENABLED") {
		app.Use(ForceTrailingSlash())
	}

	// Middleware - HSTS
	if viper.GetBool("MW_HSTS_ENABLED") {
		app.Use(HSTS(&HSTSConfig{
			MaxAge:            viper.GetInt("MW_HSTS_MAXAGE"),
			IncludeSubdomains: viper.GetBool("MW_HSTS_INCLUDESUBDOMAINS"),
			Preload:           viper.GetBool("MW_HSTS_PRELOAD"),
		}))
	}

	// Middleware - Suppress WWW
	if viper.GetBool("MW_SUPPRESS_WWW_ENABLED") {
		app.Use(SuppressWWW())
	}

	// Middleware - Recover
	if viper.GetBool("MW_FIBER_RECOVER_ENABLED") {
		app.Use(recover.New())
	}

	// TODO: Middleware - Basic Authentication

	// Middleware - Cache
	if viper.GetBool("MW_FIBER_CACHE_ENABLED") {
		app.Use(cache.New(cache.Config{
			Expiration:   viper.GetDuration("MW_FIBER_CACHE_EXPIRATION"),
			CacheControl: viper.GetBool("MW_FIBER_CACHE_CACHECONTROL"),
		}))
	}

	// Middleware - Compress
	if viper.GetBool("MW_FIBER_COMPRESS_ENABLED") {
		lvl := compress.Level(viper.GetInt("MW_FIBER_COMPRESS_LEVEL"))
		app.Use(compress.New(compress.Config{
			Level: lvl,
		}))
	}

	// Middleware - CORS
	if viper.GetBool("MW_FIBER_CORS_ENABLED") {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     viper.GetString("MW_FIBER_CORS_ALLOWORIGINS"),
			AllowMethods:     viper.GetString("MW_FIBER_CORS_ALLOWMETHODS"),
			AllowHeaders:     viper.GetString("MW_FIBER_CORS_ALLOWHEADERS"),
			AllowCredentials: viper.GetBool("MW_FIBER_CORS_ALLOWCREDENTIALS"),
			ExposeHeaders:    viper.GetString("MW_FIBER_CORS_EXPOSEHEADERS"),
			MaxAge:           viper.GetInt("MW_FIBER_CORS_MAXAGE"),
		}))
	}

	// Middleware - CSRF
	if viper.GetBool("MW_FIBER_CSRF_ENABLED") {
		app.Use(csrf.New(csrf.Config{
			TokenLookup: viper.GetString("MW_FIBER_CSRF_TOKENLOOKUP"),
			Cookie: &fiber.Cookie{
				Name:     viper.GetString("MW_FIBER_CSRF_COOKIE_NAME"),
				SameSite: viper.GetString("MW_FIBER_CSRF_COOKIE_SAMESITE"),
			},
			CookieExpires: viper.GetDuration("MW_FIBER_CSRF_COOKIE_EXPIRES"),
			ContextKey:    viper.GetString("MW_FIBER_CSRF_CONTEXTKEY"),
		}))
	}

	// Middleware - ETag
	if viper.GetBool("MW_FIBER_ETAG_ENABLED") {
		app.Use(etag.New(etag.Config{
			Weak: viper.GetBool("MW_FIBER_ETAG_WEAK"),
		}))
	}

	// Middleware - Expvar
	if viper.GetBool("MW_FIBER_EXPVAR_ENABLED") {
		app.Use(expvar.New())
	}

	// Middleware - Favicon
	if viper.GetBool("MW_FIBER_FAVICON_ENABLED") {
		app.Use(favicon.New(favicon.Config{
			File:         viper.GetString("MW_FIBER_FAVICON_FILE"),
			CacheControl: viper.GetString("MW_FIBER_FAVICON_CACHECONTROL"),
		}))
	}

	// TODO: Middleware - Filesystem

	// Middleware - Limiter
	if viper.GetBool("MW_FIBER_LIMITER_ENABLED") {
		app.Use(limiter.New(limiter.Config{
			Max:        viper.GetInt("MW_FIBER_LIMITER_MAX"),
			Expiration: viper.GetDuration("MW_FIBER_LIMITER_EXPIRATION"),
			// TODO: Key
			// TODO: LimitReached
		}))
	}

	// Middleware - Monitor
	if viper.GetBool("MW_FIBER_MONITOR_ENABLED") {
		app.Use(monitor.New())
	}

	// Middleware - Pprof
	if viper.GetBool("MW_FIBER_PPROF_ENABLED") {
		app.Use(pprof.New())
	}

	// TODO: Middleware - Proxy

	// Middleware - RequestID
	if viper.GetBool("MW_FIBER_REQUESTID_ENABLED") {
		app.Use(requestid.New(requestid.Config{
			Header: viper.GetString("MW_FIBER_REQUESTID_HEADER"),
			// TODO: Generator
			ContextKey: viper.GetString("MW_FIBER_REQUESTID_CONTEXTKEY"),
		}))
	}

	// TODO: Middleware - Timeout

	// Middleware - Logger
	if viper.GetBool("MW_FIBER_LOGGER_ENABLED") {
		app.Use(logger.New(logger.Config{
			Format:       viper.GetString("MW_FIBER_LOGGER_FORMAT"),
			TimeFormat:   viper.GetString("MW_FIBER_LOGGER_TIMEFORMAT"),
			TimeInterval: viper.GetDuration("MW_FIBER_LOGGER_TIMEINTERVAL"),
			TimeZone:     viper.GetString("MW_FIBER_LOGGER_TIMEZONE"),
			Output:       &l.ZapWriter{l.GetLogger()},
			// TODO: Output
		}))
	}

	if viper.GetBool("MW_FIBER_HELMET_ENABLED") {
		app.Use(helmet.New(helmet.Config{
			XSSProtection:         viper.GetString("MW_FIBER_HELMET_XSS_PROTECTION"),
			ContentTypeNosniff:    viper.GetString("MW_FIBER_HELMET_CONTENT_TYPE_NOSNIFF"),
			XFrameOptions:         viper.GetString("MW_FIBER_HELMET_X_FRAMEOPTIONS"),
			HSTSMaxAge:            viper.GetInt("MW_FIBER_HELMET_HSTS_MAX_AGE"),
			HSTSExcludeSubdomains: viper.GetBool("MW_FIBER_HELMET_HSTS_EXCLUDE_SUBDOMAINS"),
			ContentSecurityPolicy: viper.GetString("MW_FIBER_HELMET_CONTENT_SECURITY_POLICY"),
			CSPReportOnly:         viper.GetBool("MW_FIBER_HELMET_CSP_REPORT_ONLY"),
			HSTSPreloadEnabled:    viper.GetBool("MW_FIBER_HELMET_HSTS_PRELOAD_ENABLED"),
			ReferrerPolicy:        viper.GetString("MW_FIBER_HELMET_REFERRER_POLICY"),
			PermissionPolicy:      viper.GetString("MW_FIBER_HELMET_PERMISSION_POLICY"),
		}))
	}

	if viper.GetBool("MW_FIBER_PROMETHEUS_ENABLED") {
		pr := fiberprometheus.New(viper.GetString("MW_FIBER_PROMETHEUS_SERVICE_NAME"))
		pr.RegisterAt(app, "/metrics")
		app.Use(pr.Middleware)
	}
}

package cmdargs

type RunArgs struct {
	DatabaseURL  string `mapstructure:"DATABASE_URL"`
	ShortenerURL string `mapstructure:"SHORTENER_URL"`
}

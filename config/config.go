package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Server struct {
		HTTPPort int32 `env:"HTTPPORT,default=3000"`
	}
}

var c config

const (
	//IMGPrefix is the imagen prefix
	IMGPrefix string = "IMG"
	//OtpEmailTemplate is the otp email template
	OtpEmailTemplate string = "email_otp.html"
)

//ReadConfig read config
func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	return err
}

// HTTPListener the listener string for the http service
func HTTPListener() string {
	return fmt.Sprintf(":%d", c.Server.HTTPPort)
}

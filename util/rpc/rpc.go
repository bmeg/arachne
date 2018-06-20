package rpc

import (
	"encoding/base64"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Config describes configuration for gRPC clients
type Config struct {
	User          string
	Password      string
	ServerAddress string
	// The timeout to use for making RPC client connections in nanoseconds
	// This timeout is Only enforced when used in conjunction with the
	// grpc.WithBlock dial option.
	Timeout time.Duration
	// The maximum number of times that a request will be retried for failures.
	// Time between retries follows an exponential backoff starting at 5 seconds
	// up to 1 minute
	MaxRetries uint
}

// ConfigWithDefaults returns a gRPC client config with default values set
func ConfigWithDefaults(serverAddress string) Config {
	return Config{
		User:          os.Getenv("ARACHNE_USER"),
		Password:      os.Getenv("ARACHNE_PASSWORD"),
		ServerAddress: serverAddress,
		Timeout:       30 * time.Second,
		MaxRetries:    10,
	}
}

// PerRPCPassword returns a new gRPC DialOption which includes a basic auth.
// password header in each RPC request.
func PerRPCPassword(user, password string) grpc.DialOption {
	return grpc.WithPerRPCCredentials(&loginCreds{
		User:     user,
		Password: password,
	})
}

type loginCreds struct {
	User     string
	Password string
}

func (c *loginCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	v := base64.StdEncoding.EncodeToString([]byte(c.User + ":" + c.Password))

	return map[string]string{
		"Authorization": "Basic " + v,
	}, nil
}

func (c *loginCreds) RequireTransportSecurity() bool {
	return false
}

// Dial returns a new gRPC ClientConn with some default dial and call options set
func Dial(pctx context.Context, conf Config, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(pctx, time.Duration(conf.Timeout))
	defer cancel()

	defaultOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		PerRPCPassword(conf.User, conf.Password),
	}
	opts = append(opts, defaultOpts...)
	opts = append(
		opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(conf.MaxRetries),
			grpc_retry.WithBackoff(newExponentialBackoff().BackoffWithJitter),
		)),
	)

	conn, err := grpc.DialContext(ctx,
		conf.ServerAddress,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	go func() {
		<-pctx.Done()
		conn.Close()
	}()
	return conn, nil
}

type exponentialBackoff struct {
	Initial             time.Duration
	Max                 time.Duration
	Multiplier          float64
	RandomizationFactor float64
}

func newExponentialBackoff() *exponentialBackoff {
	return &exponentialBackoff{
		Initial:             5 * time.Second,
		Max:                 1 * time.Minute,
		Multiplier:          2.0,
		RandomizationFactor: 0.1,
	}
}

func (b *exponentialBackoff) Backoff(attempt uint) time.Duration {
	return time.Duration(float64(b.Initial) * math.Pow(b.Multiplier, float64(attempt)))
}

func (b *exponentialBackoff) Jitter(val time.Duration) time.Duration {
	v := float64(val)
	delta := b.RandomizationFactor * v
	minInterval := v - delta
	maxInterval := v + delta

	// Get a random value from the range [minInterval, maxInterval].
	// The formula used below has a +1 because if the minInterval is 1 and the maxInterval is 3 then
	// we want a 33% chance for selecting either 1, 2 or 3.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return time.Duration(minInterval + (r.Float64() * (maxInterval - minInterval + 1)))
}

func (b *exponentialBackoff) BackoffWithJitter(attempt uint) time.Duration {
	nextBackoff := b.Jitter(b.Backoff(attempt))

	if nextBackoff > b.Max {
		return b.Jitter(b.Max)
	}

	return nextBackoff
}

///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"context"
	"flag"
	"fmt"
	fruitshop "fruitshop"
	cart "fruitshop/gen/cart"
	coupon "fruitshop/gen/coupon"
	discount "fruitshop/gen/discount"
	fruit "fruitshop/gen/fruit"
	payment "fruitshop/gen/payment"
	user "fruitshop/gen/user"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[fruitshop] ", log.Ltime)
	}

	// Initialize the services.
	var (
		userSvc     user.Service
		couponSvc   coupon.Service
		fruitSvc    fruit.Service
		cartSvc     cart.Service
		paymentSvc  payment.Service
		discountSvc discount.Service
	)
	{
		userSvc = fruitshop.NewUser(logger)
		couponSvc = fruitshop.NewCoupon(logger)
		fruitSvc = fruitshop.NewFruit(logger)
		cartSvc = fruitshop.NewCart(logger)
		paymentSvc = fruitshop.NewPayment(logger)
		discountSvc = fruitshop.NewDiscount(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		userEndpoints     *user.Endpoints
		couponEndpoints   *coupon.Endpoints
		fruitEndpoints    *fruit.Endpoints
		cartEndpoints     *cart.Endpoints
		paymentEndpoints  *payment.Endpoints
		discountEndpoints *discount.Endpoints
	)
	{
		userEndpoints = user.NewEndpoints(userSvc)
		couponEndpoints = coupon.NewEndpoints(couponSvc)
		fruitEndpoints = fruit.NewEndpoints(fruitSvc)
		cartEndpoints = cart.NewEndpoints(cartSvc)
		paymentEndpoints = payment.NewEndpoints(paymentSvc)
		discountEndpoints = discount.NewEndpoints(discountSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8080/api/v1"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, userEndpoints, couponEndpoints, fruitEndpoints, cartEndpoints, paymentEndpoints, discountEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

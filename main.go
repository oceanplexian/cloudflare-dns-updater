package main

import (
	"context"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zoneID := flag.String("zoneid", "", "Zone ID")
	timeout := flag.String("timeout", "", "Timeout in seconds")
	subdomain := flag.String("subdomain", "", "Subdomain to check")

	flag.Parse()

	if *zoneID == "" || *timeout == "" || *subdomain == "" {
		log.Fatal("Missing flag. Usage: <program> -zoneid=<zone_id> -timeout=<timeout> -subdomain=<subdomain>")
	}

	timeoutDuration, err := time.ParseDuration(*timeout + "s")
	if err != nil {
		log.Fatal(err)
	}

	// Create a Zap logger configuration for development
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// Construct a new API object using a global API key
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		logger.Error("Error creating Cloudflare API object", zap.Error(err))
	}

	// Run as a daemon
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)

		// Fetch Public IP
		expectedIP, err := GetPublicIP()
		if err != nil {
			logger.Error("Error fetching Public IP", zap.Error(err))
		}
		logger.Info("Public IP is ", zap.String("expectedIP", expectedIP))

		// Fetch DNS Records
		e, _, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(*zoneID), cloudflare.ListDNSRecordsParams{})
		if err != nil {
			logger.Error("Error fetching DNS records", zap.Error(err))
		}

		// Check each record
		for _, record := range e {
			if record.Name == *subdomain {
				if net.ParseIP(record.Content) == nil {
					logger.Error("Invalid IP",
						zap.String("Content", record.Content),
					)
				}

				if record.Content != expectedIP {
					logger.Warn("Cloudflare IP does not match expected IP",
						zap.String("CloudflareIP", record.Content),
						zap.String("ExpectedIP", expectedIP),
					)

					record, err = updateDNSRecord(ctx, api, *zoneID, record, expectedIP)
					if err != nil {
						logger.Error("Error updating DNS record", zap.Error(err))
					}

					logger.Info("Record updated successfully", zap.Any("Record", record))
				} else {
					logger.Info("Cloudflare IP matches expected IP",
						zap.String("CloudflareIP", record.Content),
						zap.String("ExpectedIP", expectedIP),
					)
				}
			}
		}

		cancel()

		// Wait until the next execution time
		time.Sleep(timeoutDuration)
	}
}

func updateDNSRecord(ctx context.Context, api *cloudflare.API, zoneID string, record cloudflare.DNSRecord, expectedIP string) (cloudflare.DNSRecord, error) {
	// Update the record with the expected IP
	return api.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateDNSRecordParams{ID: record.ID, Content: expectedIP})
}

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://ifconfig.me/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := net.ParseIP(string(ipBytes))
	if ip == nil {
		return "", errors.New("Invalid IP address")
	}

	return ip.String(), nil
}

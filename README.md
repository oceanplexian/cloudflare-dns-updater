# Cloudflare DNS Updater Script

The Cloudflare DNS Updater Script is a command-line tool written in Go that automatically checks and updates Cloudflare DNS records to ensure they match the expected IP address.

## Features

- Automatically fetches DNS records from Cloudflare.
- Compares the current IP address with the expected IP address.
- Updates the DNS records if they don't match the expected IP address.
- Supports subdomains within a specific Cloudflare zone.
- Configurable timeout for periodic checks.
- Uses the Cloudflare API to interact with the DNS records.
- Logs updates and warnings using the Uber Zap logging library.

## Prerequisites

- Go programming language (version X.X.X or later)
- Cloudflare API key or API token
- Cloudflare Zone ID for the relevant domain
- Environment variables: `CLOUDFLARE_API_KEY` (or `CLOUDFLARE_API_TOKEN`), and `CLOUDFLARE_API_EMAIL` (only for API key)

## Installation

1. Clone the repository or download the source code.
2. Ensure that Go is properly installed and configured.
3. Install the project dependencies by running: `go get -d ./...`

## Usage

To use the Cloudflare DNS Updater Script, follow these steps:

1. Set the required environment variables:
   - `CLOUDFLARE_API_KEY` or `CLOUDFLARE_API_TOKEN`: Your Cloudflare API key or API token.
   - `CLOUDFLARE_API_EMAIL` (only for API key): The email associated with your Cloudflare account.
2. Run the script with the desired flags:
Replace `<zone_id>` with the Cloudflare Zone ID for the relevant domain.
Replace `<timeout>` with the desired timeout duration in seconds.
Replace `<subdomain>` with the subdomain you want to check and update.

For example:
`go run main.go -zoneid=abc123 -timeout=60 -subdomain=sub.example.com`

3. The script will start running as a daemon and will periodically check and update the Cloudflare DNS records.

## Logging

The script uses the Uber Zap logging library for logging. The logs are configured to provide informative and colorized output. The log levels and log messages are designed to assist with monitoring and troubleshooting.

## Contributing

Contributions to the Cloudflare DNS Updater Script are welcome! If you have suggestions, bug reports, or feature requests, please open an issue or submit a pull request. Please follow the project's code of conduct.

## License

This project is licensed under the [MIT License](LICENSE).


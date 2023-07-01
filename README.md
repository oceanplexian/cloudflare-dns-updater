# Cloudflare DNS Updater Script

The Cloudflare DNS Updater Script is a command-line tool written in Go that automatically checks and updates Cloudflare DNS records to ensure they match the expected IP address. You 
could probably use https://github.com/K0p1-Git/cloudflare-ddns-updater instead, but I wanted a systemd service written in Go.

![Screenshot](https://raw.githubusercontent.com/oceanplexian/cloudflare-dns-updater/main/screenshot.jpg)

## Features

- Automatically fetches DNS records from Cloudflare.
- Compares the current IP address with the expected IP address.
- Updates the DNS records if they don't match the expected IP address.
- Supports subdomains within a specific Cloudflare zone.
- Configurable timeout for periodic checks.
- Uses the Cloudflare API to interact with the DNS records.
- Logs updates and warnings using the Uber Zap logging library.

## Prerequisites

- Go programming language (version 1.18.1 or later)
- Cloudflare API key or API token
- Cloudflare Zone ID for the relevant domain
- Environment variables: `CLOUDFLARE_API_KEY` (or `CLOUDFLARE_API_TOKEN`), and `CLOUDFLARE_API_EMAIL` (only for API key)

## Installation

1. Clone the repository or download the source code.
2. Ensure that Go (version 1.18.1 or later) is properly installed and configured.
3. Install the project dependencies by running: `go get -d ./...`

### Unit File Installation

To install the Cloudflare DNS Updater as a systemd service, follow these steps:

1. Open a terminal and navigate to the project directory.
2. Run the `install_unit.sh` script with sudo: `sudo bash install_unit.sh`
3. Follow the prompts and provide the requested information:
- Timeout (in seconds): Enter the desired timeout duration for periodic checks.
- Cloudflare API key: Enter your Cloudflare API key or API token.
- Cloudflare API email: Enter the email associated with your Cloudflare account (required only for API key).
- Zone ID: Enter the Cloudflare Zone ID for the relevant domain. You can find this by clicking "Overview" for the domain in the Cloudflare dashboard and copying the Zone ID from the API Zone ID section.
- Subdomain to check: Enter the subdomain you want to check and update. Use the full path format, e.g., `subdomain.domain.com`.

4. The script will install the systemd unit file and start the Cloudflare DNS Updater service.

## Usage

To use the Cloudflare DNS Updater Script without the systemd service, follow these steps:

1. Set the required environment variables:
- `CLOUDFLARE_API_KEY` or `CLOUDFLARE_API_TOKEN`: Your Cloudflare API key or API token.
- `CLOUDFLARE_API_EMAIL` (only for API key): The email associated with your Cloudflare account.
2. Run the script with the desired flags: `go run main.go -zoneid=<zone_id> -timeout=<timeout> -subdomain=<subdomain>`

Replace `<zone_id>` with the Cloudflare Zone ID for the relevant domain.
Replace `<timeout>` with the desired timeout duration in seconds.
Replace `<subdomain>` with the subdomain you want to check and update.

For example: `go run main.go -zoneid=abc123 -timeout=60 -subdomain=subdomain.example.com`
3. The script will start running as a daemon and will periodically check and update the Cloudflare DNS records.

## License
This project is licensed under the MIT License.


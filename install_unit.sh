#!/bin/bash

set -e

# Check if running with sudo
if [[ $EUID -ne 0 ]]; then
    echo "Please run this script with sudo."
    exit 1
fi

# Prompt for input
read -rp "Enter the timeout (in seconds): " timeout
read -rp "Enter the Cloudflare API key: " api_key
read -rp "Enter the Cloudflare API email: " api_email
read -rp "Enter the Zone ID: " zone_id
read -rp "Enter the Subdomain to check: " subdomain

# Prepare the unit file content
unit_file_content="
[Unit]
Description=Cloudflare DNS Updater
After=network.target

[Service]
Type=simple
ExecStart=/opt/source/cloudflare-dns-updater/cloudflare-dns-updater -zoneid=$zone_id -timeout=$timeout -subdomain=$subdomain
Environment=\"CLOUDFLARE_API_KEY=$api_key\"
Environment=\"CLOUDFLARE_API_EMAIL=$api_email\"

[Install]
WantedBy=multi-user.target
"

# Write the unit file
unit_file_path="/etc/systemd/system/cloudflare-dns-updater.service"
echo "$unit_file_content" > "$unit_file_path"

# Set permissions for the binary
binary_path="/opt/source/cloudflare-dns-updater/cloudflare-dns-updater"
chmod +x "$binary_path"

# Reload systemd daemon
systemctl daemon-reload

# Stop and disable the service if it's already running
systemctl stop cloudflare-dns-updater || true
systemctl disable cloudflare-dns-updater || true

# Enable and start the service
systemctl enable cloudflare-dns-updater
systemctl start cloudflare-dns-updater

echo "Cloudflare DNS Updater service installed and started successfully."


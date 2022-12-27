#!/usr/bin/env bash

# Copy unit file to /etc/systemd/system/
sudo cp furyd.service /etc/systemd/system/

# Reload all unit files
sudo /bin/systemctl daemon-reload

# Enable and start the service
sudo /bin/systemctl enable furyd.service
sudo /bin/systemctl restart furyd.service

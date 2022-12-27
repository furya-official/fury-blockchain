#!/usr/bin/env bash

# Disable, stop, and remove unit file
sudo /bin/systemctl disable furyd.service
sudo /bin/systemctl stop furyd.service
sudo rm /etc/systemd/system/furyd.service

# Reload all unit files and reset failed
sudo /bin/systemctl daemon-reload
sudo /bin/systemctl reset-failed

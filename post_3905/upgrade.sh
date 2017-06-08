#!/usr/bin/env bash
# /usr/local/bin/upgrade.sh
# Refreshes the APT repository applies any package upgrades.

sudo apt-get -y update
sudo apt-get -y upgrade
sudo apt-get -y dist-upgrade
sudo apt-get -y autoremove
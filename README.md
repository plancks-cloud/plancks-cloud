# Planck's Cloud

<p align="center">
  <img src="docs/logo.png" width="64px" />
  <img src="docs/logob.png" width="64px" />
  <img src="docs/logo.png" width="64px" />
  <img src="docs/logob.png" width="64px" />
  <img src="docs/logo.png" width="64px" />
  <img src="docs/logob.png" width="64px" />
  <img src="docs/logo.png" width="64px" />
  <img src="docs/logob.png" width="64px" />
  <img src="docs/logo.png" width="64px" />
  <img src="docs/logob.png" width="64px" />
</p>


<a href="https://coggle.it/diagram/XEgmhoO3UopF8htc/t/logo"><img src="https://img.shields.io/badge/Roadmap-Coggle-brightgreen.svg" /></a>
<img src="https://europe-west1-captains-badges.cloudfunctions.net/function-clone-badge-pc?project=plancks-cloud/plancks-cloud" />
<img src="https://goreportcard.com/badge/github.com/plancks-cloud/plancks-cloud">
<a href="https://codeclimate.com/github/plancks-cloud/plancks-cloud/maintainability"><img src="https://api.codeclimate.com/v1/badges/81aff827de3938808c2d/maintainability" /></a>
[![codebeat badge](https://codebeat.co/badges/25407218-e856-4f5e-ac7c-9d045dc0fe5a)](https://codebeat.co/projects/github-com-plancks-cloud-plancks-cloud-master)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)

Planck's Cloud turns every home into a data center. Host your next project from your own home with Planck's Cloud.

# How it works

Planck's Cloud allows you to run containers on your network at home and make endpoints available on the public Internet.

A simple command line interface allows you to change things simply and quickly.


# Architecture

TBA

# Setup

## Pre-install
- Buy a domain with DNS provided.
- Point the DNS record at your public IP address.
- Setup your DNS provider's DNS updater to keep up-to-date with your public IP address.
- Give your "server" a static IP on the network.
- Create a DST-NAT rule on your router to point at your servers's static IP address.
- Install docker on your server.
- Run `docker swarm init` on your server.

## Installation
### Install the CLI
Either
- `go get github.com/plancks-cloud/plancks-cli`

or 
- Run installer at <a href="https://github.com/plancks-cloud/plancks-cli/releases">https://github.com/plancks-cloud/plancks-cli</a>

### Install the Daemon
- TBA



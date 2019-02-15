# Planck's Cloud
<img align="right" width="290" height="290" src="docs/logo.png" />

[![](https://images.microbadger.com/badges/version/planckscloud/plancks-cloud.svg)](https://microbadger.com/images/planckscloud/plancks-cloud "Get your own version badge on microbadger.com")&nbsp;<a href="https://trello.com/b/NutXeZwS/plancks-roadmap"><img src="https://img.shields.io/badge/Roadmap-Trello-brightgreen.svg" /></a>
<a href="https://coggle.it/diagram/XEgmhoO3UopF8htc/t/logo"><img src="https://img.shields.io/badge/Ideas-Coggle-brightgreen.svg" /></a>&nbsp;[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)

# What is Planck's Cloud?

Planck's Cloud aims to turn every home into a data center. Our goal is to allow you to host your next project from your own home with Planck's Cloud. We believe anyone should be able to run their side project from home.

# How it works

Planck's Cloud allows you to run containers on your network at home and make endpoints available on the public Internet. Traffic entering your network will be proxied to the relevant containers &amp; services. Planck's Cloud automatically configures LetsEncrypt TLS and is able to serve up traffic over HTTPS. 

A simple command line interface allows you to change things on your servers and also to enable CI/CD pipelines.

# What can Planck's Cloud do?

- Host a website from home.
- Host an API available for your next React, Vue or Angular app.
- Run a database for your other services.
- Allows you the freedom, flexibility and control to setup your own cloud infrastructure
- Run your application in a configuration which would allow you to scale out at home or lift and shift to a cloud provider like AWS, GCP, Azure etc.

# Current Status for Plancks Cloud?

The features currently available are:
- Create, update and delete docker services.
- Create routes for ingress. Route traffic for a hostname to a service or machine on your network.
- SSL offloading. Expose endpoints with LetEncrypt provided HTTPS.

# Architecture and Design
<img src="https://goreportcard.com/badge/github.com/plancks-cloud/plancks-cloud">&nbsp;<a href="https://codeclimate.com/github/plancks-cloud/plancks-cloud/maintainability"><img src="https://api.codeclimate.com/v1/badges/81aff827de3938808c2d/maintainability" /></a>&nbsp;[![codebeat badge](https://codebeat.co/badges/25407218-e856-4f5e-ac7c-9d045dc0fe5a)](https://codebeat.co/projects/github-com-plancks-cloud-plancks-cloud-master)

<img align="center" width="800" src="docs/pc-arch.png" />

Planck's Cloud runs is an Open Source Golang app that runs inside a docker container. The standard method of deployment is to run the container as a service inside a docker swarm. Planck's Cloud communicates with the docker daemon to create, update and delete services. Docker also proxies traffic from 80 and 443 on the host to the Planck's Cloud container. Planck's Cloud matches DNS entries to reverse proxy traffic to various services hosted in the swarm or other computers on your network.

# Setup
[![Docker Pulls](https://img.shields.io/docker/pulls/planckscloud/plancks-cloud.svg?maxAge=86400)](https://hub.docker.com/r/planckscloud/plancks-cloud)
<img src="https://europe-west1-captains-badges.cloudfunctions.net/function-clone-badge-pc?project=plancks-cloud/plancks-cloud" /><br />

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

# Getting Involved

We're dieing to hear what you think about Plancks Cloud, do you like it? do you wish it actually worked? do you want help setting it up? did it help you? or did you want to help us instead?

We're so open to the thoughts and comments, feel free to open an issue on this repository.

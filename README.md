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

<p align="center">

[![](https://images.microbadger.com/badges/version/planckscloud/plancks-cloud.svg)](https://microbadger.com/images/planckscloud/plancks-cloud "Get your own version badge on microbadger.com")
[![Docker Pulls](https://img.shields.io/docker/pulls/planckscloud/plancks-cloud.svg?maxAge=86400)](https://hub.docker.com/r/planckscloud/plancks-cloud)
<a href="https://trello.com/b/NutXeZwS/plancks-roadmap"><img src="https://img.shields.io/badge/Roadmap-Trello-brightgreen.svg" /></a>
<a href="https://coggle.it/diagram/XEgmhoO3UopF8htc/t/logo"><img src="https://img.shields.io/badge/Ideas-Coggle-brightgreen.svg" /></a>
<img src="https://europe-west1-captains-badges.cloudfunctions.net/function-clone-badge-pc?project=plancks-cloud/plancks-cloud" /><br />
<br />
<img src="https://goreportcard.com/badge/github.com/plancks-cloud/plancks-cloud">
<a href="https://codeclimate.com/github/plancks-cloud/plancks-cloud/maintainability"><img src="https://api.codeclimate.com/v1/badges/81aff827de3938808c2d/maintainability" /></a>
[![codebeat badge](https://codebeat.co/badges/25407218-e856-4f5e-ac7c-9d045dc0fe5a)](https://codebeat.co/projects/github-com-plancks-cloud-plancks-cloud-master)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)
</p>
# What is Plancks Cloud?

Planck's Cloud is a private or public cloud that aims to turn every home into a datacenter. Host your next project from your own home with Planck's Cloud.

# How it works

Planck's Cloud allows you to run containers on your network at home and make endpoints available on the public Internet.

A simple command line interface allows you to change things simply and quickly.

# What can Plancks Cloud do?

- Pool all your computing resources together to power more intensive operations.
- Share your computing resources with your friends and families who would like their own computing cloud for their own little projects 
- Allows you the freedom, flexibility and control to setup your own cloud infrastructure
- Use old and outdated hardware to dedicate more resources for your plancks cloud infrastructure
- Planck's cloud will utilize resouces at efficiently rate by eliminating overheads which is not required.
- Gives you the flexibility to deploy your applications to your network with minimal time and effort.

# Current Status for Plancks Cloud?

Currently the plan for Planck's Cloud is to make it possible for anyone and everyone to be able to easily spin up their own Planck's Cloud network from any hardware that is available, and use it to run their projects.


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

# Getting Involved

We're dieing to hear what you think about Plancks Cloud, do you like it? do you wish it actually worked? do you want help setting it up? did it help you? or did you want to help us instead?

We're so open to the thoughts and comments, feel free to send us a message via the different channels on github, or via our email justin@plancks.cloud .

# Evoke 365 Web Server 2.0

![CI](https://github.com/evoke365/webserver/actions/workflows/go.yml/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/evoke365/webserver/branch/master/graph/badge.svg?token=P63AVPT88Z)](https://codecov.io/gh/evoke365/webserver)
[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=evoke365_webserver&metric=alert_status)](https://sonarcloud.io/dashboard?id=evoke365_webserver)

Current Status: In Progress

## Overview

The codebase of the next gen backend system for evoke365.net.

The API Interface layer is generated with [go-swagger](https://github.com/go-swagger/go-swagger).

The Controller layer is implemented with [Event Sourcing](https://www.martinfowler.com/eaaDev/EventSourcing.html) and [Pub-sub design pattern](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern).

The DataStore implementation is backed by [MongoDB](https://www.mongodb.com/).


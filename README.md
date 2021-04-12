[![Build Status](https://travis-ci.org/evoke365/webserver.svg?branch=master)](https://travis-ci.org/evoke365/webserver)
[![codecov](https://codecov.io/gh/evoke365/webserver/branch/master/graph/badge.svg)](https://codecov.io/gh/evoke365/webserver)

# Evoke 365 Web Server 2.0

Current Status: In Progress

## Overview

The codebase of the next gen backend system for evoke365.net.

The API Interface layer is generated with [go-swagger](https://github.com/go-swagger/go-swagger).

The Controller layer is implemented with [Event Sourcing](https://www.martinfowler.com/eaaDev/EventSourcing.html) and [Pub-sub design pattern](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern).

The DataStore implementation is backed by [MongoDB](https://www.mongodb.com/).

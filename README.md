# Shor.ty

## Overview

Shor.ty is a Golang-based link shortener with the following list of features:

* Configurable link shortening, based on random or pattern-based UUID
* Public link shortening
* Public API for creating short links with Swagger documentation and Swagger UI for testing

## Quickstart

Clone or download the project, and for each type of build below perform the listed steps:

### Plain build
* Write into .env file configuration settings to connect to PostgreSQL and Redis servers
* Ensure that you have Go in your PATH  
* Execute "make plain-build" - the resulting server executable will be compiled and placed into bin folder
* Execute "make plain-run" - .env file will be copied to bin folder, next to shorty executable that will be started

### Docker-compose build
* Login to dockerhub.com for docker images

## Roadmap
* Form validation for checking user input
* Private area with user registration and login
* User-based format for UUID
* Query parameters substitution
* Link analytics (collection)
* Locale choice from IP, session or user selection 
* Add translation (i18n) for text, errors and Swagger API documentation
* Rate limits for API
* Bulk import and export
* Link analytics (reports)
* QR code for the link

## Versioning

Shor.ty uses [Semantic Versioning](http://semver.org/)

## License

    Copyright (C) 2021 Daniel Protopopov

    This program is free software; you can redistribute it and/or
    modify it under the terms of the GNU General Public License
    as published by the Free Software Foundation; either version 2
    of the License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program; if not, write to the Free Software
    Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
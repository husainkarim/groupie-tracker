# GROUPIE TRACKER

Welcome to the documentation for the GROUPIE TRACKER. This file provides essential information about the application, its setup, features, and usage. Whether you're a developer, a tester, or a user, this guide will help you get started.

## Table of Contents

* Introduction
* Usage
* Features
* Configuration
* Troubleshooting
* Contributing
* License

## INTRODUCTION

GROUPIE TRACKER is a cutting-edge web application designed to display a list of artists and their details by fetching data from API saved. It's built with HTML/CSS/JS framework for frontend and GO for backend, making it exceptional among other applications.

## USAGE

To run the application:

```
go run main.go
```

## FEATURES

* Searching bar: enables searching by the name of the artist.
* filtering: enables filtering the results to only the artist that matches the search.
* Geolocalization: enables mapping the different concerts locations of a certain artist/band given by the Client.

## CONFIGURATION

* GROUPIE TRACKER web application connects to an outsourced JSON API to fetch data of artists.
* A welcome page will open once server is started and clicked on the webpage link 'http://localhost:8080'.
* A path handler function in Go handles the navigations in the backend.
* Home page fetches all artists' data.
* An empty artist page that displays names of all artist if none selected.
* An artist page that displays data of selected artist.

## TROUBLESHOOTING

An error handler is used to troubleshoot errors that occurs and display appropriate data accordingly.

## CONTRIBUTING

A lot of time and efforts were contributed to this project.

## LICENSE

Open-source software license.
Under Reboot01 mentoring.

## RESOURCES
https://developer.here.com/documentation/examples/maps-js/geoshapes/polyline-on-the-map
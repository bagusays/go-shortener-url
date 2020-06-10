# Go URL Shortener

## Introduction

> This is a server-side platform (API) for making URL shorter. URL shortening is a technique on the World Wide Web in which a Uniform Resource Locator may be made substantially shorter and still direct to the required page. This is achieved by using a redirect which links to the web page that has a long URL. -Wikipedia

## Installation

> You just restore the SQL file which is urlshortener.sql file for database handling, and then you just run like any others Go ecosystem

## Code Samples

For adding a short URL
```
Endpoint: http://localhost:8888/link/create
JSON Body:
    - shortUrl (string) (required)
    - longUrl (string) (required)
```
<br>
For editing a long URL
```
Endpoint: http://localhost:8888/link/edit
JSON Body:
    - shortUrl (string) (required)
    - longUrl (string) (required)
```
<br>
For deleting a long URL
```
Endpoint: http://localhost:8888/link/delete
JSON Body:
    - shortUrl (string) (required)
```
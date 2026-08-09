# GEO_API

# Description
The website is using API from 2GIS and sending picture of map to user.

# Stack

* Echo
* Golang
* Docker

# How to begin using website
To begin using the website you need to create `.env` file in the main root and fill the lines API_KEY and SERVER_PORT as in `.env.example`. To get free API from 2GIS you need to authorize here https://platform.2gis.ru/ru/keys and you can get free API for a month.
Also to download all dependencies write this code in console:

    go mod tidy

# Screenshots

Header

![alt text](web/assets/image2.png)

Main page

![alt text](web/assets/image1.png)

About page

![alt text](web/assets/image3.png)

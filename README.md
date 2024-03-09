# Ascii-Art-Web
Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of Ascii Art which involves recieving a string and then outputting the string in graphic representation using Ascii.

<img width="722" alt="image" src="https://github.com/RedshiftRach/Ascii-Art-Web/assets/159306159/d37eba8e-7165-4033-b8f7-635708837079">

Usage: How to Run

`go run ascii.go`

`localhost:8080`

Your webpage must allow the use of the different banners:

    shadow
    standard
    thinkertoy

Implement the following HTTP endpoints:

    GET /: Sends HTML response, the main page.
    1.1. GET Tip: go templates to receive and display data from the server.

    POST /ascii-art: that sends data to Go server (text and a banner)
    2.1. POST Tip: use form and other types of tags to make the post request.\

The way you display the result from the POST is up to you. What we recommend are one of the following :

    Display the result in the route /ascii-art after the POST is completed. So going from the home page to another page.
    Or display the result of the POST in the home page. This way appending the results in the home page.

The main page must have:

    text input
    radio buttons, select object or anything else to switch between banners
    button, which sends a POST request to '/ascii-art' and outputs the result on the page.

Your endpoints must return appropriate HTTP status codes.

    OK (200), if everything went without errors.
    Not Found, if nothing is found, for example templates or banners.
    Bad Request, for incorrect requests.
    Internal Server Error, for unhandled errors.
    
Implementation details: algorithm

The webpage uses a html option select to change between the different banners. After selecting an option from the menu and typing some text it will output the text in ascii art form. If you don't select any banner it will automatically select the standard one. You should be able to print normally.

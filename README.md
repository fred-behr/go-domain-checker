# Go Domain Name Checker

This is a script that can run through a list of domain names and check for availability and more.

## What it can do

- Check if a domain is available
- Check if a domain is for sale privately
- Check for several top-level domains (eg. .com, .tech etc.)

## Setup

Assuming you already have the source code in your editor.

1. Create a .env file and specify:

`API_KEY:{your_api_key}`

Replace "{your_api_key}" with an api key, that can be easily obtained and found for free [here](https://api-ninjas.com/profile)

<br>
<br>

2. Have Go installed on your machine, so you can run the script.

Macos:

`brew install go`

or follow the instructions on [the official Go website](https://go.dev/doc/install)

<br>
<br>

3. Modify the .txt files in /input to have the domain names and top level domain names, that you wish to check for.

<br>
<br>

4. Run the code.

Open terminal in the root and run:

`go run .`

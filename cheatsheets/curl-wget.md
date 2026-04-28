---
title: curl & wget
icon: fa-download
primary: "#6C5CE7"
lang: bash
---

## fa-globe Basic Requests

```bash
curl https://example.com              # fetch URL, print to stdout
curl -o page.html https://example.com # save to file
curl -O https://example.com/file.zip  # save with remote filename
curl -s https://example.com           # silent mode (no progress)
curl -i https://api.example.com       # include response headers
curl -I https://example.com           # HEAD request (headers only)

wget https://example.com/file.zip     # download file
wget -q https://example.com/file.zip  # quiet mode
wget -O output.zip https://example.com/file.zip  # save as custom name
wget -P /tmp/downloads https://example.com/file.zip  # save to directory
```

## fa-paper-plane HTTP Methods

```bash
curl -X POST https://api.example.com/users     # POST request
curl -X PUT https://api.example.com/users/1     # PUT request
curl -X DELETE https://api.example.com/users/1  # DELETE request
curl -X PATCH https://api.example.com/users/1   # PATCH request

# POST with form data
curl -d "name=Alice&age=25" https://api.example.com/users

# POST with JSON
curl -d '{"name":"Alice","age":25}' \
     -H "Content-Type: application/json" \
     https://api.example.com/users
```

## fa-key Headers & Auth

```bash
curl -H "Authorization: Bearer TOKEN" https://api.example.com  # bearer token
curl -H "X-Custom-Header: value" https://api.example.com       # custom header
curl -H "Accept: application/json" https://api.example.com      # accept JSON
curl -u user:pass https://api.example.com                       # basic auth
curl -u user https://api.example.com                            # prompt for password

wget --header="Authorization: Bearer TOKEN" https://api.example.com
wget --user=user --password=pass https://api.example.com        # basic auth
wget --header="Content-Type: application/json" https://api.example.com
```

## fa-file-arrow-up File Upload

```bash
curl -F "file=@photo.jpg" https://api.example.com/upload         # upload file
curl -F "file=@photo.jpg;type=image/jpeg" https://api.example.com  # with MIME type
curl -F "file=@/path/to/file" -F "name=Alice" https://api.example.com  # file + fields

# multipart form
curl -F "document=@report.pdf" \
     -F "title=Annual Report" \
     -F "category=finance" \
     https://api.example.com/upload
```

## fa-cookie Cookie Handling

```bash
curl -c cookies.txt https://example.com/login     # save cookies to file
curl -b cookies.txt https://example.com/dashboard  # send saved cookies
curl -b "session=abc123" https://example.com       # send cookie directly
curl -c cookies.txt -b cookies.txt https://example.com  # save and send

wget --save-cookies cookies.txt https://example.com/login
wget --load-cookies cookies.txt https://example.com/dashboard
```

## fa-arrows-rotate Redirect & Retry

```bash
curl -L https://example.com               # follow redirects
curl -L --max-redirs 5 https://example.com # max redirect hops
curl --retry 3 https://example.com         # retry on failure
curl --retry-delay 2 https://example.com   # wait 2s between retries
curl --connect-timeout 10 https://example.com  # connection timeout (seconds)
curl --max-time 30 https://example.com     # total timeout (seconds)

wget --max-redirect=5 https://example.com  # max redirects
wget --tries=3 https://example.com         # retry 3 times
wget --waitretry=5 https://example.com     # wait 5s between retries
wget --timeout=30 https://example.com      # timeout (seconds)
```

## fa-download Download & Resume

```bash
curl -C - -O https://example.com/large-file.zip  # resume download
curl -# -O https://example.com/file.zip           # progress bar

wget -c https://example.com/large-file.zip        # resume download
wget -r -np -k https://example.com/docs/          # mirror site (recursive)
wget --mirror https://example.com/                # full site mirror
wget --limit-rate=1m https://example.com/file.zip # limit download speed
wget -i urls.txt                                  # download from URL list
```

## fa-shield SSL & Proxy

```bash
curl -k https://example.com                # skip SSL verification
curl --cacert ca.pem https://example.com   # use custom CA cert
curl --cert client.pem --key key.pem https://example.com  # client cert
curl -x http://proxy:8080 https://example.com             # HTTP proxy
curl -x socks5://proxy:1080 https://example.com          # SOCKS5 proxy

wget --no-check-certificate https://example.com  # skip SSL verification
wget --ca-certificate=ca.pem https://example.com
wget -e use_proxy=yes http_proxy=http://proxy:8080 https://example.com
```

## fa-file-code Response Handling

```bash
curl -w "\nHTTP Code: %{http_code}\nTime: %{time_total}s\n" \
     -o /dev/null -s https://example.com       # show timing & status

curl -D headers.txt https://example.com         # dump response headers
curl -o response.json -w "%{http_code}" https://api.example.com  # save body + status

curl -s https://api.example.com | jq '.'        # pretty print JSON
curl -s https://api.example.com | python3 -m json.tool  # pretty print JSON

wget -S https://example.com                     # print server response headers
wget --server-response https://example.com      # print response headers
```

## fa-lightbulb Useful Patterns

```bash
# check if URL is reachable
curl -s -o /dev/null -w "%{http_code}" https://example.com

# download with authentication and resume
curl -L -C - -u user:pass -o file.zip https://example.com/file.zip

# POST JSON and pretty print response
curl -s -X POST -H "Content-Type: application/json" \
     -d '{"key":"value"}' https://api.example.com | jq '.'

# debug HTTP request
curl -v https://example.com                    # verbose output
curl --trace-ascii debug.log https://example.com  # full trace to file

# download multiple files
curl -O https://example.com/a.txt -O https://example.com/b.txt

# mirror website with wget
wget --mirror --convert-links --adjust-extension --page-requisites \
     --no-parent https://example.com/docs/
```

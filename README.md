Simple go web wrapper for getting images from Figma API. Useful for demo sites and automated updating of resources on site. DO NOT use direct urls from this service in production, Figma has very slow export API, so your site will be laggy

## Installation


* run `git clone https://git.shbbl.ru/shabablinchikow/figma-api.git`
* run `sudo ./install.sh`
* add code from `nginx.conf` to your configuration of nginx site (it's something like `/etc/nginx/sites-enabled/<sitename>.conf`
* run `sudo nginx -s reload`
* rename `config.json.example` to `config.json` at `/srv/figma/figma-api` and add to it your api token
* run `sudo systemctl start figma`

Now you can request your.site/figma?file=file_id&id=object_id and get a pretty image c:
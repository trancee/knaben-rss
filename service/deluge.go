package service

/*
request = "POST"
compressed
cookie = "cookie_deluge.txt"
cookie-jar = "cookie_deluge.txt"
header = "Content-Type: application/json"
header = "Accept: application/json"
url = "http://localhost:8112/json"
write-out = "\n"

curl -d '{"method": "auth.login", "params": ["deluge"], "id": 1}' -K curl.cfg
{
    "error": null,
    "id": 1,
    "result": true
}

curl -d '{"method": "web.connected", "params": [], "id": 1}' -K curl.cfg

curl -d '{"method": "web.get_hosts", "params": [], "id": 1}' -K curl.cfg
curl -d '{"method": "web.get_host_status", "params": ["<hostID>"], "id": 1}' -K curl.cfg

curl -d '{"method": "web.connect", "params": ["<hostID>"], "id": 1}' -K curl.cfg
curl -d '{"method": "web.disconnect", "params": [], "id": 1}' -K curl.cfg

curl -d '{"method": "core.add_torrent_magnet", "params": ["<magnet_uri>", {}], "id": 1}' -K curl.cfg

*/

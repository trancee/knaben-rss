package service

/*
# Settings
url = 'https://address.domain:port/transmission/rpc'
username = 'transmission'
password = 'secret'

# Get RPC Session ID
def get_session_id():
    sessionid_request = requests.get(url, auth=(username, password), verify=False)
    return sessionid_request.headers['x-transmission-session-id']

# Post Magnet Link
def post_link(magnetlink):
    sessionid = get_session_id()
    if sessionid:
        headers = {"X-Transmission-Session-Id": sessionid}
        body = dumps({"method": "torrent-add", "arguments": {"filename": magnetlink}})
        post_request = requests.post(url, data=body, headers=headers, auth=(username, password), verify=False)
        if str(post_request.text).find("success") == -1:
            title =  argv[0] + ' - Error'
            message = 'Magnet Link: ' + magnetlink + '\nAnswer: ' + post_request.text
            show_message(title, message)
        else:
            message = '\n'.join(names) + '\n' + '\n'.join(trackers)
            show_tray_message('Magnet Link Added', message)
*/

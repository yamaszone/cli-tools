import http.client
import json

BASE_URL = "08ad1pao69.execute-api.us-east-1.amazonaws.com"
ENDPOINT = "/dev/random_joke"
def get_json_payload():
    conn = http.client.HTTPSConnection(BASE_URL)
    conn.request('GET', ENDPOINT)
    response = conn.getresponse()
    return response.read().decode("utf-8")

jokes = json.loads(get_json_payload())
print("Setup:\t\t", jokes["setup"])
print("Punchline:\t", jokes["punchline"])

# Before refactor
#def get_json_payload2():
#    conn = http.client.HTTPSConnection("08ad1pao69.execute-api.us-east-1.amazonaws.com")
#    conn.request('GET', '/dev/random_joke')
#    response = conn.getresponse()
#    return response.read().decode("utf-8")

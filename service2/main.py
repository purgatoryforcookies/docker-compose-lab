import time
import requests
import os

print("Starting service2")

while True:
    
    if os.getenv("SERVICE1_URL") is None:
        print("SERVICE1_URL not set")
        exit(2)
    
    try:
        resp = requests.get(os.getenv("SERVICE1_URL"))

        if resp.status_code != 200:
            print("Requests failing!")

        else:
            print("Requests are working!")
        
        time.sleep(2)
    except Exception as e:
        print("Requests failing, no connection")
        
        time.sleep(2)


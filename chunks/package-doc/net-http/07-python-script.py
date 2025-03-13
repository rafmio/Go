import requests
import time

# list of routes
routes = ["/", "/about", "/contact"]

# number of iterations
iterations = len(routes)

# main loop
for _ in range(iterations):
    for route in routes:
        url = f"http://localhost:8080{route}"
        response = requests.get(url)
        print(response.text) # output server's response
        time.sleep(1)

# Load Balancer

Example configuration:

`````
services:
  - name: dummy                         # Service Name - Unique
    path: /service                      # Request path to be matched
    type: RoundRobin                    # *Type of algorithm
    servers:                            # Servers/endpoints to distrubute traffic to
      - host: http://host.org
        port: 8080
      - host: http://another-host.org
        port: 8188
``````

## Rate Limiter

Configured at the service level with maximum number of requests and window size in ms.

E.G.

max-requests = 2
window-size  = 2000ms

means that only two requests in time windows of 2 seconds are allowed.
In case maximum number of request is exceeded, 429 status code is returned.

`````
services:
  - name: dummy                                  
    path: /service   
    rate-limiter:
      name: dummy                       # Service Name - Unique
      max-requests: 2                   # Maximum number of requests allowed
      window-size: 20000                # Window size in ms. 
    
    
``````

* Currently only Round Robin is supported
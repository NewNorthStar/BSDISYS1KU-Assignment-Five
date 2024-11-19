# Distributed Auction System

## Service Functional Requirements

- Holds a single auction.
- Provides an API where clients can query on the status of the auction and place bids. 
- Auctions conclude after a predetermined duration. 
- For a bid to count, it must top the previously received bid by amount. 
- Bids of equal amount are time-compared on a first to process basis at the leader. Subsequent bids are considered _underbids_.
- For fairness, the system will satisfy _linearizability_. 

## Service Redundancy Requirements

- For redundancy, multiple backup nodes follow the leader of the auction. 
- If the leader fails, then the logically most up-to-date node should take over the auction. 
- A front-end service will 
    - allow clients to make contact with the auction leader, and
    - allow discovery of backup front-ends for clients to retain in case of front-end failure. 
- Consider the possibility of clients using different front-end nodes simultaneously. 

## Client Functional Requirements

- Clients bid on the auction. 
- Clients want to poll the state of the auction at times. 
- Clients have an ID which is consistent through the duration of the auction, even if they are not in permanent contact with the auction service. 
- Clients bid up the item as the auction progresses. 
- Clients discover front-end services for the auction through the discovery service provided. Should a front-end node fail, then it will try to use one of the other known front-ends. 
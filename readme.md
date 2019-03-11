## UID Generator

It has a simple interface.

	GenUID() uint64
	
	
Generates a partially-ordered UIDs based on timestamp (in seconds), server id, and an in-server counter.

Timestamp takes first 4 bytes, so sorting by this ID will IDs by the time of creation with up to a second precision.

This will not be a sequence in which they were generated, unless you only have a single server (or very lucky)  

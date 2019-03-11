## UID Generator

UID generator with a simple interface that does not require a global (project-wide) counter.

Interface is as follows:

	GenUID() uint64
	
	
Generates a partially-ordered UIDs based on timestamp (in seconds), server id, and an in-server counter.

Timestamp takes first 4 bytes, so sorting by this ID will sort IDs by the time of creation with up to a second precision.

This will not be a sequence in which they were generated, unless you only have a single server (or very lucky)  

| Bits | Format          | Comment                                                       | 
| ---- | --------------- | ------------------------------------------------------------- | 
| 32   | Unix Timestamp  |                                                               |
| 8    | Server id       | 256 Servers is enough for any project !                       | 
| 24   | Server Counter  | It limits this generator to 16m new IDs per second per server |

 
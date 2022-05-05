# hyperism-go
Backend REST-api server for hyperism 

## Quick Overview

### Step 1. Deploy REST API, Database and IPFS
```bash
git clone https://github.com/hyperism/hyperism-go
cd hyperism-go
docker-compose up -d
```

### Step 2. Private IPFS Peer-to-Server Setup
```bash
# Enter Server bash
docker exec -it server.hyperism.com bash
ipfs bootstrap rm --all
ipfs id -f='<addrs>' # Keep ipfs address except 127.0.0.1
exit
# Enter Peer bash
docker exec -it peer.hyperism.com bash
ipfs bootstrap rm --all
ipfs id -f='<addrs>' # Keep ipfs address except 127.0.0.1
ipfs bootstrap add <server ipfs address>
exit
# Enter Server bash
docker exec -it server.hyperism.com bash
ipfs bootstrap add <peer ipfs address>
exit
# Local bash
cd swarmkeygen && go run . generate > ../ipfs/swarm.key && cd ..
docker cp ipfs/swarm.key server.hyperism.com:/var/ipfsfb
docker cp ipfs/swarm.key peer.hyperism.com:/var/ipfsfb
docker-compose restart
# Check IPFS demo
./ipfs/e2e/test.sh p2s server.hyperism.com peer.hyperism.com
```
### Step 3. Rest API Usage
```bash
GET localhost:3000/api/catchphrases
GET localhost:3000/api/catchphrases/:id
POST localhost:3000/api/catchphrases
PATCH localhost:3000/api/catchphrases/:id
DELETE localhost:3000/api/catchphrases/:id
```

# Demo
TBA..

## Contact

You can contact me via e-mail (sinjihng at gmail.com). I am always happy to answer questions or help with any issues you might have, and please be sure to share any additional work or your creations with me, I love seeing what other people are making.

## License
<img align="right" src="http://opensource.org/trademarks/opensource/OSI-Approved-License-100x137.png">

The class is licensed under the [MIT License](http://opensource.org/licenses/MIT):

Copyright (c) 2022 Team Hyperism
*   [Jihong Shin](https://github.com/Snowapril)
*   [Hyungeun Lee](https://github.com/leehyunk6310)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

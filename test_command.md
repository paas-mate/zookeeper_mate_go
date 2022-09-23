### start
```bash
docker run --rm ttbb/zookeeper:mate
```
### start with port expose
```bash
docker run -p 2181:2181 -e REMOTE_MODE="false" --rm ttbb/zookeeper:mate
```
### start tls,port expose
```bash
docker run -p 2181:2181 -p 2182:2182 -e REMOTE_MODE="false" -e ZOOKEEPER_TLS_ENABLE="true" --rm ttbb/zookeeper:mate
```

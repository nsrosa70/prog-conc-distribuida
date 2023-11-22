docker stop mqtt
docker rm mqtt
docker run -d --memory="6g" --cpus="5.0" --name mqtt -p 1883:1883 -v C:\Users\user\go\prog-conc-distribuida\distribuida\pubsub\mqtt\mosquitto.conf:/mosquitto/config/mosquitto.conf eclipse-mosquitto
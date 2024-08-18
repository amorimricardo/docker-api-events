import docker
import datetime

client = docker.DockerClient(base_url='tcp://127.0.0.1:2375')

for event in client.events(decode=True, filters={"event": ["die"]}):
  container_id = event['Actor']['ID']
  container_name = event['Actor']['Attributes']['name']
  epoch_time = event['time']
  time_str = datetime.datetime.fromtimestamp(epoch_time)
  print(f"Container {container_name} with ID {container_id} died at {time_str}")
  
  
  
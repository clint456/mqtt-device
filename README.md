# mqtt-device

## 1.启动edgex 
```bash
# 将docker-compose.override.yml中的`/home/clint/work/edgexFoundry/v3.1/mqtt-device/mqtt-mock/`修改为自己的路径
cd ./degex-compose/compose-builder/
sudo bash start.sh
```

## 2.启动mqtt-mock
```bash
cd ./mqtt-mock/mock-scripts
# 将mock-device.sh中的`/home/clint/work/edgexFoundry/v3.1/mqtt-device/mqtt-mock`改为自己的路径
sudo bash mock-device.sh
```

## 3.登陆远程测试
docker-compose.yml中的微服务均设为了0.0.0.0，可以在宿主机直接`虚拟机ip:[微服务端口号]`访问ui管理界面



## 1.启动edgex 
```bash
cd ./degex-compose/compose-builder/
```
> 将docker-compose.override.yml中的`/home/clint/work/edgexFoundry/v3.1/mqtt-device/mqtt-mock/custom-config`修改为自己的路径
![image](https://github.com/user-attachments/assets/d7ef56a7-72e5-47fe-8e67-e7df88fa06be)

```bash
sudo bash start.sh
```

## 2.启动mqtt-mock
```bash
cd ./mqtt-mock/mock-scripts
```
> 将mock-device.sh中的`/home/clint/work/edgexFoundry/v3.1/mqtt-device/mqtt-mock/mock-scripts`改为自己的路径
![image](https://github.com/user-attachments/assets/253cbc2b-3dde-49ed-a1e2-eb97611fa04e)
```bash
sudo bash mock-device.sh
```
> 显示如下表示成功
![image](https://github.com/user-attachments/assets/65a65056-0543-4ab3-9c2e-f53cc68dde3a)



## 3.登陆远程测试
docker-compose.yml中的微服务均设为了0.0.0.0，可以在宿主机直接`虚拟机ip:[微服务端口号]`访问ui管理界面
![image](https://github.com/user-attachments/assets/660b4c3a-7a90-4934-9dd4-025ba8ac52e7)



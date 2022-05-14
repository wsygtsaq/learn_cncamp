1.创建namespace用于测试
kubectl create ns istio-test

2.创建configmap
kubectl create -f configmap.yaml -n istio-test

3.创建pod
kubectl create -f httpserver-metric.yaml -n istio-test

4.创建service
kubectl create -f httpserver-svc.yaml -n istio-test

5.测试服务访问
curl http://10.103.164.157/healthz
![1652535401068](https://wengao-1259242939.cos.ap-beijing.myqcloud.com/uPic/1652535401068.jpg)

6.创建istio网关
kubectl create -f istio-http.yaml -n istio-test

7.测试访问
export INGRESS_IP=10.98.144.74
kubectl get svc -nistio-system
curl -H "Host: httpserver.ldtest.io" $INGRESS_IP/healthz -v
![1652537289516](https://wengao-1259242939.cos.ap-beijing.myqcloud.com/uPic/1652537289516.jpg)

8.生成证书
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=ldtest Inc./CN=*.ldtest.io' -keyout ldtest.io.key -out ldtest.io.crt

9.创建secret
kubectl create -n istio-system secret tls ldtest-credential --key=ldtest.io.key --cert=ldtest.io.crt

10.创建istio https网关
kubectl create -f istio-https.yaml -n istio-test

11.测试
curl --resolve httpsserver.ldtest.io:443:$INGRESS_IP https://httpsserver.ldtest.io/healthz -v -k
![1652537175067](https://wengao-1259242939.cos.ap-beijing.myqcloud.com/uPic/1652537175067.jpg)
kubectl delete -f postgres.yaml &
kubectl delete -f redis.yaml &
kubectl delete -f rabbit.yaml &
kubectl delete -f user.yaml &
kubectl delete -f course.yaml &
kubectl delete -f exam.yaml &
kubectl delete -f overseer.yaml &
kubectl delete -f api_gateway.yaml
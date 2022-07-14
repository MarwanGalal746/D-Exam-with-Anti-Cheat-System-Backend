kubectl apply -f dexam-secret.yaml &
kubectl apply -f postgres.yaml &
kubectl apply -f redis.yaml &
kubectl apply -f rabbit.yaml &
kubectl apply -f user.yaml &
kubectl apply -f course.yaml &
kubectl apply -f exam.yaml &
kubectl apply -f overseer.yaml &
kubectl apply -f fr.yaml &
kubectl apply -f api_gateway.yaml
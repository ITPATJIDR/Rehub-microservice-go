dev:
	gin -i -a 8080 -p 3000 run .

build:
	docker stop rehub-microservice || true
	docker system prune -a -f
	docker build -t rehub-microservice .
	docker run --name rehub-microservice -d -p 8080:8080 rehub-microservice

stop:
	docker stop rehub-microservice

update_minikube: 
	cd k8
	kubectl apply -f ~/Code/CoSI/Rehab-Microservice/k8/deployment-go.yaml
	kubectl apply -f ~/Code/CoSI/Rehab-Microservice/k8/service-go.yaml


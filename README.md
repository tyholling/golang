# golang

Example project to collect metrics from clients over gRPC.

## Build

1. Requires: `go`, `gcc`, `make`, `podman`

1. `make setup` to install the build tools

1. `make builder` to build the builder image

1. `make build` to build

1. `make check` to run the checkers

1. `make images` to build the images

## Deploy

1. `cd deploy`

1. Install [`ingress-nginx`](https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx)

		kubectl create namespace ingress-nginx
		helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
		helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx

1. Install [`prometheus`](https://artifacthub.io/packages/helm/prometheus-community/prometheus)

		helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
		helm install prometheus prometheus-community/prometheus -f prometheus-values.yaml
		kubectl apply -f prometheus-ingress.yaml

1. Install [`grafana`](https://artifacthub.io/packages/helm/grafana/grafana)

		helm repo add grafana https://grafana.github.io/helm-charts
		helm install grafana grafana/grafana -f grafana-values.yaml \
		--set-file dashboards.default.client-metrics.json=client-metrics.json
		kubectl apply -f grafana-ingress.yaml
		kubectl get secret grafana -o jsonpath="{.data.admin-password}" | base64 --decode; echo

1. Deploy http service

		kubectl apply -f http-deployment.yaml
		kubectl apply -f http-service.yaml
		kubectl apply -f http-ingress.yaml

1. Deploy grpc service

		kubectl apply -f server-deployment.yaml
		kubectl apply -f server-service.yaml

## Collect metrics

1. `./bin/client | jq .msg` to run the client

1. Navigate to `/grafana` and open the Client Metrics dashboard

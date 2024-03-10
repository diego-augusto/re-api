dc:
	docker build --build-arg REACT_APP_API_URL=${API_URL} -t re-ui .

dc-run: dc
	docker run -d -p 80:80 re-ui

dc-stop:
	docker stop $(shell docker ps -q --filter ancestor=re-ui)
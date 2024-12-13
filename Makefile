target: go react

go:
	sudo systemctl restart postgresql.service
	air

react:
	cd ./web; npm start

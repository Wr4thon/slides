serve:
	present --http=0.0.0.0:3999 -play -use_playground -notes

serve-no-play:
	present --http=0.0.0.0:3999 -use_playground -notes

docker:
	docker build . -t docker.io/wr4thon/slides:latest

present:
	

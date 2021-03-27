build b:
	echo "[NASA GRAPH QL] Making binary in 'bin/main'" && go build -o bin/main src/main.go && echo "[NASA GRAPH QL] Done."

run r:
	echo "[NASA GRAPH QL] Running project" && go run src/main.go
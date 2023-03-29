run:
	go run cmd/main.go

logs:
	tail -f logs.txt

remotelogs:
	ssh pickup 'tail -f ~/PickupandGO/logs.txt'

# openremotelogs:
# 	ssh pickup 'vi ~/PickupandGO/logs.txt'

attach:
	tmux attach-session -t 0

remoteattach:
	ssh pickup 'tmux attach-session -t 0'
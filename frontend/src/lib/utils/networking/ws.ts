function NewWebSocketConnection() {
	// let wsEndpoint = 'ws://localhost:8000/ws';
	let wsEndpoint = 'wss://live-judge.herokuapp.com/ws';
	if (process.env.NODE_ENV === 'production') wsEndpoint = 'wss://' + process.env.HOST_NAME + '/ws';
	return new WebSocket(wsEndpoint);
}

export default function NewWebSocket(
	token: string,
	onMessage: (data: string) => void,
	onConnect?: () => void,
	onDisconnect?: () => void
) {
	const ws = NewWebSocketConnection();
	ws.onopen = () => {
		console.log('Connected to socket');
		if (onConnect) onConnect();
		ws.send(token);
	};

	ws.onmessage = (e: MessageEvent) => {
		onMessage(e.data);
	};

	ws.onerror = () => {
		console.error('Socket encountered error. Closing.');
		ws.close();
		if (onDisconnect) onDisconnect();
	};

	ws.onclose = () => {
		console.log('Disconnected from socket');
		if (onDisconnect) onDisconnect();
	};

	return ws;
}

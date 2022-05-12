import { WS_URL } from '$lib/data/constants/url';

function NewWebSocketConnection() {
	const wsEndpoint = WS_URL[process.env.NODE_ENV];
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

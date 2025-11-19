import ws from 'k6/ws';
import { check, sleep } from 'k6';

export default function () {
  const url = 'ws://localhost:9000/ws';
  const params = { tags: { my_tag: 'hello' } };

  const res = ws.connect(url, params, function (socket) {
    socket.on('open', () => console.log('connected'));
    socket.on('message', (data) => console.log('Message received: ', data));
    socket.on('close', () => console.log('disconnected'));
    socket.on('error', (e) => {
      if (e.error() != 'websocket: close sent') {
        console.log('An unexpected error occurred: ', e.error());
      }
    });

    socket.send('Hello from k6!');

    socket.setTimeout(function () {
      console.log('Closing the socket forcefully');
      socket.close();
    }, 3000);
  });

  check(res, { 'status is 101': (r) => r && r.status === 101 });
  sleep(1);
}

import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const url = 'http://localhost:2000/ping';

  const requestBody = {
    value: 'Pong'
  };

    const response = http.post(url, JSON.stringify(requestBody));

    check(response, {        'status is 200': (r) => r.status === 200,    });
  
}
import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 30,
  // A string specifying the total duration of the test run.
  duration: '60s',
};

// The function that defines VU logic.
//
// See https://grafana.com/docs/k6/latest/examples/get-started-with-k6/ to learn more
// about authoring k6 scripts.
//
export default function () {
  http.post('http://localhost:8080/items', JSON.stringify({
    "name": "lorem ipsum",
    "stock": 1
  }), {
    headers: { 'Content-Type': 'application/json' },
  })
  sleep(2);
}

export function handleSummary(data) {
  return {
    'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
    './summary.json': JSON.stringify(data), // and a JSON with all the details...
  };
}
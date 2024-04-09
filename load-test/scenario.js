import http from 'k6/http';
import { check, sleep } from 'k6';
import execution from 'k6/execution';
import { SharedArray } from 'k6/data';

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 3,
  // A string specifying the total duration of the test run.
  duration: '60s',
};

const sharedData = new SharedArray("credentials", () => {
  let data = JSON.parse(open('users.json'))
  console.log(data)
  return data
})

/**
 * user flow
 */
export default function () {
  let max_id = 100
  let id = Math.floor(Math.random() * max_id)
  let userCred = sharedData[execution.vu.idInTest - 1]
  console.log(userCred)

  // login to obtain JWT to get access to restricted resource
  const loginRes = http.post(`http://localhost:8080/login`, JSON.stringify({
    "email": userCred.email,
    "password": "password"
  }), {
    headers: { 'Content-Type': 'application/json' },
  })
  const token = loginRes.json().token
  console.log(token)
  sleep(1);

  // request params with authorization token
  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
  };

  // getting list of items in database
  let res = http.get(`http://localhost:8080/items`, params);
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  sleep(1);

  // adding new item to the database
  res = http.post(`http://localhost:8080/items`,
    JSON.stringify({
      "name": "items",
      "stock": 10,
    }),
    params);
  sleep(1);

  // get item data from the response of post request, and use the ite id to get the item detail
  let postedItem = res.json()
  res = http.get(`http://localhost:8080/items/${postedItem.item.ID}`, params);
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  sleep(1);

  res = http.put(`http://localhost:8080/items/${postedItem.item.ID}`,
    JSON.stringify({
      "stock": 15,
    }),
    params);
  sleep(1)
}

// export function handleSummary(data) {
//   return {
//     'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
//     './summary.json': JSON.stringify(data), // and a JSON with all the details...
//   };
// }
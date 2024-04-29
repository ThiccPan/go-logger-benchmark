import http from 'k6/http';
import { check, sleep } from 'k6';
import execution from 'k6/execution';
import { SharedArray } from 'k6/data';

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 3,
  // A string specifying the total iteration of the test execution.
  iterations: 10,
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
  const loginRes = http.post(
    `http://localhost:8080/login`,
    JSON.stringify({
      "email": userCred.email,
      "password": "password"
    }), {
    headers: { 'Content-Type': 'application/json' },
    tags: { name: "login request" }
  })
  const token = loginRes.json().token
  console.log(token)
  sleep(1);

  // getting list of items in database
  let res = http.get(`http://localhost:8080/items`, {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
    tags: { name: "get list of items request" }
  });
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  sleep(1);

  // adding new item to the database
  res = http.post(`http://localhost:8080/items`,
    JSON.stringify({
      "name": "items",
      "stock": 10,
    }), {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
    tags: { name: "add new item request" }
  });
  sleep(1);

  // get item data from the response of post request, and use the item id to get the item detail
  let postedItem = res.json()
  res = http.get(
    `http://localhost:8080/items/${postedItem.item.ID}`,
    {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      tags: { name: "get item by id request" }
    });

  check(res, {
    'is status 200': (r) => r.status === 200,
  });

  sleep(1);

  // update item data stock property
  res = http.put(`http://localhost:8080/items/${postedItem.item.ID}`,
    JSON.stringify({
      "stock": 15,
    }),
    {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      tags: { name: "update item request" }
    });
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  sleep(1)
}

// export function handleSummary(data) {
//   return {
//     'stdout': textSummary(data, { indent: ' ', enableColors: true }), // Show the text summary to stdout...
//     './summary.json': JSON.stringify(data), // and a JSON with all the details...
//   };
// }
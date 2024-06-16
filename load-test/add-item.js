import http from 'k6/http';
import { check, sleep } from 'k6';
import execution from 'k6/execution';
import { SharedArray } from 'k6/data';

export const options = {
    // A number specifying the number of VUs to run concurrently.
    vus: 100,
    // A string specifying the total duration of the test run.
    // duration: '60s',
    iterations: 3000,
};

const sharedData = new SharedArray("credentials", () => {
    let data = JSON.parse(open('users.json'))
    console.log(data)

    const vusNum = options.vus
    const usersCred = []
    for (let i = 0; i < vusNum; i++) {
        usersCred[i] = data[i]
    }
    return usersCred
})

export default function () {
    let userCred = sharedData[execution.vu.idInTest - 1]
    console.log(userCred)
    const loginRes = http.post(`http://localhost:8080/login`, JSON.stringify({
        "email": userCred.email,
        "password": "password"
    }), {
        headers: { 'Content-Type': 'application/json' },
    })

    const token = loginRes.json().token
    console.log(token)
    sleep(1);

    // adding new item to the database
    let res = http.post(`http://localhost:8080/items`,
        JSON.stringify({
            "name": "item lorem",
            "stock": 10,
        }), {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
    });

    check(res, {
        'is status 200': (r) => r.status === 200,
    });
    sleep(1);
}
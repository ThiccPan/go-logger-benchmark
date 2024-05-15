import http from 'k6/http';
import { check, sleep } from 'k6';
import execution from 'k6/execution';
import { SharedArray } from 'k6/data';

export const options = {
    // A number specifying the number of VUs to run concurrently.
    vus: 100,
    // A string specifying the total duration of the test run.
    duration: '1m',
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
    console.log("token:", token)
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
}
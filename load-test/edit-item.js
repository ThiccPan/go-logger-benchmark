import http from 'k6/http';
import { check, sleep } from 'k6';
import execution from 'k6/execution';
import { SharedArray } from 'k6/data';

export const options = {
    // A number specifying the number of VUs to run concurrently.
    vus: 3,
    // A string specifying the total iteration of the test execution.
    iterations: 100,
};

const sharedData = new SharedArray("credentials", () => {
    let data = JSON.parse(open('users.json'))
    console.log(data)
    return data
})

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

    // update item data stock property
    let stock = Math.floor(Math.random() * 100)
    let res = http.put(`http://localhost:8080/items/1`,
        JSON.stringify({
            "stock": stock,
        }),
        {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
        });

    check(res, {
        'is status 200': (r) => r.status === 200,
    });
    sleep(1)
}
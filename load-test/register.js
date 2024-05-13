import http from 'k6/http';
import { check, sleep } from 'k6';
import execution from 'k6/execution';
import { SharedArray } from 'k6/data';

export const options = {
    // A number specifying the number of VUs to run concurrently.
    vus: 1,
    // A string specifying the total duration of the test run.
    // duration: '96s',
    iterations: 1,
};

export default function () {
    // login to obtain JWT to get access to restricted resource
    for (let index = 35; index <= 100; index++) {
        const registerRes = http.post(`http://localhost:8080/register`, JSON.stringify({
            "username": `user${index}`,
            "email": `user${index}@gmail.com`,
            "password": "password"
        }), {
            headers: { 'Content-Type': 'application/json' },
        })
        sleep(0.5);
    }
}
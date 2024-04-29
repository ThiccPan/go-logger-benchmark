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
    let userCred = sharedData[execution.vu.idInTest - 1]
    console.log(userCred)

    
}
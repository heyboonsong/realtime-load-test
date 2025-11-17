import sse from "k6/x/sse";
import { check } from "k6"


export default function () {
    const url = "http://app:9000/sse"

    const response = sse.open(url, function (client) {
        client.on('open', function open() {
            console.log('connected')
        })

        client.on('event', function (event) {
            console.log(`event id=${event.id}, name=${event.name}, data=${event.data}`)
        })

        client.on('error', function (e) {
            console.log('An unexpected error occurred: ', e.error())
        })
    })

    check(response, { "success": (r) => r && r.status === 200 })
}
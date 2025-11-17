import http from 'k6/http'
import { check,sleep } from 'k6'

export default function () {
  let res = http.get('http://app:9000/polling')
  check(res, { 'success': (r) => r.status === 200 })
  sleep(1);

}

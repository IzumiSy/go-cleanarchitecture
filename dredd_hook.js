var hooks = require('hooks')

hooks.beforeAll(function(transactions) { })

session = {}

after('/login > POST', function(transaction) {
  session.token = JSON.parse(transaction.real.body).session_id
})

beforeEach(function(transaction) {
  if (session.token) {
    transaction.request.headers['Authorization'] = `bearer ${session.token}`
  }
})

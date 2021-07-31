var hooks = require('hooks')

hooks.beforeAll(function(transactions) { })

session = {
  token: null,
  email: null,
  todoName: null
}

hooks.beforeAll(function(transaction) {
  var prefix = Math.round(Math.random() * 10000)
  session.email = `user${prefix + ''}@example.com`
  session.todoName = `todo ${prefix + ''}`
  console.log("Email:", session.email)
})

hooks.before('/signup > POST', function(transaction) {
  var requestBody = JSON.parse(transaction.request.body);
  requestBody['email'] = session.email
  transaction.request.body = JSON.stringify(requestBody)
})

hooks.before('/login > POST', function(transaction) {
  var requestBody = JSON.parse(transaction.request.body);
  requestBody['email'] = session.email
  transaction.request.body = JSON.stringify(requestBody)
})

hooks.after('/login > POST', function(transaction) {
  session.token = JSON.parse(transaction.real.body).session_id
})

hooks.beforeEach(function(transaction) {
  if (session.token) {
    transaction.request.headers['Authorization'] = `Bearer ${session.token}`
  }
})

hooks.before('/todo > POST', function(transaction) {
  var requestBody = JSON.parse(transaction.request.body);
  requestBody['name'] = session.todoName
  transaction.request.body = JSON.stringify(requestBody)
})

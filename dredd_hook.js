var hooks = require('hooks')

hooks.beforeAll(function(transactions) {
  hooks.log('before all')
})

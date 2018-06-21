(function(Vue) {
  "use strict";
  
  Vue.component('todo-item', {
    props: ['todo'],
    template: '<li>{{ todo }}</li>'
  })

  var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue!',
      todos: ['We', 'make', 'it', 'awesome'],
      apiData: 0,
    },
    mounted: function() {
      console.log("this: ", this.$http)
      this.$http.get('/api/data').then(
        function(res) {
        console.log(res)
          if (res.data.success) {
            this.apiData = res.data.data ? res.data.data : 100
          }
        })
    },
    methods: {
      upcaseData: function (index) {
        alert(this.todos[index].toUpperCase())
      }
    }
  })
})(Vue);
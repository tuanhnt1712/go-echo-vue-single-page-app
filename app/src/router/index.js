import Vue from 'vue'
import Router from 'vue-router'
import Message from '@/components/Message'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
    	path: '/message',
    	name: 'Message',
    	component: Message
    }
  ]
})

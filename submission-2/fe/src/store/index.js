import { createStore } from 'vuex'
import taskStore from './task'

export default createStore({
  state: {
    countMe: 1
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    task: taskStore
  }
})

import axios from 'axios'

export default {
  namespaced: true,
  state: {
    count: 1,
    data: [],
    show_alert: false,
    status: '',
    message: '',
  },
  getters: {

  },
  mutations: {

  },
  actions: {
    async getData({ state }, params) {
      console.log('process getdata')
      try {

        const { data } = await axios.get(`http://localhost:3007/api/task`)
        
        state.data = data.data

        console.log('data ', data.data)

      } catch (err) {
        alert("Error")
      }
    },
    async storeData({ state }, params) {
      try {

        const {data} = await axios.post(`http://localhost:3007/api/task`, params, {
          headers: {
            'Content-Type': 'application/json'
          }
        })

        console.log('success ', data)

      } catch (err) {

        console.log('PRINT ERROR ', err)

      }
    },
  }
}
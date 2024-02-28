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
       
        state.show_alert = true
        state.status = "error"
        state.message = "Get data failed. try again later or contact your administrator"

      }
    },
    async storeData({ state }, params) {
      try {

        const {data} = await axios.post(`http://localhost:3007/api/task`, params, {
          headers: {
            'Content-Type': 'application/json'
          }
        })

        state.show_alert = true
        state.status = "success"
        state.message = "create data successfully"

      } catch (err) {

        state.show_alert = true
        state.status = "error"
        state.message = "create data failed. try again later or contact your administrator"

      }
    },
    async deleteData({ state }, id) {
      try {
        const { data } = await axios.delete(`http://localhost:3007/api/task/${id}`, {
          headers: {
            'Content-Type': 'application/json'
          }
        })

        state.show_alert = true
        state.status = "success"
        state.message = "delete data successfully"

      } catch (err) {
        
        state.show_alert = true
        state.status = "error"
        state.message = "delete data failed. try again later or contact your administrator"

      }
    }
  }
}
<template>
  <div class="task">
    
    <h2>TASK</h2>
    <div style="float: right;">
      <button class="button button1" type="button" @click="openModal">
        Add
      </button>
    </div>
    <table class="custom-table">
    <thead>
      <tr>
        <th v-for="column in columnHeader" :key="column">{{ column }}</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(row, index) in rowsData" :key="index">
       <td>
        {{ row.name }}
       </td>
       <td>
        {{ row.content }}
       </td>
       <td>
        {{ row.status }}
       </td>
       <td>
        {{ moment(row.created_at).format("D-M-YYYY H:ss") }} 
       </td>
       <td>
        <button class="button button2" type="button" @click="openModalDelete(row.id, row.name)">
          Delete
        </button>
       </td>
      </tr>
    </tbody>
  </table>

  <ModalComponent :isOpen="isModalOpened" @modal-close="closeModal" name="first-modal" />
  <ModalDeleteComponent :isOpen="isModalDeleteOpened" :idTask="idTask" :taskName="taskName" @modal-close="closeModalDelete" name="second-modal" />
  </div>
</template>
<script setup>
import moment from 'moment'
import { ref, onMounted, computed, watch } from 'vue'
import { useStore } from 'vuex'
import ModalComponent from '@/components/task/add.vue'
import ModalDeleteComponent from '@/components/task/delete'

// store
const store = useStore()

// data
const idTask = ref(0)
const taskName = ref("")
const form = ref({
  email: '',
  password: '',
  remember: false,
})
const isModalOpened = ref(false)
const isModalDeleteOpened = ref(false)
const rows = ref([
  {
    taskName: "Demo 1",
    createdBy: "Lorem ipsum",
    dateCreated: "2024-02-27",
    status: "on-going",
    actions: ""
  }
])
const columnHeader = ref(["Task Name", "Content", "Date Created", "Status", "actions"])

// computed
const rowsData = computed(() => {
  return store.state.task.data
})

// mounted
onMounted(() => {
  getData()
})

// methods
const getData = (() => { store.dispatch('task/getData') })
const openModal = () => {
  isModalOpened.value = true
}
const closeModal = () => {
  console.log('closeee')
  isModalOpened.value = false
  getData()
}

const openModalDelete = (id, name) => {
  idTask.value = id
  taskName.value = name
  isModalDeleteOpened.value = true
}
const closeModalDelete = () => {
  isModalDeleteOpened.value = false
}
</script>

<style scoped>
.custom-table {
  width: 100%;
  border-collapse: collapse;
}

.custom-table th,
.custom-table td {
  border: 1px solid #dddddd;
  padding: 8px;
  text-align: left;
}

.custom-table th {
  background-color: #f2f2f2;
}

.button {
  border: none;
  color: white;
  padding: 10px 15px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  margin: 4px 2px;
  transition-duration: 0.4s;
  cursor: pointer;
}

.button1 {
  background-color: white; 
  color: black; 
  border: 1px solid #04AA6D;
}

.button1:hover {
  background-color: #04AA6D;
  color: white;
}

.button2 {
  background-color: white; 
  color: black; 
  border: 1px solid #ff4d4d;
}

.button2:hover {
  background-color: #ff1a1a;
  color: white;
}
</style>
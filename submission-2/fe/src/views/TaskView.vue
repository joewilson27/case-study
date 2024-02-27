<template>
  <div class="task">
    
    <h2>TASK</h2>
    
    <table class="custom-table">
    <thead>
      <tr>
        <th v-for="column in columnHeader" :key="column">{{ column }}</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(row, index) in rows" :key="index">
        <td v-for="(value, key) in row" :key="key">
          <span v-if="key != 'actions'">
            {{ value }}
          </span>
          <span v-else>
            TOMBOL
          </span>
        </td>
      </tr>
    </tbody>
  </table>

  </div>
</template>
<script setup>
import { ref } from 'vue'

const form = ref({
  email: '',
  password: '',
  remember: false,
})

const rows = ref([
  {
    taskName: "Demo 1",
    createdBy: "John Doe",
    dateCreated: "2024-02-27",
    status: "on-going",
    actions: ""
  }
])

const columnHeader = ref(["Task Name", "Created By", "Date Created", "Status", "actions"])

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
</style>
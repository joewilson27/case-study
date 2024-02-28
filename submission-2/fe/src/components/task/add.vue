<script setup>
import { defineProps, defineEmits, ref } from "vue"
import { useStore } from 'vuex'

// store
const store = useStore()

// props
const props = defineProps({
  isOpen: Boolean,
})

// emits
const emit = defineEmits(["modal-close"])

const target = ref(null)

// data
const form = ref({
  name: "",
  content: "",
  status: "ongoing"
})
const errMessage = ref("")

// methods
const addHandler = () => {
  errMessage.value = ""
  if (form.value.name != ""
      && form.value.content != ""
      && form.value.status != "")
      {
        // process
        store.dispatch("task/storeData", form.value).then((data) => {
          emit("modal-close")
        })
        
      }
  else {
    // throw an error
    errMessage.value = "All input fields must be filled"
  }
}
</script>

<template>
  <div v-if="isOpen" class="modal-mask">
    <div class="modal-wrapper">
      <div class="modal-container" ref="target">
        <div class="modal-header">
          <slot name="header"> Add Task </slot>
          <p style="color:red" v-if="errMessage != ''">{{ errMessage }}</p>
        </div>
        <div class="modal-body">
          <slot name="content">
            <form v-on:submit.prevent="addHandler">
              <p>Name:
                <input v-model="form.name" placeholder="Task Name" />
              </p>
              <p>Content:
                <textarea v-model="form.content" placeholder="Task Name" />
              </p>
              <p>
                Status: 
                <select v-model="form.status">
                  <option>ongoing</option>
                  <option>complete</option>
                  <option>reject</option>
                </select>
              </p>
            </form>
          </slot>
        </div>
        <div class="modal-footer">
          <slot name="footer">
            <div>
              <button @click="addHandler">Submit</button>
            </div>
          </slot>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
}
.modal-container {
  width: 300px;
  margin: 150px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
}

</style>
<script setup>
import { defineProps, defineEmits, ref } from "vue"
import { useStore } from 'vuex'

// emits
const emit = defineEmits()

// store
const store = useStore()

// props
const props = defineProps({
  isOpen: Boolean,
  idTask: Number,
  taskName: String,
})

// method
const submitHandler = () => {
  if (props.idTask != "" || props.idTask != 0) {
    store.dispatch("task/deleteData", props.idTask)
         .then((data) => {
          emit("modal-close")
         })
  }
}

const closeHandler = () => {
  emit("modal-close")
}

</script>
<template>
  <div v-if="isOpen" class="modal-mask">
    <div class="modal-wrapper">
      <div class="modal-container" ref="target">
        <div class="modal-header">
          <h4 style="color: red"> WARNING! </h4>
        </div>
        <div class="modal-body">
          <br />
          <slot name="content">
            Are you sure want to delete this data? <strong>{{ taskName }}</strong> 
          </slot>

          <br />
          <br />
          <button class="button button1" type="button" @click="submitHandler">Yes</button>  |
          <button class="button button2" type="button" @click="closeHandler">Cancel</button>
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
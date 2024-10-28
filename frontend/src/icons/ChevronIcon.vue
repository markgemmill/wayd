<template>
   <svg v-if="expanded" 
      xmlns="http://www.w3.org/2000/svg" 
      :width="width" 
      :height="height" 
      @click="toggleExpand(false)"
      fill="currentColor" 
      class="bi bi-chevron-down" 
      viewBox="0 0 16 16"
      >
     <path fill-rule="evenodd" d="M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708"/>
   </svg>

   <svg v-else xmlns="http://www.w3.org/2000/svg" 
      :width="width" 
      :height="height" 
      @click="toggleExpand(true)"
      fill="currentColor" 
      class="bi bi-chevron-right" 
      viewBox="0 0 16 16"
      >
     <path fill-rule="evenodd" d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708"/>
   </svg> 
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const expanded = ref(false)

const { width=16, height=16, expand=false }= defineProps<{
    width: number
    height: number
    expand: boolean
}>()

const emit = defineEmits<{
    "expand": [boolean]
}>()

const toggleExpand = (expand: boolean) => {
    expanded.value = expand
    emit("expand", expand)
}

onMounted(() => {
    expanded.value = expand 
})

</script>

<style scoped>
</style>
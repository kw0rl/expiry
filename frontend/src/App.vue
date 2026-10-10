<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { checkHealth } from './services/api'

const status = ref<'loading' | 'connected' | 'error'>('loading')
const message = ref('Checking the connection…')

async function connect() {
  status.value = 'loading'
  message.value = 'Checking the connection…'
  try {
    await checkHealth()
    status.value = 'connected'
    message.value = 'The frontend and backend are connected.'
  } catch (error: unknown) {
    status.value = 'error'
    message.value = error instanceof Error ? error.message : 'Unable to reach the API.'
  }
}

onMounted(connect)
</script>

<template>
  <main>
    <p class="eyebrow">EXPIRY · PROJECT FOUNDATION</p>
    <h1>A little preparation.<br />Fewer surprises.</h1>
    <p class="intro">Keep track of what expires and when it needs your attention.</p>
    <section aria-labelledby="connection-title">
      <h2 id="connection-title">Connection check</h2>
      <p role="status" :class="status">{{ message }}</p>
      <button :disabled="status === 'loading'" @click="connect">
        {{ status === 'loading' ? 'Checking…' : 'Check again' }}
      </button>
    </section>
  </main>
</template>

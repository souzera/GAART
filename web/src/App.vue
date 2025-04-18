<script setup>
import LoginForm from './templates/LoginForm.vue';
</script>

<template>
  <div class="login-container">
    <img class="image-login" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcREyCF62WRPcmg4ET1T6LQqtvS3shJSXhElsQ&s" alt="sample image">
    <LoginForm/>
  </div>
</template>

<style scoped>

.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100vw;
  gap: 20px;
}

.image-login {
  width: 75vh;
  height: auto;
  margin-bottom: 20px;
}

</style>

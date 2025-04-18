<template>
    <div>
        <form class="login-form" @submit.prevent="handleSubmit">
            <h2 class="login-title">Login</h2>

            <div class="form-group">
                <label for="usuario">Usuario</label>
                <input type="text" id="usuario" v-model="usuario" placeholder="Seu usuário" required
                    class="form-input" />
            </div>

            <div class="form-group">
                <label for="senha">Senha</label>
                <input type="password" id="senha" v-model="senha" placeholder="Sua senha" required class="form-input" />
            </div>

            <div class="form-footer">
                <div class="remember-me">
                    <input type="checkbox" id="remember" v-model="remember">
                    <label for="remember" style="color: darkgray;">Lembrar-me</label>
                </div>
                <a href="#" class="forgot-password">Esqueceu a senha?</a>
            </div>

            <button type="submit" class="login-button">Entrar</button>

            <div class="register-link">
                Não tem uma conta? <a href="#">Registre-se</a>
            </div>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue';

import login from '../services/login.js';

const email = ref('');
const password = ref('');
const remember = ref(false);

const handleSubmit = async() => {
    login(usuario.value, senha.value)
    .then(response => {
        if (response.status === 200) {
            console.log('Login successful:', response.data);
        }
    })
};
</script>

<style scoped>
.login-form {
    width: 100%;
    max-width: 400px;
    padding: 30px;
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
}

.login-title {
    font-size: 24px;
    font-weight: 600;
    color: #333;
    margin-bottom: 30px;
    text-align: center;
}

.form-group {
    margin-bottom: 20px;
}

.form-group label {
    display: block;
    font-size: 14px;
    font-weight: 500;
    color: #555;
    margin-bottom: 6px;
}

.form-input {
    width: 100%;
    padding: 12px 15px;
    background-color: #f8f9fa;
    border: 1px solid #e1e5eb;
    border-radius: 6px;
    font-size: 14px;
    color: #333;
    transition: all 0.3s ease;
}

.form-input:focus {
    border-color: #4a6cf7;
    outline: none;
    box-shadow: 0 0 0 3px rgba(74, 108, 247, 0.15);
}

.form-footer {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    margin-bottom: 20px;
    font-size: 14px;
}

.remember-me {
    display: flex;
    align-items: center;
}

.remember-me input {
    margin-right: 6px;
}

.forgot-password {
    color: #4a6cf7;
    text-decoration: none;
    transition: color 0.2s;
}

.forgot-password:hover {
    color: #2a4cd7;
    text-decoration: underline;
}

.login-button {
    width: 100%;
    padding: 12px;
    background-color: #4a6cf7;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 16px;
    font-weight: 500;
    cursor: pointer;
    transition: background-color 0.2s;
}

.login-button:hover {
    background-color: #2a4cd7;
}

.register-link {
    text-align: center;
    margin-top: 20px;
    font-size: 14px;
    color: #555;
}

.register-link a {
    color: #4a6cf7;
    text-decoration: none;
    font-weight: 500;
    transition: color 0.2s;
}

.register-link a:hover {
    color: #2a4cd7;
    text-decoration: underline;
}
</style>
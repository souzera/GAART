const botao = document.querySelector('#doar01');

const disqueDenuncia = document.querySelector('#denuncia-section');

botao.addEventListener('click', () => {
    window.location.href = "donations.html";
})

disqueDenuncia.addEventListener('click', (email) =>{
    console.log("ola mundo")
})
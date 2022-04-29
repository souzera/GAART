const nome = document.querySelector('#nome-input');
const valor = document.querySelector('#valor-input');
const btnSubmit = document.querySelector('#btn-doar');

nome.addEventListener('click', ()=>{
    nome.value='';
});

valor.addEventListener('click', () =>{
    valor.value = "";
});
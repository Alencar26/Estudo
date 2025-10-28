console.log("Script carregado com sucesso!");
document.addEventListener("DOMContentLoaded", function() {
  document.body.addEventListener("htmx:configRequest", function(event) {
    
    const form = event.target;
    const username = form.querySelector('#username').value;
    const divError = form.querySelector('#username-error');

    if (username.length < 3) {
      event.preventDefault();
      divError.innerText = "O nome de usuário deve ter pelo menos 3 caracteres.";
    } else {
      divError.innerText = "";
    }
  });
});

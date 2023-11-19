const urlParams = new URLSearchParams(window.location.search);
const tabelaSelecionada = urlParams.get("tabela");

mostrarTituloTabela(tabelaSelecionada);
mostrarTabelaDashboard(tabelaSelecionada);

const profileBtn = document.getElementById("profile");
const profileMenu = document.getElementById("profile-menu");
profileBtn.addEventListener("click", mostrarMenuPerfilUsuario);
document.addEventListener("click", (e) => {
    // Fecha o menu de perfil quando a página é clicada
    if (!profileMenu.contains(e.target) && e.target !== profileBtn)
        profileMenu.style.display = "none";
});

const logoutBtn = document.getElementById("profile-log-out");
logoutBtn.addEventListener("click", (e) => {
    e.preventDefault(); 
    fazerLogout();
});


/**
 * Muda h1 (título) da tabela de forma dinâmica.
 * @param {string} nomeTabela - Nome da tabela selecionada
 * pelo usuário.
 */
function mostrarTituloTabela(nomeTabela) {
    const titulosTabelas = new Map();

    titulosTabelas.set("clientes", "Clientes");
    titulosTabelas.set("componentes", "Componentes");
    titulosTabelas.set("entradas", "Entradas");
    titulosTabelas.set("estoque", "Estoque");
    titulosTabelas.set("fabricantes", "Fabricantes");
    titulosTabelas.set("funcionarios", "Funcionários");
    titulosTabelas.set("ordemServico", "Ordens de serviço");
    titulosTabelas.set("saidas", "Saídas");

    if (titulosTabelas.get(nomeTabela)) {
        const tituloTabelaH1 = document.getElementById("titulo-tabela");
        tituloTabelaH1.textContent = titulosTabelas.get(nomeTabela);
    }
}

/**
 * Verifica a partir dos parâmetros da URL qual tabela
 * da dashboard deve ser mostrada e a expõe ao usuário.
 * @param {string} nomeTabela - Nome da tabela selecionada
 * pelo usuário.
 */
function mostrarTabelaDashboard(nomeTabela) {
    let section;
    const script = document.createElement("script");

    if (nomeTabela) {
        section = document.getElementById(nomeTabela);
        script.src = `/JS/dashboard/${nomeTabela}.js`;
    } else {
        section = document.getElementById("visao-geral");
        script.src = "/JS/dashboard/visaoGeral.js";
    }

    section.style.display = "block";
    document.querySelector("head").appendChild(script);
}

/**
 * Mostra o menu (dropdown) ao clicar na imagem do
 * perfil do usuário na navbar.
 */
function mostrarMenuPerfilUsuario() {
    if (profileMenu.style.display === "block") {
        profileMenu.style.display = "none";
    } else {
        profileMenu.style.display = "block";
    }
}

/**
 * Faz logout do sistema.
 */
function fazerLogout() {
    fetch("/sessao", {
        method: "DELETE"
    })
    .then((res) => res.json())
    .then((res) => {
        if (res.error) {
            mostrarMensagemErro(res.error);
            return;
        }

        window.location.reload();
    })
    .catch((err) => {
        mostrarMensagemErro("Erro ao conectar com o servidor. Tente novamente mais tarde.");
        console.error(err);
    });
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Erro a ser mostrado.
 */
function mostrarMensagemErro(erro) {
    const mensagemErroContainer = document.getElementById("mensagem-erro-container");
    const mensagemErro = document.getElementById("mensagem-erro");

    mensagemErroContainer.style.display = "block";

    mensagemErro.textContent = erro;

    setTimeout(() => {
        mensagemErroContainer.style.display = "none";
    }, 5000);
}

/**
 * Gera o popup de forma dinâmica.
 * @param {string} mensagem a ser mostrada no popup.
 */
function mostrarPopup(mensagem) {
    const popup = document.getElementById("popup");
    const mensagemTexto = document.getElementById("mensagem-popup");
    const btnVoltar = document.getElementById("popup-btn-voltar");
    const btnContinuar = document.getElementById("popup-btn-continuar");

    popup.style.visibility = "visible";

    const customPromise = new Promise((resolve, reject) => {
        btnVoltar.addEventListener("click", (e) => {
            e.preventDefault(); 
            popup.style.visibility = "hidden";
            reject(false);
        });
    
        btnContinuar.addEventListener("click", (e) => {
            e.preventDefault(); 
            popup.style.visibility = "hidden";
            //precisa retornar true (provavelmente por meio de uma promisse)
            resolve(true);
        });
    })

    mensagemTexto.innerHTML = mensagem;
    
    return customPromise
}
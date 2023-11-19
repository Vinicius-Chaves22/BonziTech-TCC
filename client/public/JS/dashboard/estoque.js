let dadosEstoque;
fetchEstoque()
.then((res) => {
    dadosEstoque = res;
    mostrarTabelaEstoque(dadosEstoque.estoque);
}); 

const addEstoqueFormBtn = document.getElementById("add-table-row-estoque");
const cancelarCriacaoEstoqueBtn = document.getElementById("cancel-btn-estoque");
addEstoqueFormBtn.addEventListener("click", mostrarFormCriacaoItemEstoque);
cancelarCriacaoEstoqueBtn.addEventListener("click", (e) => {
    e.preventDefault();
    const addEstoqueForm = document.getElementById("add-form-estoque");
    addEstoqueForm.style.display = "none";
});

const cancelarEdicaoEstoqueBtn = document.getElementById("cancel-btn-edit-estoque");
cancelarEdicaoEstoqueBtn.addEventListener("click", (e) => {
    e.preventDefault();
    const addEstoqueForm = document.getElementById("edit-form-estoque");
    addEstoqueForm.style.display = "none";
});

const searchBar = document.getElementById("search-bar");
searchBar.addEventListener("keyup", () => {
    const estq = procurarItemEstoque(searchBar.value.trim(), dadosEstoque.estoque);
    mostrarTabelaEstoque(estq);
});

const confirmarCriacaoEstoqueBtn = document.getElementById("confirm-btn-estoque");
confirmarCriacaoEstoqueBtn.addEventListener("click", async (event) => {
    event.preventDefault();

    const codCompTextbox = document.getElementById("add-codigo-componentes-estoque");
    const quantMinTextbox = document.getElementById("add-quantidade-min-estoque");
    const quantMaxTextbox = document.getElementById("add-quantidade-max-estoque");
    const quantidadeTextbox = document.getElementById("add-quantidade-estoque");

    await criarItemEstoque(
        Number(codCompTextbox.value.trim()),
        Number(quantMinTextbox.value.trim()),
        Number(quantMaxTextbox.value.trim()),
        Number(quantidadeTextbox.value.trim())
    );

    window.location.reload();
});

const confirmarAtualizacaoEstoqueBtn = document.getElementById("update-btn-edit-estoque");
confirmarAtualizacaoEstoqueBtn.addEventListener("click", async (event) => {
    event.preventDefault();

    const codEstqTextbox = document.getElementById("edit-codigo-estoque-estoque");
    const codCompTextbox = document.getElementById("edit-codigo-componentes-estoque");
    const quantMinTextbox = document.getElementById("edit-quantidade-min-estoque");
    const quantMaxTextbox = document.getElementById("edit-quantidade-max-estoque");
    const quantidadeTextbox = document.getElementById("edit-quantidade-estoque");

    await atualizarItemEstoque(
        Number(codEstqTextbox.value.trim()),
        Number(codCompTextbox.value.trim()),
        Number(quantMinTextbox.value.trim()),
        Number(quantMaxTextbox.value.trim()),
        Number(quantidadeTextbox.value.trim())
    );

    window.location.reload();
});


/**
 * Mostra tabela de estoque com seus devidos dados.
 * @param {Array} dadosTabela - Lista de estoque a serem mostrados.
 */
async function mostrarTabelaEstoque(dadosTabela) {
    const tbody = document.getElementById("tbody-estoque");
    tbody.innerHTML = "";

    if (dadosTabela.length === 0) {
        tbody.innerHTML = `
            <h2 class="texto-404"> Nenhum item encontrado no estoque. </h2>`;
    }

    for (const estq of dadosTabela) {
        const coluna = document.createElement("tr");
        coluna.innerHTML = `
        <td> ${estq.codEstq} </td>
        <td> ${estq.codComp} </td>
        <td> ${estq.min} </td>
        <td> ${estq.max} </td>
        <td> ${estq.quantidade} </td>
        `;

        const acoesCell = document.createElement("td");

        // Botão de inativação 
        const btnDelete = document.createElement("button");
        btnDelete.classList.add("delete-btn");
        btnDelete.addEventListener("click", (event) => {
            event.preventDefault();
            confirmarExclusaoEstoque()
            .then(() => {
                excluirItemEstoque(estq.codComp);
            })
                
                
        });
        btnDelete.innerHTML = '<i class="fa-solid fa-ban"> </i>';

        // Botão de edição 
        const btnEdit = document.createElement("button");
        btnEdit.classList.add("update-btn-icon");
        btnEdit.addEventListener("click", (event) => {
            event.preventDefault();
            mostrarFormEdicaoEstoque(estq);
        });
        btnEdit.innerHTML = '<i class="fa-solid fa-pen-to-square"> </i>';

        acoesCell.appendChild(btnDelete);
        acoesCell.appendChild(btnEdit);

        coluna.appendChild(acoesCell);

        tbody.appendChild(coluna);
    }
}

/**
 * Pega todos os item cadastrados no estoque pela API e
 * os insere na tabela.
 * @returns {object} - Resposta da API ou Objeto de erro.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function fetchEstoque() {
    return await fetch("/estoque")
    .then((res) => res.json())
    .then((res) => {
        if (res == null) 
            return null

        if (res.error) {
            mostrarMensagemErro(res.error);
            return new Error(res.error);
        }

        return res;
    })
    .catch((err) => {
        mostrarMensagemErro("Erro ao conectar com o servidor. Tente novamente mais tarde.");
        return new Error(err);
    });
}

/**
 * Envia dados de item do estoque para criação à API.
 * @param {Number} codComp - Código do componente.
 * @param {Number} min - Quantidade mínima aceitável no estoque.
 * @param {Number} max - Quantidade máxima aceitável no estoque.
 * @param {Number} quantidade - Quantidade atual do item no estoque.
 * @returns {object} - Mensagem de erro ou sucesso.
 * @throws Retorna erro em caso de falha de conexão com a API ou servidor.
 */
async function criarItemEstoque(codComp, min, max, quantidade) {
    return await fetch(`/estoque`, {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            codComp,
            min,
            max,
            quantidade
        })
    })
    .then((res) => res.json())
    .then((res) => {
        if (res.error) {
            mostrarMensagemErro(res.error);
            return new Error(res.error);
        }

        return res;
    })
    .catch((err) => {
        mostrarMensagemErro("Erro ao conectar com o servidor. Tente novamente mais tarde.");
        return new Error(err);
    });
}

/**
 * Envia dados de item do estoque para atualização à API.
 * @param {Number} codEstq - Código do registro no estoque.
 * @param {Number} codComp - Código do componente.
 * @param {Number} min - Quantidade mínima aceitável no estoque.
 * @param {Number} max - Quantidade máxima aceitável no estoque.
 * @param {Number} quantidade - Quantidade atual do item no estoque.
 * @returns {object} - Mensagem de erro ou sucesso.
 * @throws Retorna erro em caso de falha de conexão com a API ou servidor.
 */
async function atualizarItemEstoque(codEstq, codComp, min, max, quantidade) {
    return await fetch(`/estoque`, {
        method: "PUT",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            codEstq,
            codComp,
            min,
            max,
            quantidade
        })
    })
    .then((res) => res.json())
    .then((res) => {
        if (res.error) {
            mostrarMensagemErro(res.error);
            return new Error(res.error);
        }

        return res;
    })
    .catch((err) => {
        mostrarMensagemErro("Erro ao conectar com o servidor. Tente novamente mais tarde.");
        return new Error(err);
    });
}

/**
 * Remove item do estoque e mostra resposta.
 * @param {string} codComp - Código do componente a ser 
 * removido do estoque.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function excluirItemEstoque(codComp) {
    const res = await fetch(`/estoque/${codComp}`, {
        method: "DELETE"
    });

    if (res.error) {
        mostrarMensagemErro(res.error);
        return new Error(res.error);
    }

    window.location.reload();
}

/**
 * Confirma se o usuário realmente deseja excluir o estoque.
 * @returns {boolean} - Retorna confirmação do usuário.
 */
function confirmarExclusaoEstoque() {
    return mostrarPopup("Espere! Ao excluir um item diretamente no estoque, ele não será excluido em entradas Isso pode levar a complicações no dashboard e outras partes do sistema.")
        .then(result => {
            return result;
        });
}

/**
 * Avisa ao usuário sobre criação de item diretamente no estoque.
 * @returns {boolean} - Retorna confirmação do usuário.
 */
function confirmarCriacaoEstoque() {
    return mostrarPopup("Espere! Ao criar um item diretamente no estoque, ele não será criado em entradas Isso pode levar a complicações no dashboard e outras partes do sistema.")
        .then(result => {
            return result;
        });
}

/**
 * Avisa ao usuário sobre edição de item diretamente no estoque.
 * @returns {boolean} - Retorna confirmação do usuário.
 */
function confirmarAtualizacaoEstoque() {
    return mostrarPopup("Espere! Ao editar um item diretamente no estoque, ele não será editado em entradas Isso pode levar a complicações no dashboard e outras partes do sistema.")
        .then(result => {
            return result;
        });
}

/**
 * Mostra o forms para edição de dados de estoque.
 * @param {object} estq - Dados do estoque a ser alterado.
 */
function mostrarFormEdicaoEstoque(estq) {
    const confirm = confirmarAtualizacaoEstoque();
    confirm.then(() => {
        const editEstoqueForm = document.getElementById("edit-form-estoque");
    
        const codEstqTextbox = document.getElementById("edit-codigo-estoque-estoque");
        const codCompTextbox = document.getElementById("edit-codigo-componentes-estoque");
        const quantMinTextbox = document.getElementById("edit-quantidade-min-estoque");
        const quantMaxTextbox = document.getElementById("edit-quantidade-max-estoque");
        const quantidadeTextbox = document.getElementById("edit-quantidade-estoque");

        editEstoqueForm.style.display = "block";

        codEstqTextbox.value = estq.codEstq;
        codCompTextbox.value = estq.codComp;
        quantMinTextbox.value = estq.min;
        quantMaxTextbox.value = estq.max;
        quantidadeTextbox.value = estq.quantidade;
    })
    .catch((err) => {
        const editEstoqueForm = document.getElementById("edit-form-estoque");
        editEstoqueForm.style.display = "none";
    });
}

/**
 * Mostra forms para cadastro de estoque.
 */
function mostrarFormCriacaoItemEstoque(event) {
    event.preventDefault();
    const confirm = confirmarCriacaoEstoque();
    confirm.then(() => {
        const addEstoqueForm = document.getElementById("add-form-estoque");
        addEstoqueForm.style.display = "block";
    })
    .catch((err) => {
        const addEstoqueForm = document.getElementById("add-form-estoque");
        addEstoqueForm.style.display = "none";
    });
}

/**
 * Retorna estoque com nome contendo uma dada string.
 * @param {string} codComp - String para busca de nome.
 * @param {Array} estq - Array com todos os itens do estoque.
 */
function procurarItemEstoque(codComp, estq) {
    if (!codComp)
        return estq;

    const estoque = [];
    for (const item of estq) {
        if (item.codComp == codComp) 
            estoque.push(item);
    }

    return estoque;
}
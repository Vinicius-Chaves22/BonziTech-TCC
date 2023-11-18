let dadosComponentes, codComp;

/**
 * Formatação do texto para truncar o tamanho maximo de exibição.
 * @param {String} str - Valor a ser truncado
 * @param {String} num - tamanho maximo
 */
function truncateString(str, num) {
    if (str.length > num) {
      return str.slice(0, num) + "...";
    } else {
      return str;
    }
  }

/**
 * Formatação do texto para valor monetário.
 * @param {String} stringNumber - Valor a ser formatado
 */
function parseLocaleNumber(stringNumber) {
    var valor = stringNumber.replace(/\D/g,'');
    valor = (valor/100).toFixed(2) + '';
    valor = valor.replace(".", ",");
    valor = valor.replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1.');
    return valor;
}

const valorEntradaAdd = document.getElementById("add-valor-entrada-componentes");
valorEntradaAdd.onkeyup = function () {
    valorEntradaAdd.value = parseLocaleNumber(valorEntradaAdd.value);
};

const valorSaidaAdd = document.getElementById("add-valor-saida-componentes");  
valorSaidaAdd.onkeyup = function () {
    valorSaidaAdd.value = parseLocaleNumber(valorSaidaAdd.value);
};

const valorEntradaEdit = document.getElementById("edit-valor-entrada-componentes");
valorEntradaEdit.onkeyup = function () {
    valorEntradaEdit.value = parseLocaleNumber(valorEntradaEdit.value);
};

const valorSaidaEdit = document.getElementById("edit-valor-saida-componentes");  
valorSaidaEdit.onkeyup = function () {
    valorSaidaEdit.value = parseLocaleNumber(valorSaidaEdit.value);
};

function detectFloat(source) {
    return source.replace(/(\.|\s)|(\,)/g,(m,p1,p2) => p1 ? "" : ".");
}

/**
 * Adicão de nulos para campos não usados no formulário de insert e edit dos componentes.
 * @param {string} codSubcat        - código da sub categoria do componente, representado por numero.
 * @param {string} diamInterno      - diametro interno do componente, representado por numero.
 * @param {string} diamExterno      - diametro externo do componente, representado por numero.
 * @param {string} diamNominal      - diametro nominal do componente, representado por numero.
 * @param {string} medidaD          - medida d do componente, representado por numero.
 * @param {string} costura          - se tem costura no componente, representado por numero.
 * @param {string} prensadoReusavel - se é prensado ou reutilizável o componente, representado por numero.
 * @param {string} mangueira        - valor da mangueira componente, representado por numero.
 * @param {string} material         - descrição/codigo material do componente, representado por numero.
 * @param {string} norma            - norma do componente, representado por numero.
 * @param {string} bitola           - tamanho da bitola componente, representado por numero.
 */
function colocaNull(
    codSubcat, diamInterno, diamExterno, diamNominal, medidaD, costura, 
    prensadoReusavel, mangueira, material, norma, bitola
    ){

    if ((codSubcat == '') || (codSubcat == 'null')) {
        codSubcat = null;
    } else {
        codSubcat = Number(codSubcat);
    }

    if ((diamInterno == '') || (diamInterno == 'null')) {
        diamInterno = null;
    }

    if ((diamExterno == '') || (diamExterno == 'null')) {
        diamExterno = null;
    } else {
        diamExterno = Number(diamExterno);
    }

    if ((diamNominal == '') || (diamNominal == 'null')) {
        diamNominal = null;
    }

    if ((medidaD == '') || (medidaD == 'null')) {
        medidaD = null;
    } else {
        medidaD = Number(medidaD);
    }

    if ((costura == '') || (costura == 'null')) {
        costura = null;
    } else {
        costura = (costura.toLowerCase() === "true");
    }

    if ((prensadoReusavel == '') || (prensadoReusavel == 'null')) {
        prensadoReusavel = null;
    }

    if ((mangueira == '') || (mangueira == 'null')) {
        mangueira = null;
    }

    if ((material == '') || (material == 'null')) {
        material = null;
    }

    if ((norma == '') || (norma == 'null')) {
        norma = null;
    }

    if ((bitola == '') || (bitola == 'null')) {
        bitola = null;
    } else {
        bitola = Number(bitola);
    }
}

fetchComponentes()
.then((res) => {
    mostrarTabelaComponentes(res.componentes);
}); 

fetchCategorias()
.then((res) => {
    mostrarCategorias(res.categorias);
}); 

const addcomponentesFormBtn = document.getElementById("add-table-row-componentes");
const cancelarCriacaoComponentesIcone = document.getElementById("cancel-componente-post-icon");
const cancelarCriacaoComponentesBtn = document.getElementById("cancel-criacao-componente-btn");

addcomponentesFormBtn.addEventListener("click", mostrarFormCriacaoComponentes);
cancelarCriacaoComponentesIcone.addEventListener("click", mostrarFormCriacaoComponentes);
cancelarCriacaoComponentesBtn.addEventListener("click", mostrarFormCriacaoComponentes);

const cancelarEdicaoComponentesIcone = document.getElementById("cancel-componente-edit-icon");
const cancelarEdicaoComponentesBtn = document.getElementById("edit-componente-cancel-btn");

cancelarEdicaoComponentesIcone.addEventListener("click", mostrarFormEdicaoComponentes);
cancelarEdicaoComponentesBtn.addEventListener("click", mostrarFormEdicaoComponentes);

const searchBar = document.getElementById("search-bar");
searchBar.addEventListener("keyup", () => {
    fetchComponentes()
    .then((res) => {
        console.log(searchBar.value.trim());
        console.log(res);
        let compArr = procurarComponentes(searchBar.value.trim(), res.componentes);
        console.log("======================");
        console.log(compArr);
        console.log("======================");
        mostrarTabelaComponentes(compArr);
        console.log("======================");
    });    
});

const confirmarCriacaoComponentesBtn = document.getElementById("confirm-criacao-componente-btn");

//Criação dos componentes, com tratamento das variáves
confirmarCriacaoComponentesBtn.addEventListener("click", async (event) => {
    event.preventDefault();

    const codpeca        	    = document.getElementById("add-codpeca-componentes");
    const especificacao    	    = document.getElementById("add-especificacao-componentes");
    const categoria        	    = document.getElementById("add-categoria-componentes");
    const subcategoria        	= document.getElementById("add-subcategoria-componentes");
    const diametroInt 			= document.getElementById("add-dia-interno-componentes");
    const diametroExt 			= document.getElementById("add-dia-externo-componentes");
    const diametroNom 			= document.getElementById("add-dia-nominal-componentes");
    const medida      			= document.getElementById("add-medida-componentes");
    const costura     			= document.getElementById("add-costura-componentes");
    const prensadoReusavel      = document.getElementById("add-pren-reusavel-componentes");
    const mangueira      		= document.getElementById("add-mangueira-componentes");
    const material      		= document.getElementById("add-material-componentes");
    const norma      			= document.getElementById("add-norma-componentes");
    const bitola      			= document.getElementById("add-bitola-componentes");
    const valorEntrada      	= document.getElementById("add-valor-entrada-componentes");
    const valorSaida      		= document.getElementById("add-valor-saida-componentes");  

    await criarComponentes(
            codpeca.value.trim(), 
            especificacao.value.trim(), 
            categoria.value.trim(), 
            subcategoria.value.trim(), 
            diametroInt.value.trim(), 
            diametroExt.value.trim(), 
            diametroNom.value.trim(), 
            medida.value.trim(), 
            costura.value.trim(), 
            prensadoReusavel.value.trim(), 
            mangueira.value.trim(), 
            material.value.trim(), 
            norma.value.trim(), 
            bitola.value.trim(), 
            valorEntrada.value.trim(), 
            valorSaida.value.trim()
    ).then((res) => {
        if ('message' in res) {
            if (res.message == "Componente cadastrado com sucesso!"){
                const form = document.getElementById("add-formulario-componente");
                form.reset();
                window.location.reload();
                return;
            }
        } else {
            mostrarMensagemErro(res.Error);
            return;
        }
    })
    .catch((err) => {
        mostrarMensagemErro("Erro ao inserir componente.");
        return new Error(err);
    });
});

const confirmarAtualizacaocomponentesBtn = document.getElementById("update-componente-btn");

//Update dos componentes, com tratamento das variáves
confirmarAtualizacaocomponentesBtn.addEventListener("click", async (event) => {
    event.preventDefault();

    const codpeca        	    = document.getElementById("edit-codpeca-componentes");
    const especificacao    	    = document.getElementById("edit-especificacao-componentes");
    const categoria        	    = document.getElementById("edit-categoria-componentes");
    const subcategoria        	= document.getElementById("edit-subcategoria-componentes");
    const diametroInt 			= document.getElementById("edit-dia-interno-componentes");
    const diametroExt 			= document.getElementById("edit-dia-externo-componentes");
    const diametroNom 			= document.getElementById("edit-dia-nominal-componentes");
    const medida      			= document.getElementById("edit-medida-componentes");
    const costura     			= document.getElementById("edit-costura-componentes");
    const prensadoReusavel      = document.getElementById("edit-pren-reusavel-componentes");
    const mangueira      		= document.getElementById("edit-mangueira-componentes");
    const material      		= document.getElementById("edit-material-componentes");
    const norma      			= document.getElementById("edit-norma-componentes");
    const bitola      			= document.getElementById("edit-bitola-componentes");
    const valorEntrada      	= document.getElementById("edit-valor-entrada-componentes");
    const valorSaida      		= document.getElementById("edit-valor-saida-componentes"); 

    await atualizarComponentes(
        codComp,
        codpeca.value.trim(),        	 
        categoria.value.trim(),
        subcategoria.value.trim(),
        especificacao.value.trim(),
        diametroInt.value.trim(), 	
        diametroExt.value.trim(), 	
        diametroNom.value.trim(), 	
        medida.value.trim(),      	
        costura.value.trim(),     	
        prensadoReusavel.value.trim(), 
        mangueira.value.trim(),      	
        material.value.trim(),      	
        norma.value.trim(),      		
        bitola.value.trim(),      	
        valorEntrada.value.trim(),     
        valorSaida.value.trim()
    );

    window.location.reload();
});


/**
 * Mostra tabela de componentes com seus devidos dados.
 * @param {Array} dadosTabela - Lista de componentes a
 * serem mostrados.
 */
async function mostrarTabelaComponentes(dadosTabela) {
    const tbody = document.getElementById("tbody-componentes");
    tbody.innerHTML = "";
    
    if ((dadosTabela == null) || (dadosTabela.length === 0)) {
        tbody.innerHTML = `
        <tr><td colspan="17"><h2 class="texto-404"> Nenhum componente encontrado. </h2><td/></tr>`;
    } else {
        for (const comp of dadosTabela) {
            const coluna = document.createElement("tr");
            
            let codPeca = comp.codPeca;
            let especificacao = comp.especificacao;
            let codCat = comp.codCat;
            let codSubcat = comp.codSubcat;
            let diamInterno = comp.diamInterno;
            let diamExterno = comp.diamExterno;
            let diamNominal = comp.diamNominal;
            let medidaD = comp.medidaD;
            let costura = comp.costura;
            let prensadoReusavel = comp.prensadoReusavel;
            let mangueira = comp.mangueira;
            let material = comp.material;
            let norma = comp.norma;
            let bitola = comp.mangueira;
            let valorEntrada = comp.valorEntrada;
            let valorVenda = comp.valorVenda;

            if ( codPeca === null || codPeca === undefined ){
                codPeca = "";
            }

            if ( especificacao === null || especificacao === undefined ){
                especificacao = "";
            }

            especificacao = truncateString(especificacao, 30);

            if ( codCat === null || codCat === undefined ){
                codCat = "";
            }

            if ( codSubcat === null || codSubcat === undefined ){
                codSubcat = "";
            }
            
            if ( material === null || material === undefined ){
                material = "";
            }

            if ( valorEntrada === null || valorEntrada === undefined ){
                valorEntrada = "";
            }
            
            if ( valorVenda === null || valorVenda === undefined ){
                valorVenda = "";
            }

            if ( diamInterno === null || diamInterno === undefined ){
                diamInterno = "";
            }

            if ( diamExterno === null || diamExterno === undefined ){
                diamExterno = "";
            }
            
            if ( diamNominal === null || diamNominal === undefined ){
                diamNominal = "";
            }

            if ( medidaD === null || medidaD === undefined ){
                medidaD = "";
            }

            if ( costura === null || costura === undefined ){
                costura = "";
            }

            if ( prensadoReusavel === null || prensadoReusavel === undefined ){
                prensadoReusavel = "";
            }
            
            if ( mangueira === null || mangueira === undefined ){
                mangueira = "";
            }       

            coluna.innerHTML = `
            <td> ${codPeca} </td>
            <td> ${especificacao} </td>
            <td> ${codCat} </td>
            <td> ${codSubcat} </td>
            <td> ${diamInterno} </td>
            <td> ${diamExterno} </td>
            <td> ${diamNominal} </td>
            <td> ${medidaD} </td>
            <td> ${costura} </td>
            <td> ${prensadoReusavel} </td>
            <td> ${mangueira} </td>
            <td> ${material} </td>
            <td> ${norma} </td>
            <td> ${bitola} </td>
            <td> ${valorEntrada} </td>
            <td> ${valorVenda} </td>
            `;

            const acoesCell = document.createElement("td");
    
            // Botão de inativação de componente
            const btnDelete = document.createElement("button");
            btnDelete.classList.add("delete-btn");
            btnDelete.addEventListener("click", () => {
                if(confirmarInativacaoComponentes())
                    inativarComponentes(comp.codComp);
            });
            btnDelete.innerHTML = '<i class="fa-solid fa-ban"> </i>';
    
            // Botão de edição de componente
            const btnEdit = document.createElement("button");
            btnEdit.classList.add("update-btn-icon");
            btnEdit.addEventListener("click", (event) => {
                event.preventDefault();
                mostrarFormEdicaoComponentes(comp);
                codComp = comp.codComp;
            });
            btnEdit.innerHTML = '<i class="fa-solid fa-pen-to-square"> </i>';
    
            acoesCell.appendChild(btnDelete);
            acoesCell.appendChild(btnEdit);
    
            coluna.appendChild(acoesCell);
    
            tbody.appendChild(coluna);
        }
    }

    
}

/**
 * Pega todos os componentes cadastrados pela API e
 * os insere na tabela.
 * @returns {object} - Resposta da API.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function fetchComponentes() {
    return await fetch("/componentes/")
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
 * Pega todos os componentes cadastrados pela API e
 * os insere na tabela.
 * @returns {object} - Resposta da API.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function fetchCategorias() {
    return await fetch("/categorias/")
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
 * Pega todos os componentes cadastrados pela API e
 * os insere na tabela.
 * @returns {object} - Resposta da API.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function fetchSubCategorias(codCat) {
	
	//GET /api/subcategorias/categoria/${codCat}	
    return await fetch(`/subcategorias/categoria/${codCat}`)
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
 * Envia dados de componente para criação à API.
 * @param {string} codPeca          - texto interno de código da peça componnente a ser cadastrado.
 * @param {string} especificacao    - especificacao do componente.
 * @param {string} codCat           - código da categoria do componente, representado por numero.
 * @param {string} codSubcat        - código da sub categoria do componente, representado por numero.
 * @param {string} diamInterno      - diametro interno do componente, representado por numero.
 * @param {string} diamExterno      - diametro externo do componente, representado por numero.
 * @param {string} diamNominal      - diametro nominal do componente, representado por numero.
 * @param {string} medidaD          - medida d do componente, representado por numero.
 * @param {string} costura          - se tem costura no componente, representado por numero.
 * @param {string} prensadoReusavel - se é prensado ou reutilizável o componente, representado por numero.
 * @param {string} mangueira        - valor da mangueira componente, representado por numero.
 * @param {string} material         - descrição/codigo material do componente, representado por numero.
 * @param {string} norma            - norma do componente, representado por numero.
 * @param {string} bitola           - tamanho da bitola componente, representado por numero.
 * @param {string} valorEntrada     - valor de entrada do componente, representado por numero.
 * @param {string} valorVenda       - valor de venda do componente, representado por numero.
 * @returns {object} - Mensagem de erro ou sucesso.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function criarComponentes( 
    codPeca, especificacao, codCat, codSubcat, diamInterno, diamExterno, 
    diamNominal, medidaD, costura, prensadoReusavel, 
    mangueira, material, norma, bitola, valorEntrada, valorVenda) {

    colocaNull(
        codSubcat, diamInterno, diamExterno, 
        diamNominal, medidaD, costura, 
        prensadoReusavel, mangueira, material,  
        norma, bitola)

    console.log({
        codPeca, 
        especificacao, 
        "codCat": Number(codCat), 
        "codSubcat": codSubcat, 
        diamInterno, 
        "diamExterno": Number(diamExterno), 
        diamNominal, 
        "medidaD": Number(medidaD), 
        "costura": (costura == "true"), 
        prensadoReusavel, 
        mangueira, 
        material, 
        norma, 
        "bitola": bitola, 
        "valorEntrada": parseFloat(detectFloat(valorEntrada)), 
        "valorVenda": parseFloat(detectFloat(valorVenda))
    });

    //execução do request de envio do JSON
    return await fetch(`/componentes`, {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify(
            {
                codPeca, 
                especificacao, 
                "codCat": Number(codCat), 
                "codSubcat": (codSubcat === 'null' ? null : Number(codSubcat)), 
                diamInterno, 
                "diamExterno": Number(diamExterno), 
                diamNominal, 
                "medidaD": Number(medidaD), 
                "costura": (costura == "true"), 
                prensadoReusavel, 
                mangueira, 
                material, 
                norma, 
                "bitola": bitola, 
                "valorEntrada": parseFloat(detectFloat(valorEntrada)), 
                "valorVenda": parseFloat(detectFloat(valorVenda))
        })

    })
    .then((res) => {
        return res.json();
    })
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
 * Envia dados de componente para atualização à API.
 * @param {string} codPeca          - texto interno de código da peça componnente a ser cadastrado.
 * @param {string} especificacao    - especificacao do componente.
 * @param {string} codCat           - código da categoria do componente, representado por numero.
 * @param {string} codSubcat        - código da sub categoria do componente, representado por numero.
 * @param {string} diamInterno      - diametro interno do componente, representado por numero.
 * @param {string} diamExterno      - diametro externo do componente, representado por numero.
 * @param {string} diamNominal      - diametro nominal do componente, representado por numero.
 * @param {string} medidaD          - medida d do componente, representado por numero.
 * @param {string} costura          - se tem costura no componente, representado por numero.
 * @param {string} prensadoReusavel - se é prensado ou reutilizável o componente, representado por numero.
 * @param {string} mangueira        - valor da mangueira componente, representado por numero.
 * @param {string} material         - descrição/codigo material do componente, representado por numero.
 * @param {string} norma            - norma do componente, representado por numero.
 * @param {string} bitola           - tamanho da bitola componente, representado por numero.
 * @param {string} valorEntrada     - valor de entrada do componente, representado por numero.
 * @param {string} valorVenda       - valor de venda do componente, representado por numero.
 * @returns {object} - Mensagem de erro ou sucesso.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function atualizarComponentes(
    codComp, codPeca, codCat, 
    codSubcat, especificacao, diamInterno, 
    diamExterno, diamNominal, medidaD, 
    costura, prensadoReusavel, mangueira, 
    material, norma, bitola, 
    valorEntrada, valorVenda) {

    colocaNull(
        codSubcat, diamInterno, diamExterno, 
        diamNominal, medidaD, costura, 
        prensadoReusavel, mangueira, material,  
        norma, bitola)

    return await fetch(`/componentes`, {
        method: "PUT",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            "codComp": Number(codComp),
            codPeca, 
            especificacao, 
            "codCat": Number(codCat), 
            "codSubcat": (codSubcat === 'null' ? null : Number(codSubcat)), 
            diamInterno, 
            "diamExterno": Number(diamExterno), 
            diamNominal, 
            "medidaD": Number(medidaD), 
            "costura": (costura == "true"), 
            prensadoReusavel, 
            mangueira, 
            material, 
            norma, 
            "bitola": bitola, 
            "valorEntrada": parseFloat(detectFloat(valorEntrada)), 
            "valorVenda": parseFloat(detectFloat(valorVenda))
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
 * Inativa componente no banco de dados e mostra resposta.
 * @param {string} codComp - Código do componente a ser excluido.
 * @throws Retorna erro em caso de falha de conexão com a 
 * API ou servidor.
 */
async function inativarComponentes(codComp) {
    console.log("Excluir:"+codComp);
    const res = await fetch(`/componentes/${codComp}`, {
        method: "DELETE"
    });

    if (res.error) {
        mostrarMensagemErro(res.error);
        return new Error(res.error);
    }

    window.location.reload();
}

/**
 * Confirma se o usuário realmente deseja inativar o componente.
 * @returns {boolean} - Retorna confirmação do usuário.
 */
function confirmarInativacaoComponentes() {
    return confirm("Você tem certeza que deseja excluir este componente?");
}

/**
 * Mostra o forms para edição de componente.
 * @param {object} comp - Dados do componente a ser alterado.
 */
function mostrarFormEdicaoComponentes(comp) {
    const editcomponentesForm = document.getElementById("edit-componentes-form");
    
    if (editcomponentesForm.style.display !== "block") {
        editcomponentesForm.style.display = "block";

        let codpecaTextbox        	    = document.getElementById("edit-codpeca-componentes");
        let especificacaoTextbox        = document.getElementById("edit-especificacao-componentes");
        let categoriaTextbox        	= document.getElementById("edit-categoria-componentes");
        let subcategoriaTextbox        	= document.getElementById("edit-subcategoria-componentes");
        let diametroIntTextbox 			= document.getElementById("edit-dia-interno-componentes");
        let diametroExtTextbox 			= document.getElementById("edit-dia-externo-componentes");
        let diametroNomTextbox 			= document.getElementById("edit-dia-nominal-componentes");
        let medidaTextbox      			= document.getElementById("edit-medida-componentes");
        let costuraTextbox     			= document.getElementById("edit-costura-componentes");
        let prensadoReusavelTextbox     = document.getElementById("edit-pren-reusavel-componentes");
        let mangueiraTextbox      		= document.getElementById("edit-mangueira-componentes");
        let materialTextbox      		= document.getElementById("edit-material-componentes");
        let normaTextbox      			= document.getElementById("edit-norma-componentes");
        let bitolaTextbox      			= document.getElementById("edit-bitola-componentes");
        let valorEntradaTextbox      	= document.getElementById("edit-valor-entrada-componentes");
        let valorSaidaTextbox      		= document.getElementById("edit-valor-saida-componentes");        
        
        codpecaTextbox.value            = comp.codPeca;
        especificacaoTextbox.value      = comp.especificacao;
        categoriaTextbox.value          = comp.codCat;
        subcategoriaTextbox.value       = comp.codSubcat;
        diametroIntTextbox.value        = comp.diamInterno;
        diametroExtTextbox.value        = comp.diamExterno;
        diametroNomTextbox.value        = comp.diamNominal;
        medidaTextbox.value             = comp.medidaD;

        for (var i = 0; i < costuraTextbox.options.length; i++)
        {
            if (comp.diamNominal == costuraTextbox.options[i].value)
            costuraTextbox.options[i].selected = true;
        }

        prensadoReusavelTextbox.value   = comp.prensadoReusavel;
        mangueiraTextbox.value          = comp.mangueira;
        materialTextbox.value           = comp.material;
        normaTextbox.value              = comp.norma;
        bitolaTextbox.value             = comp.bitola;

        valorEntradaTextbox.value       = comp.valorEntrada;
        valorSaidaTextbox.value         = comp.valorVenda;

    } else {
        editcomponentesForm.style.display = "none";
    }
}

/**
 * Mostra forms para criação de componente.
 */
function mostrarFormCriacaoComponentes() {
    const addcomponentesForm = document.getElementById("add-componentes-form");
    const firstField = document.getElementById("add-codpeca-componentes");

    if (addcomponentesForm.style.display !== "block") {
        addcomponentesForm.style.display = "block";
    } else {
        addcomponentesForm.style.display = "none";
    }
        
    firstField.focus();
}

/**
 * Mostra combo de componentes para criação de componente.
 * @param {Array} dadosCategorias - Array com os dados coletados da API de categorias.
 */
function mostrarCategorias(dadosCategorias) {
	const comboCat = document.getElementById("add-categoria-componentes");
	const comboSubCat = document.getElementById("add-subcategoria-componentes");
	
	//controle de seleção da categoria para exibir a subcategoria correspondente
	comboCat.addEventListener("change", function(){
		const codCat = comboCat.options[comboCat.selectedIndex].value;
		fetchSubCategorias(codCat)
			.then((res) => {
				mostrarSubCategorias(res.subcategorias, comboSubCat);
			}); 
	});
	
	//carrega as categorias da adicionar
	for (const comp of dadosCategorias) {
			var opt = document.createElement("option");
			opt.value = comp.codCat;
			opt.text = comp.nomeCat;
			comboCat.add(opt, comboCat.options[comboCat.selectedIndex + 1]);		    
	}
	
	
	const comboCat2 = document.getElementById("edit-categoria-componentes");
	const comboSubCat2 = document.getElementById("edit-subcategoria-componentes");
	
	//controle de seleção da categoria para exibir a subcategoria correspondente
	comboCat2.addEventListener("change", function(){
		const codCat = comboCat2.options[comboCat2.selectedIndex].value;
		fetchSubCategorias(codCat)
			.then((res) => {
				mostrarSubCategorias(res.subcategorias, comboSubCat2);
			}); 
	});
	
	//carrega as categorias da edição
	for (const comp of dadosCategorias) {
			var opt = document.createElement("option");
			opt.value = comp.codCat;
			opt.text = comp.nomeCat;
			comboCat2.add(opt, comboCat2.options[comboCat2.selectedIndex + 1]);		    
	}
}

/**
 * Mostra combo de componentes para criação de componente.
 * @param {Array} dadosCategorias - Array com os dados coletados da API de categorias.
 */
function mostrarSubCategorias(dadosSubCategorias, comboSubCat) {
	comboSubCat.innerHTML = '<option value="null" selected>Selecionar Sub-Categoria</option>';
	//carrega as subcategorias da a
	for (const comp of dadosSubCategorias) {
			var opt = document.createElement("option");
			opt.value = comp.codSubcat;
			opt.text = comp.nome;
			comboSubCat.add(opt, comboSubCat.options[comboSubCat.selectedIndex + 1]);		    
	}

}

/**
 * Retorna componentes com especificacao contendo uma
 * dada string.
 * @param {string} str - String para busca de especificacao.
 * @param {Array} compArr - Array com componentes 
 * a serem buscados.
 */
function procurarComponentes(str, compArr) {
    const componentes = [];
    for (const comp of compArr) {
        if (String(comp.codPeca).includes(str))
            componentes.push(comp);            
        else if (comp.especificacao.includes(str)) 
            componentes.push(comp);
        else if (String(comp.codCat).includes(str)) 
            componentes.push(comp);
        else if (String(comp.codSubcat).includes(str)) 
            componentes.push(comp);
        else if (String(comp.diamInterno).includes(str)) 
            componentes.push(comp);
        else if(String(comp.diamExterno).includes(str)) 
            componentes.push(comp);
        else if(String(comp.diamNominal).includes(str)) 
            componentes.push(comp);
        else if (String(comp.medidaD).includes(str)) 
            componentes.push(comp);
        else if (String(comp.costura).includes(str)) 
            componentes.push(comp);
        else if (String(comp.prensadoReusavel).includes(str)) 
            componentes.push(comp);
        else if (String(comp.mangueira).includes(str)) 
            componentes.push(comp);
        else if (String(comp.material).includes(str)) 
            componentes.push(comp);
        else if (String(comp.norma).includes(str)) 
            componentes.push(comp);
        else if (String(comp.bitola).includes(str)) 
            componentes.push(comp);
        else if (String(comp.valorEntrada).includes(str)) 
            componentes.push(comp);
        else if (String(comp.valorVenda).includes(str)) 
            componentes.push(comp);
    }

    return componentes;
}
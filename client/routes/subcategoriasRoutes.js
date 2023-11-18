require("dotenv").config();
const express = require("express");
const compRouter = express.Router();

const validarSessao = require("../modules/validarSessao");

compRouter.get("/categoria/:codCat", validarSessao, (req, res) => {
    const codCat = req.params.codCat;

    fetch(process.env.APIURL + `/subcategorias/categoria/${codCat}`, {
        method: "GET",
        headers: {
            "Content-type": "Application/JSON",
            "codSessao": req.session.codSessao
        }
    })
    .then((response) => response.json())
    .then((response) => res.json(response))
    .catch((err) => {
        console.log(err);
        res.status(500).json({ error: "Erro ao conectar com o servidor." });
    });
});

module.exports = compRouter;
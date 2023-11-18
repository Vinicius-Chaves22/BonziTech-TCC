require("dotenv").config();
const express = require("express");
const compRouter = express.Router();

const validarSessao = require("../modules/validarSessao");

compRouter.get("/", validarSessao, (req, res) => {
    fetch(process.env.APIURL + `/categorias`, {
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

compRouter.get("/:codCat", validarSessao, (req, res) => {
    const codComp = req.params.codComp;

    fetch(process.env.APIURL + `/categorias/${codCat}`, {
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
require("dotenv").config();
const express = require("express");
const compRouter = express.Router();

const validarSessao = require("../modules/validarSessao");

compRouter.get("/", validarSessao, (req, res) => {
    fetch(process.env.APIURL + `/componentes`, {
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

compRouter.get("/:codComp", validarSessao, (req, res) => {
    const codComp = req.params.codComp;

    fetch(process.env.APIURL + `/componentes/${codComp}`, {
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

compRouter.post("/", validarSessao, (req, res) => {
    fetch(process.env.APIURL + `/componentes`, {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON",
            "codSessao": req.session.codSessao
        },
        body: JSON.stringify(req.body)
    })
    .then((response) => response.json())
    .then((response) => res.json(response))
    .catch((err) => {
        console.log(err);
        res.status(500).json({ error: "Erro ao conectar com o servidor." });
    });
});

compRouter.put("/", validarSessao, (req, res) => {
    fetch(process.env.APIURL + `/componentes`, {
        method: "PUT",
        headers: {
            "Content-type": "Application/JSON",
            "codSessao": req.session.codSessao
        },
        body: JSON.stringify(req.body)
    })
    .then((response) => response.json())
    .then((response) => res.json(response))
    .catch((err) => {
        console.log(err);
        res.status(500).json({ error: "Erro ao conectar com o servidor." });
    });
});

compRouter.delete("/:codComp", validarSessao, (req, res) => {
    const codComp = req.params.codComp;

    fetch(process.env.APIURL + `/componentes/${codComp}`, {
        method: "DELETE",
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
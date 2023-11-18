require("dotenv").config();
const express = require("express");
const session = require("express-session");
const app = express();

app.use(express.json());
app.use(session({
    secret: process.env.key,
    resave: false,
    saveUninitialized: false,
    cookie: {
        secure: false
    }
}));

app.use(express.static("./public/"));
app.use("/", require("./routes/routes"));
app.use("/clientes", require("./routes/clientesRoutes"));
app.use("/componentes", require("./routes/componentesRoutes"));
app.use("/categorias", require("./routes/categoriasRoutes"));
app.use("/subcategorias", require("./routes/subcategoriasRoutes"));
app.use("/entradas", require("./routes/entradasRoutes"));
app.use("/estoque", require("./routes/estoqueRoutes"));
app.use("/fabricantes", require("./routes/fabricantesRoutes"));
app.use("/funcionarios", require("./routes/funcionariosRoutes"));
app.use("/saidas", require("./routes/saidasRoutes"));
app.use("/sessao", require("./routes/sessaoRoutes"));

const port = process.env.PORT || 5000;
app.listen(port, () => console.log('Server is running -> ' + port));

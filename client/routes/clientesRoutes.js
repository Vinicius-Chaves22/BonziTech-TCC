require("dotenv").config();
const express = require("express");
const cliRouter = express.Router();

const path = require("path");
const publicFolder = "../public/";


module.exports = cliRouter;
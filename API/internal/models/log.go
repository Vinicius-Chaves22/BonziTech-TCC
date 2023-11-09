package models

type Log struct {
	CodLog    int    `json:"codLog"`
	TipoReq   string `json:"tipoReq"`
	Caminho   string `json:"caminho"`
	StatusRes int    `json:"statusRes"`
	CodSessao int    `json:"codSessao"`
	Data      string `json:"data"`
}

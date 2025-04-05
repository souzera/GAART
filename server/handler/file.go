package handler

import "github.com/gin-gonic/gin"

func UploadArquivo(contexto *gin.Context) {

	arquivo, err := contexto.FormFile("arquivo")
	if err != nil {
		sendError(contexto, 400, "Erro ao fazer upload do arquivo")
		return
	}

	logger.Infof("Arquivo: %v", arquivo.Filename)

	if err := contexto.SaveUploadedFile(arquivo, "./uploads/"+arquivo.Filename); err != nil {
		logger.Errorf("[UPLOAD-ARQUIVO] Error: %v", err)
		sendError(contexto, 500, "Erro ao salvar o arquivo")
		return
	}

	sendSucess(contexto, "upload-arquivo", "Arquivo salvo com sucesso")

}

func UploadMultiplosArquivos(contexto *gin.Context) {

	arquivos := contexto.Request.MultipartForm.File["arquivos"]
	if len(arquivos) == 0 {
		sendError(contexto, 400, "Erro ao fazer upload dos arquivos")
		return
	}

	for _, arquivo := range arquivos {
		logger.Infof("Arquivo: %v", arquivo.Filename)

		// TODO: loop para salvar os arquivos
	}

	sendSucess(contexto, "upload-arquivo", "Arquivos salvos com sucesso")

}

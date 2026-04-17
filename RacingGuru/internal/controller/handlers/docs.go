package handlers

import (
	_ "embed"
	"net/http"
)

var (
	//go:embed assets/swagger/swagger.yaml
	swaggerSpec []byte
	swaggerUI   = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>RacingGuru API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.ui = SwaggerUIBundle({
      url: '/swagger.yaml',
      dom_id: '#swagger-ui'
    });
  </script>
</body>
	</html>`
)

// SwaggerSpec godoc
// @Summary OpenAPI specification
// @Tags system
// @Produce plain
// @Success 200 {string} string
// @Router /swagger.yaml [get]
func (h *Handler) SwaggerSpec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(swaggerSpec)
}

// SwaggerUI godoc
// @Summary Swagger UI
// @Tags system
// @Produce html
// @Success 200 {string} string
// @Router /docs [get]
func (h *Handler) SwaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(swaggerUI))
}

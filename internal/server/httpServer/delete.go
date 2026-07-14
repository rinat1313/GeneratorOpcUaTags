package httpServer

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем название solution из query параметра
	solutionName := r.URL.Query().Get("name")
	if solutionName == "" {
		// Пробуем получить из тела запроса
		var requestData struct {
			Filename string `json:"filename"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err == nil && requestData.Filename != "" {
			solutionName = requestData.Filename
		}
	}

	// Проверяем, что имя файла указано
	if solutionName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Не указано имя решения",
		})
		return
	}

	// Очищаем имя файла от возможных путей для безопасности
	solutionName = filepath.Base(solutionName)

	// Убеждаемся, что файл имеет расширение .xml
	if !strings.HasSuffix(strings.ToLower(solutionName), ".xml") {
		solutionName = solutionName + ".xml"
	}

	// Формируем полный путь к файлу
	// !!! ВАЖНО: Укажите здесь правильный путь к директории с решениями !!!
	solutionsDir := "./data" // Замените на актуальный путь

	filePath := filepath.Join(solutionsDir, solutionName)

	// Проверяем существование файла
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  false,
			"error":    "Файл не найден",
			"filename": solutionName,
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Ошибка при проверке файла: " + err.Error(),
		})
		return
	}

	// Проверяем, что это действительно файл, а не директория
	if fileInfo.IsDir() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Указанный путь является директорией",
		})
		return
	}

	// !!! ЗДЕСЬ МЫ УДАЛЯЕМ ФАЙЛ !!!
	err = os.Remove(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Ошибка при удалении файла: " + err.Error(),
		})
		return
	}

	// Возвращаем успешный результат
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Файл успешно удален",
		"filename": solutionName,
	})
}

func init() {
	AddNewFunction(CreateHandlerCommand("/delete", "DELETE", DeleteHandler))
}
